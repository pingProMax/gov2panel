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
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type sUser struct {
	Cornerstone cornerstone.Cornerstone
}

func init() {
	service.RegisterUser(New())

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
	if user.CommissionType != 3 {
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
			dao.V2User.Columns().CreatedAt,
		).
			Where(dao.V2User.Columns().Id, data.Id).
			Update(data)
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
			dao.V2User.Columns().ExpiredAt:      time.Now(),
		},
	)

	_, err =
		db.WhereGT(dao.V2User.Columns().GroupId, 0).Where(
			db.Builder().WhereLT(dao.V2User.Columns().ExpiredAt, time.Now()).WhereOr(fmt.Sprintf("%s + %s >= %s", dao.V2User.Columns().U, dao.V2User.Columns().D, dao.V2User.Columns().TransferEnable)),
		).
			Update()
	return err
}

// 用户注册
func (s *sUser) RegisterUser(UserName, Passwd, CommissionCode string) error {

	if utils.CheckStr(UserName) {
		return errors.New("用户名存在特殊字符！")
	}

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
		if err == sql.ErrNoRows {
			return errors.New("邀请码错误")
		}
		if err != nil {
			return err
		}
	} else {
		uu = &entity.V2User{
			Id: 0,
		}
	}

	err = s.AddUser(&entity.V2User{
		UserName:       UserName,
		Password:       Passwd,
		InviteUserId:   uu.Id,
		Banned:         -1,
		IsAdmin:        -1,
		IsStaff:        -1,
		CommissionType: 3,
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
		if err != nil {
			return
		}
		//查询用户更新到上报缓存
		err = service.User().MGetDb2UserMap(data.Id)
		return
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
		Where(dao.V2User.Columns().Banned, -1).
		WhereGT(dao.V2User.Columns().ExpiredAt, time.Now()).Scan(u)
	return u, err
}

// 根据 token 获取用户
func (s *sUser) GetUserByToken(token string) (u *entity.V2User, err error) {
	u = new(entity.V2User)
	err = s.Cornerstone.GetDB().Where(dao.V2User.Columns().Token, token).Scan(u)
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

// 获取用户并且检测用户状态
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
		Where(dao.V2User.Columns().Banned, -1).
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
		Where(dao.V2User.Columns().Banned, -1).
		Wheref("`%s` != ''", dao.V2User.Columns().Uuid).
		Count()
	return
}

// 更新用户 流量使用情况 直接更新数据库 u+值、d+值、t+值
func (s *sUser) UpUserUAndDBy(data []*model.UserTraffic) (err error) {

	colId := dao.V2User.Columns().Id
	colU := dao.V2User.Columns().U
	colD := dao.V2User.Columns().D
	colT := dao.V2User.Columns().T
	sql := fmt.Sprintf("UPDATE `%s` SET ", s.Cornerstone.Table)
	sqlSetU := fmt.Sprintf("`%s` = CASE `%s` ", colU, colId)
	sqlSetD := fmt.Sprintf("`%s` = CASE `%s` ", colD, colId)
	sqlSetPublic := fmt.Sprintf("`%s` = unix_timestamp(now()) ", colT)
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
	sqlSetD = sqlSetD + "END, "

	sqlWhere = fmt.Sprintf("WHERE `%s` IN (%s)", colId, sqlWhere)

	sql = sql + sqlSetU + sqlSetD + sqlSetPublic + sqlWhere
	fmt.Println(sql)
	_, err = g.DB().Exec(gctx.New(), sql)
	if err != nil {
		return
	}

	return

}

// 更新用户 u、d、t
func (s *sUser) UpUserDUTBy(data []*model.UserTraffic) (err error) {

	colId := dao.V2User.Columns().Id
	colU := dao.V2User.Columns().U
	colD := dao.V2User.Columns().D
	colT := dao.V2User.Columns().T
	sql := fmt.Sprintf("UPDATE `%s` SET ", s.Cornerstone.Table)
	sqlSetU := fmt.Sprintf("`%s` = CASE `%s` ", colU, colId)
	sqlSetD := fmt.Sprintf("`%s` = CASE `%s` ", colD, colId)
	sqlSetPublic := fmt.Sprintf("`%s` = unix_timestamp(now()) ", colT)
	sqlWhere := ""

	for i, u := range data {
		sqlSetU = sqlSetU + fmt.Sprintf("WHEN %s THEN %s ", strconv.Itoa(u.UID), strconv.FormatInt(u.Upload, 10))
		sqlSetD = sqlSetD + fmt.Sprintf("WHEN %s THEN %s ", strconv.Itoa(u.UID), strconv.FormatInt(u.Download, 10))
		if i == 0 {
			sqlWhere = strconv.Itoa(u.UID)
		} else {
			sqlWhere = sqlWhere + "," + strconv.Itoa(u.UID)
		}

	}

	sqlSetU = sqlSetU + "END, "
	sqlSetD = sqlSetD + "END, "

	sqlWhere = fmt.Sprintf("WHERE `%s` IN (%s)", colId, sqlWhere)

	sql = sql + sqlSetU + sqlSetD + sqlSetPublic + sqlWhere
	fmt.Println(sql)
	_, err = g.DB().Exec(gctx.New(), sql)
	if err != nil {
		return
	}

	return

}

// 更新用户 7天流量使用数据
func (s *sUser) UpUserDay7Flow(data []*model.UserTraffic) (err error) {
	//用户流量使用缓存

	ctx := gctx.New()
	//服务器当天的流量使用情况 (记录7天的)
	for _, v := range data {

		ketStr := fmt.Sprintf("USER_%s_%s_FLOW_UPLOAD", strconv.Itoa(v.UID), utils.GetDateNowStr())
		userFlow, err := gcache.Get(ctx, ketStr)
		if err != nil {
			return err
		}
		err = utils.GcacheSet(ctx, ketStr, userFlow.Int64()+v.Upload+v.Download, 169*time.Hour)
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

	if user.Banned == 1 {
		return nil, errors.New("账号被冻结请联系管理员")
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
	if req.V2User.Banned != 0 {
		gdbU.Where("banned", req.V2User.Banned)
	}
	if req.V2User.GroupId != 0 {
		gdbU.Where("group_id", req.V2User.GroupId)
	}
	if req.V2User.IsAdmin != 0 {
		gdbU.Where("is_admin", req.V2User.IsAdmin)
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
	if req.V2User.CommissionType != 0 {
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
func (s *sUser) UpUserPasswdById(ctx context.Context, req *userv1.UserUpPasswdReq) (res *userv1.UserUpPasswdRes, err error) {
	res = &userv1.UserUpPasswdRes{}
	u, err := s.GetUserByIdAndCheck(s.GetCtxUser(ctx).Id)
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
		}).Where(dao.V2User.Columns().Id, s.GetCtxUser(ctx).Id).Update()

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

	uuidStr := uuid.New().String()
	_, err = s.Cornerstone.GetDB().Data(
		g.Map{
			dao.V2User.Columns().Token: strings.ReplaceAll(uuid.New().String(), "-", ""),
			dao.V2User.Columns().Uuid:  uuidStr,
		},
	).Where(dao.V2User.Columns().Id, id).Update()
	if err != nil {
		return
	}

	//更新到userMap缓存
	err = s.MGetDb2UserMap(id)

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

// 获取订阅组用户数量
func (s *sUser) GetUserCountByPlanID(id int) (count int, err error) {
	gdbU := s.Cornerstone.GetDB()
	gdbU.Where(dao.V2User.Columns().GroupId, id)
	return gdbU.Count()
}

// 创建 token
func (s *sUser) CreateToken(ctx context.Context, user *entity.V2User) (signedToken string, claims *model.JWTClaims, err error) {
	// Create claims with user information
	claims = &model.JWTClaims{
		UserName: user.UserName,
		TUserID:  user.Id,
		TPasswd:  user.Password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString(g.Cfg().MustGet(gctx.New(), "jwtkey").Bytes())
	if err != nil {
		return "", nil, gerror.NewCode(gcode.CodeInternalError, "Failed to generate token")
	}

	return
}

// 从上下文获取用户信息
func (s *sUser) GetCtxUser(ctx context.Context) *entity.V2User {
	user := &entity.V2User{}
	g.RequestFromCtx(ctx).GetCtxVar("database_user").Struct(user)
	return user
}
