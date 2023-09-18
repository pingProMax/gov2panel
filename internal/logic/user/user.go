package user

import (
	"context"
	"errors"
	v1 "gov2panel/api/admin/v1"
	userv1 "gov2panel/api/user/v1"
	"gov2panel/internal/dao"
	"gov2panel/internal/logic/cornerstone"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/model/model"
	"gov2panel/internal/service"
	"gov2panel/internal/utils"
	"strconv"
	"strings"
	"time"

	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

type sUser struct {
	Cornerstone cornerstone.Cornerstone
}

func init() {
	service.RegisterUser(New())
	auth := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "gov2panel",
		Key:             []byte("secret key"),
		Timeout:         time.Hour * 3,
		MaxRefresh:      time.Hour * 1,
		IdentityKey:     "TUserID",
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,
		Unauthorized:    Unauthorized,
		PayloadFunc:     PayloadFunc,
		IdentityHandler: IdentityHandler,
	})
	authService = auth
}

func New() *sUser {
	return &sUser{
		Cornerstone: *cornerstone.NewCornerstoneWithTable(dao.V2User.Table()),
	}
}

// 账号查询用户
func (s *sUser) GetUserByUserName(userName string) (user *entity.V2User, err error) {
	record, err := s.Cornerstone.GetDB().Where(dao.V2User.Columns().UserName, userName).One()
	if err != nil {
		return nil, err
	}
	err = record.Struct(&user)
	return
}

// 获取 user_id 的邀请数量
func (s *sUser) GetInviteCountByUserId(user_id int) (inviteCount int, err error) {
	inviteCount, err = s.Cornerstone.GetDB().Where(dao.V2User.Columns().InviteUserId, user_id).Count()
	return
}

// 获取 用户 的佣金比例\佣金类型
func (s *sUser) GetUserCTypeAndCRate(user *entity.V2User) (commissionType int, commissionRate int) {
	setting, _ := service.Setting().GetSettingAllMap()
	if user.CommissionType != 0 {
		commissionType = user.CommissionType
	} else {
		commissionType = setting["commission_type"].Int()
	}

	if user.CommissionRate != 0 {
		commissionRate = user.CommissionRate
	} else {
		commissionRate = setting["commission_rate"].Int()
	}

	return
}

// 计算用户佣金
// 用户
// 金额
func (s *sUser) CalculateUserCommission(CType, CRate int, fromUserId int, val float64) (commission float64, err error) {
	if CType == 2 { // 一次性的，判断是否已经有过佣金
		invitationRecords, err := service.InvitationRecords().GetOneByFromUserId(fromUserId)
		if utils.IgnoreErrNoRows(err) != nil {
			return 0, err
		}

		if invitationRecords != nil {
			return 0, err
		}

	}

	//计算佣金
	commission = utils.Decimal(val * float64(CRate) / 100)

	return
}

// 更新
func (s *sUser) UpUser(data *entity.V2User) (err error) {
	db := s.Cornerstone.GetDB().Data(data)
	if data.Password == "" {
		db.FieldsEx(dao.V2User.Columns().Password)
	}

	_, err =
		db.FieldsEx(
			dao.V2User.Columns().CreatedAt,
			dao.V2User.Columns().PasswordAlgo,
			dao.V2User.Columns().PasswordSalt,
			dao.V2User.Columns().InviteUserId,   //邀请id
			dao.V2User.Columns().TelegramId,     //电报id
			dao.V2User.Columns().Uuid,           //uuid
			dao.V2User.Columns().Token,          //token
			dao.V2User.Columns().CommissionCode, //优惠码
		).
			Where(dao.V2User.Columns().Id, data.Id).
			Update()
	return err
}

// 用户注册
func (s *sUser) RegisterUser(UserName, Passwd, CommissionCode string) error {

	userNameCount, err := s.Cornerstone.GetDB().Where(dao.V2User.Columns().UserName, UserName).Count()
	if err != nil {
		return err
	}
	if userNameCount > 0 {
		return errors.New("用户已经存在！")
	}

	var uu *entity.V2User
	//获取邀请码
	if CommissionCode != "" {
		uu, err = s.GetUserByCommissionCode(CommissionCode)
		if err != nil {
			return err
		}
	} else {
		uu = &entity.V2User{
			Id: 0,
		}
	}

	err = s.AddUser(&entity.V2User{
		UserName:     UserName,
		Password:     Passwd,
		InviteUserId: uu.Id,
	})
	if err != nil {
		return err
	}
	return nil
}

// 添加用户
func (s *sUser) AddUser(user *entity.V2User) error {
	user.PasswordSalt = strings.Split(uuid.New().String(), "-")[0]
	user.Password = utils.MD5V(user.Password, user.PasswordSalt)
	user.PasswordAlgo = "MD5"
	user.CommissionCode = strings.Split(uuid.New().String(), "-")[0]
	user.Uuid = uuid.New().String()
	user.Token = strings.ReplaceAll(uuid.New().String(), "-", "")
	err := s.Cornerstone.Save(user)
	if err != nil {
		return err
	}
	return nil
}

// 删除
func (s *sUser) DelUser(ids []int) error {
	return s.Cornerstone.DelByIds(ids)
}

// 冻结
func (s *sUser) UpUserBanned1(ids []int) (err error) {
	_, err = s.Cornerstone.GetDB().Data(g.Map{dao.V2User.Columns().Banned: 1}).WhereIn(dao.V2User.Columns().Id, ids).Update()
	return
}

// AE设置
func (s *sUser) AEUser(data *entity.V2User) (err error) {
	if data.Id != 0 {
		if data.Password != "" {
			user, err := s.GetUserByIdAndCheck(data.Id)
			if err != nil {
				return err
			}
			data.Password = utils.MD5V(data.Password, user.PasswordSalt)
		}

		err = s.UpUser(data)
		return err
	}

	err = s.AddUser(data)
	return
}

// 获取用户
func (s *sUser) GetUserById(id int) (u *entity.V2User, err error) {
	u = new(entity.V2User)
	err = s.Cornerstone.GetOneById(id, u)
	return u, err
}

// 邀请码获取 邀请用户id
func (s *sUser) GetUserByCommissionCode(commissionCode string) (u *entity.V2User, err error) {
	u = new(entity.V2User)
	err = s.Cornerstone.GetDB().Where(dao.V2User.Columns().CommissionCode, commissionCode).Scan(u)
	if err != nil {
		return
	}
	if u == nil || u.Banned == 1 {
		return u, errors.New("邀请码异常")
	}
	return u, err
}

// 获取用户并且检测用户装
func (s *sUser) GetUserByIdAndCheck(id int) (u *entity.V2User, err error) {
	u, err = s.GetUserById(id)
	if err != nil {
		return
	}
	if u == nil || u.Banned == 1 {
		return nil, errors.New("用户状态异常")
	}
	return
}

// 用户登录
func (s *sUser) Login(userName, passwd string) (user *entity.V2User, err error) {

	//查询用户名
	user, err = s.GetUserByUserName(userName)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}

	passwd = utils.MD5V(passwd, user.PasswordSalt)

	record, err := s.Cornerstone.GetDB().Where(dao.V2User.Columns().UserName, userName).Where(dao.V2User.Columns().Password, passwd).One()
	if err != nil {
		return nil, err
	}
	err = record.Struct(&user)

	return
}

// 获取用户数据
func (s *sUser) GetUserList(req *v1.UserReq, orderBy, orderDirection string, offset, limit int) (items []*model.UserInfo, total int, err error) {

	items = make([]*model.UserInfo, 0)
	gdbU := s.Cornerstone.GetDB()
	if req.V2User.Id != 0 {
		gdbU.Where("id like ?", "%"+strconv.Itoa(req.V2User.Id)+"%")
	}
	gdbU.Where("user_name like ?", "%"+req.V2User.UserName+"%")
	if req.V2User.Banned != -1 {
		gdbU.Where("banned", req.V2User.Banned)
	}
	if req.V2User.GroupId != -1 {
		gdbU.Where("group_id", req.V2User.GroupId)
	}

	if req.US != "" {
		switch req.US {
		case ">=":
			gdbU.Where("u >= ?", g.Slice{req.V2User.U})
		case "<=":
			gdbU.Where("u <= ?", g.Slice{req.V2User.U})
		}
	}

	if req.DS != "" {
		switch req.DS {
		case ">=":
			gdbU.Where("d >= ?", g.Slice{req.V2User.D})
		case "<=":
			gdbU.Where("d <= ?", g.Slice{req.V2User.D})
		}
	}

	if req.TransferEnableS != "" {
		switch req.TransferEnableS {
		case ">=":
			gdbU.Where("transfer_enable >= ?", g.Slice{req.V2User.TransferEnable})
		case "<=":
			gdbU.Where("transfer_enable <= ?", g.Slice{req.V2User.TransferEnable})
		}
	}

	if req.ExpiredAtS != "" {
		switch req.ExpiredAtS {
		case ">=":
			gdbU.Where("expired_at >= ?", g.Slice{req.V2User.ExpiredAt})
		case "<=":
			gdbU.Where("expired_at <= ?", g.Slice{req.V2User.ExpiredAt})
		}
	}

	if req.BalanceS != "" {
		switch req.BalanceS {
		case ">=":
			gdbU.Where("balance >= ?", g.Slice{req.V2User.Balance})
		case "<=":
			gdbU.Where("balance <= ?", g.Slice{req.V2User.ExpiredAt})
		}
	}

	if req.CommissionBalanceS != "" {
		switch req.CommissionBalanceS {
		case ">=":
			gdbU.Where("commission_balance >= ?", g.Slice{req.V2User.CommissionBalance})
		case "<=":
			gdbU.Where("commission_balance <= ?", g.Slice{req.V2User.CommissionBalance})
		}
	}

	if req.CreatedAtS != "" {
		switch req.CreatedAtS {
		case ">=":
			gdbU.Where("created_at >= ?", g.Slice{req.V2User.CreatedAt})
		case "<=":
			gdbU.Where("created_at <= ?", g.Slice{req.V2User.CreatedAt})
		}
	}

	if req.DiscountS != "" {
		switch req.CreatedAtS {
		case ">=":
			gdbU.Where("discount >= ?", g.Slice{req.V2User.Discount})
		case "<=":
			gdbU.Where("discount <= ?", g.Slice{req.V2User.Discount})
		}
	}
	if req.V2User.CommissionType != -1 {
		gdbU.Where("commission_type = ?", g.Slice{req.V2User.CommissionType})

	}
	if req.CommissionRateS != "" {
		switch req.CommissionRateS {
		case ">=":
			gdbU.Where("commission_rate >= ?", g.Slice{req.V2User.CommissionRate})
		case "<=":
			gdbU.Where("commission_rate <= ?", g.Slice{req.V2User.CommissionRate})
		}
	}

	gdbUC := *gdbU
	gdbUCCount := &gdbUC

	err = gdbU.Order(orderBy, orderDirection).Limit(offset, limit).ScanList(&items, "V2User")
	if err != nil {
		return items, 0, err
	}

	total, err = gdbUCCount.Count()
	if err != nil {
		return items, 0, err
	}
	if total > 0 {
		err = dao.V2Plan.Ctx(dao.V2Plan.DB().GetCtx()).
			Where("id", gdb.ListItemValuesUnique(items, "V2User", "GroupId")).
			ScanList(&items, "V2Plan", "V2User", "id:GroupId")
	}

	return items, total, err
}

// 修改密码
func (s *sUser) UpUserPasswdById(req *userv1.UserUpPasswdReq) (res *userv1.UserUpPasswdRes, err error) {
	res = &userv1.UserUpPasswdRes{}
	u, err := s.GetUserByIdAndCheck(req.TUserID)
	if err != nil {
		return res, err
	}
	//检查旧密码
	req.OldPasswd = utils.MD5V(req.OldPasswd, u.PasswordSalt)
	if req.OldPasswd != u.Password {
		return res, errors.New("密码错误，修改失败")
	}

	_, err = s.Cornerstone.GetDB().Data(g.Map{dao.V2User.Columns().Password: req.NewPasswd}).Where(dao.V2User.Columns().Id, req.TUserID).Update()

	return
}

// jwt 处理 ------------------------------------------------------------------
var authService *jwt.GfJWTMiddleware

func Auth() *jwt.GfJWTMiddleware {
	return authService
}

// PayloadFunc is a callback function that will be called during login.
// Using this function it is possible to add additional payload data to the webtoken.
// The data is then made available during requests via c.Get("JWT_PAYLOAD").
// Note that the payload is not encrypted.
// The attributes mentioned on jwt.io can't be used as keys for the map.
// Optional, by default no additional data will be set.
// PayloadFunc是一个回调函数，将在登录期间调用。
// 使用此函数可以向网络令牌添加额外的有效负载数据。
// 然后通过c.Get（“JWT_PAYLOAD”）在请求期间提供数据。
// 请注意，有效载荷未加密。
// jwt.io上提到的属性不能用作贴图的关键点。
// 可选，默认情况下不会设置其他数据。
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(*entity.V2User)
	claims["TUserID"] = params.Id
	claims["TUserName"] = params.UserName
	claims["TPasswd"] = params.Password
	return claims
}

// IdentityHandler get the identity from JWT and set the identity for every request
// Using this function, by r.GetParam("id") get identity
// IdentityHandler从JWT获取标识，并为每个请求设置标识
// 使用此函数，通过r.GetParam（“id”）获取标识
func IdentityHandler(ctx context.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return claims[authService.IdentityKey]
}

// Unauthorized is used to define customized Unauthorized callback function.
// Unauthorized用于定义自定义的Unauthorized回调函数。
func Unauthorized(ctx context.Context, code int, message string) {
	r := g.RequestFromCtx(ctx)
	r.Response.WriteJson(g.Map{
		"code":    code,
		"message": message,
	})
	r.ExitAll()
}

// Authenticator is used to validate login parameters.
// It must return user data as user identifier, it will be stored in Claim Array.
// if your identityKey is 'id', your user data must have 'id'
// Check error (e) to determine the appropriate error message.
// Authenticator用于验证登录参数。
// 它必须将用户数据作为用户标识符返回，它将存储在Claim Array中。
// 如果您的identityKey是“id”，则用户数据必须具有“id”
// 检查错误（e）以确定适当的错误消息。
func Authenticator(ctx context.Context) (interface{}, error) {
	var (
		r = g.RequestFromCtx(ctx)
	)

	if user, err := service.User().Login(r.Get("UserName").String(), r.Get("Passwd").String()); err != nil {
		return user, err
	} else {
		if user != nil {
			return user, nil
		}
	}

	return nil, jwt.ErrFailedAuthentication
}
