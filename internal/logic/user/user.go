package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	v1 "gov2panel/api/admin/v1"
	userv1 "gov2panel/api/user/v1"
	"gov2panel/internal/dao"
	"gov2panel/internal/logic/cornerstone"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/model/model"
	"gov2panel/internal/service"
	"gov2panel/internal/utils"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/google/uuid"
)

type sUser struct {
	Cornerstone cornerstone.Cornerstone
}

func init() {
	service.RegisterUser(New())
	jwtkey, err := g.Cfg().Get(gctx.New(), "jwtkey")
	if err != nil {
		panic(err.Error())
	}
	auth := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "gov2panel",
		Key:             jwtkey.Bytes(),
		Timeout:         time.Hour * 24,
		MaxRefresh:      time.Hour * 24,
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
	err = s.Cornerstone.GetDB().Where(dao.V2User.Columns().UserName, userName).Scan(&user)
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
		db.FieldsEx(dao.V2User.Columns().PasswordAlgo)
		db.FieldsEx(dao.V2User.Columns().PasswordSalt)
	}

	passwordSalt := strings.Split(uuid.New().String(), "-")[0]
	data.PasswordAlgo = "MD5"
	data.Password = utils.MD5V(data.Password, passwordSalt)
	data.PasswordSalt = passwordSalt

	_, err =
		db.FieldsEx(
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

// 更新过期用户的权限组和流量
func (s *sUser) ClearExpiredUserGroupIdAndUDTransferEnable() (err error) {

	db := s.Cornerstone.GetDB().Data(
		g.Map{
			dao.V2User.Columns().GroupId:        0,
			dao.V2User.Columns().U:              0,
			dao.V2User.Columns().D:              0,
			dao.V2User.Columns().TransferEnable: 0,
		},
	)

	_, err =
		db.WhereLT(dao.V2User.Columns().ExpiredAt, time.Now()).
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
	data.U = utils.GBToBytes(float64(data.U))
	data.D = utils.GBToBytes(float64(data.D))
	data.TransferEnable = utils.GBToBytes(float64(data.TransferEnable))
	if data.Id != 0 {
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

// 获取用户
func (s *sUser) GetUserByTokenAndUDAndGTExpiredAt(token string) (u *entity.V2User, err error) {
	u = new(entity.V2User)
	err = s.Cornerstone.GetDB().Where(dao.V2User.Columns().Token, token).
		Where(fmt.Sprintf("%s > %s + %s", dao.V2User.Columns().TransferEnable, dao.V2User.Columns().U, dao.V2User.Columns().D)).
		WhereGT(dao.V2User.Columns().ExpiredAt, time.Now()).Scan(u)
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

// 获取用户 订阅组下的用户数据
func (s *sUser) GetUserListByGroupIds(groupIds []int) (u []*entity.V2User, err error) {
	u = make([]*entity.V2User, 0)
	err = s.Cornerstone.GetDB().
		Where(dao.V2User.Columns().GroupId, groupIds).
		Where(fmt.Sprintf("%s > %s + %s", dao.V2User.Columns().TransferEnable, dao.V2User.Columns().U, dao.V2User.Columns().D)).
		WhereGT(dao.V2User.Columns().ExpiredAt, time.Now()).
		Where(dao.V2User.Columns().Banned, 0).
		Wheref("`%s` != ''", dao.V2User.Columns().Uuid).
		Scan(&u)
	return
}

// 获取用户数量 订阅组下的用户数据
func (s *sUser) GetUserCountByGroupIds(groupIds []int) (totle int, err error) {
	totle, err = s.Cornerstone.GetDB().
		Where(dao.V2User.Columns().GroupId, groupIds).
		Where(fmt.Sprintf("%s > %s + %s", dao.V2User.Columns().TransferEnable, dao.V2User.Columns().U, dao.V2User.Columns().D)).
		WhereGT(dao.V2User.Columns().ExpiredAt, time.Now()).
		Where(dao.V2User.Columns().Banned, 0).
		Wheref("`%s` != ''", dao.V2User.Columns().Uuid).
		Count()
	return
}

// 更新用户 流量使用情况
func (s *sUser) UpUserUAndDBy(data []model.UserTraffic) (err error) {
	colId := dao.V2User.Columns().Id
	colU := dao.V2User.Columns().U
	colD := dao.V2User.Columns().D
	sql := fmt.Sprintf("UPDATE `%s` SET ", s.Cornerstone.Table)
	sqlSetU := fmt.Sprintf("`%s` = CASE `%s` ", colU, colId)
	sqlSetD := fmt.Sprintf("`%s` = CASE `%s` ", colD, colId)
	sqlWhere := ""

	for i, u := range data {
		sqlSetU = sqlSetU + fmt.Sprintf("WHEN %s THEN `%s`+%s ", strconv.Itoa(u.UID), colU, strconv.FormatInt(u.Upload, 10))
		sqlSetD = sqlSetD + fmt.Sprintf("WHEN %s THEN `%s`+%s ", strconv.Itoa(u.UID), colD, strconv.FormatInt(u.Download, 10))

		if i == 0 {
			sqlWhere = strconv.Itoa(u.UID)
		} else {
			sqlWhere = sqlWhere + "," + strconv.Itoa(u.UID)
		}

	}

	sqlSetU = sqlSetU + "END, "
	sqlSetD = sqlSetD + "END "

	sqlWhere = fmt.Sprintf("WHERE `%s` IN (%s)", colId, sqlWhere)

	sql = sql + sqlSetU + sqlSetD + sqlWhere
	fmt.Println(sql)
	_, err = g.DB().Exec(gctx.New(), sql)
	if err != nil {
		return
	}

	//用户流量使用缓存

	ctx := gctx.New()
	//服务器当天的流量使用情况 (记录7天的)
	for _, v := range data {

		ketStr := fmt.Sprintf("USER_%s_%s_FLOW_UPLOAD", strconv.Itoa(v.UID), utils.GetDateNowStr())
		userFlow, err := gcache.Get(ctx, ketStr)
		if err != nil {
			return err
		}

		err = gcache.Set(ctx, ketStr, userFlow.Int64()+v.Upload+v.Download, 169*time.Hour)
		if err != nil {
			return err
		}
	}

	return

}

// 用户登录
func (s *sUser) Login(userName, passwd string) (user *entity.V2User, err error) {

	//查询用户名
	user, err = s.GetUserByUserName(userName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("账号或密码错误")
		}
		return nil, err
	}

	if user == nil {
		return nil, errors.New("账号或密码错误")
	}

	switch user.PasswordAlgo {
	case "MD5":
		passwd = utils.MD5V(passwd, user.PasswordSalt)

		err = s.Cornerstone.GetDB().Where(dao.V2User.Columns().UserName, userName).Where(dao.V2User.Columns().Password, passwd).Scan(&user)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.New("账号或密码错误")
			}
			return nil, err
		}
		return

	case "BCRYPT":
		if !utils.BcryptCheckPassword(passwd, user.Password) {
			return nil, errors.New("账号或密码错误")
		}
		return
	default:
		return nil, errors.New("账号异常请联系管理员！！！")

	}

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
	switch u.PasswordAlgo {
	case "MD5":
		req.OldPasswd = utils.MD5V(req.OldPasswd, u.PasswordSalt)
		if req.OldPasswd != u.Password {
			return res, errors.New("密码错误，修改失败")
		}
	case "BCRYPT":
		if utils.BcryptCheckPassword(req.OldPasswd, u.PasswordSalt) {
			return res, errors.New("密码错误，修改失败")
		}
	default:
		err = errors.New("账号密码异常，请联系管理员")
		return
	}

	passwordSalt := strings.Split(uuid.New().String(), "-")[0]
	_, err = s.Cornerstone.GetDB().Data(
		g.Map{
			dao.V2User.Columns().PasswordAlgo: "MD5",
			dao.V2User.Columns().Password:     utils.MD5V(req.NewPasswd, passwordSalt),
			dao.V2User.Columns().PasswordSalt: passwordSalt,
		}).Where(dao.V2User.Columns().Id, req.TUserID).Update()

	return
}

// 获取当月注册量
func (s *sUser) GetNowMonthCount() (count int, err error) {
	timeNow := time.Now()

	sqlStr := fmt.Sprintf("YEAR(%s) = %s and MONTH(%s) = %s",
		dao.V2RechargeRecords.Columns().CreatedAt,
		strconv.Itoa(timeNow.Year()),
		dao.V2RechargeRecords.Columns().CreatedAt,
		strconv.Itoa(int(timeNow.Month())),
	)
	count, err = s.Cornerstone.GetDB().Where(sqlStr).Count()

	return
}

// 重置用户的Token和uuid
func (s *sUser) ResetTokenAndUuidById(id int) (err error) {

	_, err = s.Cornerstone.GetDB().Data(
		g.Map{
			dao.V2User.Columns().Token: strings.ReplaceAll(uuid.New().String(), "-", ""),
			dao.V2User.Columns().Uuid:  uuid.New().String(),
		},
	).Where(dao.V2User.Columns().Id, id).Update()

	return
}

// 获取当月每一天注册量
func (s *sUser) GetNowMonthDayCount() (count []int, err error) {
	count = make([]int, 0)
	timeNow := time.Now()
	createAt := dao.V2RechargeRecords.Columns().CreatedAt

	sqlStr := fmt.Sprintf("YEAR(%s) = %s and MONTH(%s) = %s and (",
		createAt,
		strconv.Itoa(timeNow.Year()),
		createAt,
		strconv.Itoa(int(timeNow.Month())),
	)

	for i := timeNow.Day(); i > 0; i-- {
		sqlStr = sqlStr + fmt.Sprintf("DAY(%s) = %s ", createAt, strconv.Itoa(i))
		if i != 1 {
			sqlStr = sqlStr + "or "
		}
	}

	sqlStr = sqlStr + ")"

	result, err := s.Cornerstone.GetDB().
		Fields(fmt.Sprintf("DAY(%s) AS creation_date, COUNT(*) AS daily_count", createAt)).
		Where(sqlStr).
		Group(fmt.Sprintf("DAY(%s)", createAt)).
		OrderAsc("creation_date").All()
	if err != nil {
		return
	}

	for i := 1; i <= timeNow.Day(); i++ {

		var iDayCount int
		for _, v := range result {
			if v["creation_date"].Int() == i {
				iDayCount = v["daily_count"].Int()
			}
		}
		count = append(count, iDayCount)
	}

	return
}

func (s *sUser) Logout(ctx context.Context) {
	authService.LogoutHandler(ctx)
}

func (s *sUser) Refresh(ctx context.Context) (tokenString string, expire time.Time) {
	return authService.RefreshHandler(ctx)
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
	if r.RequestURI != "/login" {
		r.Response.RedirectTo("/login", http.StatusFound)
	} else {
		r.Response.WriteJson(g.Map{
			"code":    code,
			"message": message,
		})
	}

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
