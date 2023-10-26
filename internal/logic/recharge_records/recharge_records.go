package recharge_records

import (
	"context"
	"errors"
	"fmt"
	v1 "gov2panel/api/admin/v1"
	"gov2panel/internal/dao"
	d "gov2panel/internal/dao"
	"gov2panel/internal/logic/cornerstone"
	"gov2panel/internal/utils"
	"strconv"
	"time"

	"gov2panel/internal/model/entity"
	"gov2panel/internal/model/model"
	"gov2panel/internal/service"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type sRechargeRecords struct {
	Cornerstone cornerstone.Cornerstone
}

func init() {
	service.RegisterRechargeRecords(New())
}

func New() *sRechargeRecords {
	return &sRechargeRecords{
		Cornerstone: *cornerstone.NewCornerstoneWithTable(dao.V2RechargeRecords.Table()),
	}
}

// 保存数据
// payCode 充值代码，填的充值通道，人工填admin
// val 充值金额|消费金额
// id 充值的支付id|消费订阅id
// couponCode 消费的优惠码，填的优惠码
func (s *sRechargeRecords) SaveRechargeRecords(data *entity.V2RechargeRecords, payCode string, val float64, id int, couponCode string) (err error) {
	user, err := service.User().GetUserById(data.UserId)
	if utils.IgnoreErrNoRows(err) != nil {
		return err
	}
	if user.Id == 0 {
		return errors.New("用户id不存在")
	}
	err = g.DB().Transaction(context.TODO(), func(ctx context.Context, tx gdb.TX) error {
		//为用户充值/消费 金额
		switch data.OperateType {
		case 1: //充值
			data.ConsumptionName = ""
			_, err := tx.Ctx(ctx).Model(d.V2User.Table()).Where(d.V2User.Columns().Id, data.UserId).Increment(d.V2User.Columns().Balance, data.Amount)
			if err != nil {
				return err
			}

		case 2: //消费
			data.TransactionId = utils.UseOrderNo(id, data.Amount, couponCode)
			data.Amount = val
			data.RechargeName = ""
			_, err := tx.Ctx(ctx).Model(d.V2User.Table()).Where(d.V2User.Columns().Id, data.UserId).Decrement(d.V2User.Columns().Balance, val)
			if err != nil {
				return err
			}
		}

		rechargeRecordsId, err := tx.Ctx(ctx).InsertAndGetId(d.V2RechargeRecords.Table(), data)
		if err != nil {
			return err
		}

		//给邀请者添加邀请佣金
		if payCode != "admin" && data.OperateType == 1 { //不是手动添加的 并且 为充值
			// 有邀请者
			if user.InviteUserId != 0 {
				userInviteUser, err := service.User().GetUserById(user.InviteUserId)
				if utils.IgnoreErrNoRows(err) != nil {
					return err
				}
				if userInviteUser.Id != 0 {

					//获取用户的佣金模式和佣金比例
					cType, cRate := service.User().GetUserCTypeAndCRate(userInviteUser)

					//计算邀请者的佣金
					commission, err := service.User().CalculateUserCommission(cType, cRate, data.UserId, data.Amount)
					if err != nil {
						return err
					}
					fmt.Println("-------------------", commission, cType, cRate)

					if commission != 0 {
						//添加佣金
						invitationRecords := &entity.V2InvitationRecords{
							Amount:            commission,
							UserId:            userInviteUser.Id,
							FromUserId:        data.UserId,
							CommissionRate:    cRate,
							RechargeRecordsId: int(rechargeRecordsId),
							OperateType:       1,
							State:             0,
						}
						_, err = tx.Ctx(ctx).Model(d.V2InvitationRecords.Table()).Data(invitationRecords).Insert()
						if err != nil {
							return err
						}
					}

				}
			}

		}
		return nil

	})
	return err
}

// 获取数据
func (s *sRechargeRecords) GetRechargeRecordsList(req *v1.RechargeRecordsReq, orderBy, orderDirection string, offset, limit int) (m []*model.RechargeRecordsInfo, total int, err error) {
	m = make([]*model.RechargeRecordsInfo, 0)
	db := s.Cornerstone.GetDB()

	db.WhereLike(dao.V2RechargeRecords.Columns().RechargeName, "%"+req.V2RechargeRecords.RechargeName+"%")
	db.WhereLike(dao.V2RechargeRecords.Columns().ConsumptionName, "%"+req.V2RechargeRecords.ConsumptionName+"%")
	db.WhereLike(dao.V2RechargeRecords.Columns().TransactionId, "%"+req.TransactionId+"%")
	if req.Id != 0 {
		db.Where(dao.V2RechargeRecords.Columns().Id, req.Id)
	}
	if req.UserId != 0 {
		db.Where(dao.V2RechargeRecords.Columns().UserId, req.V2RechargeRecords.UserId)
	}
	if req.OperateType != 0 {
		db.Where(dao.V2RechargeRecords.Columns().OperateType, req.V2RechargeRecords.OperateType)
	}

	dbC := *db
	dbCCount := &dbC

	err = db.Order(orderBy, orderDirection).Limit(offset, limit).ScanList(&m, "V2RechargeRecords")
	if err != nil {
		return m, 0, err
	}

	total, err = dbCCount.Count()
	if err != nil {
		return m, 0, err
	}

	if total > 0 {
		err = s.Cornerstone.GetDBT(dao.V2User.Table()).
			Where("id", gdb.ListItemValuesUnique(m, "V2RechargeRecords", "UserId")).
			WhereLike(dao.V2User.Columns().UserName, "%"+req.UserName+"%").
			ScanList(&m, "V2User", "V2RechargeRecords", "id:UserId")
	}

	totaljs := 0
	for i := 0; i < len(m); i++ {
		if m[i].V2User == nil {
			m = append(m[:i], m[i+1:]...)
			i--
			totaljs++
		}
	}
	total = total - totaljs

	return m, total, err
}

// 获取数据根据用户id
func (s *sRechargeRecords) GetRechargeRecordsListByUserId(userId int, orderBy, orderDirection string, offset, limit int) (m []*entity.V2RechargeRecords, total int, err error) {
	m = make([]*entity.V2RechargeRecords, 0)
	db := s.Cornerstone.GetDB()
	db.Where(dao.V2RechargeRecords.Columns().UserId, userId)

	dbC := *db
	dbCCount := &dbC

	err = db.Order(orderBy, orderDirection).Limit(offset, limit).Scan(&m)
	if err != nil {
		return m, 0, err
	}

	total, err = dbCCount.Count()
	if err != nil {
		return m, 0, err
	}

	return m, total, err
}

// 更新备注
func (s *sRechargeRecords) UpRechargeRecordsRemarksById(id int, remarks string) (err error) {
	_, err = s.Cornerstone.GetDB().Data(dao.V2RechargeRecords.Columns().Remarks, remarks).Where(dao.V2RechargeRecords.Columns().Id, id).Update()
	return err
}

// 获取当月收入
func (s *sRechargeRecords) GetNowMonthSumAmount() (amount float64, err error) {
	var amountSum *gvar.Var
	timeNow := time.Now()

	sqlStr := fmt.Sprintf("YEAR(%s) = %s and MONTH(%s) = %s",
		dao.V2RechargeRecords.Columns().CreatedAt,
		strconv.Itoa(timeNow.Year()),
		dao.V2RechargeRecords.Columns().CreatedAt,
		strconv.Itoa(int(timeNow.Month())),
	)
	amountSum, err = s.Cornerstone.GetDB().Fields(fmt.Sprintf("SUM(%s)", dao.V2RechargeRecords.Columns().Amount)).Where(sqlStr).Value()
	if err != nil {
		return 0, err
	}

	amount = amountSum.Float64()

	return
}

// 获取当月每一天的收入
func (s *sRechargeRecords) GetNowMonthDaySum() (data []int, err error) {
	data = make([]int, 0)
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
		Fields(fmt.Sprintf("DAY(%s) AS creation_date, sum(%s) AS daily_amount", createAt, dao.V2RechargeRecords.Columns().Amount)).
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
				iDayCount = v["daily_amount"].Int()
			}
		}
		data = append(data, iDayCount)
	}

	return
}
