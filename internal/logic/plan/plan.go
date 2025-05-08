package plan

import (
	"context"
	"errors"
	"fmt"
	userv1 "gov2panel/api/user/v1"
	"gov2panel/internal/dao"
	d "gov2panel/internal/dao"
	"gov2panel/internal/logic/cornerstone"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/model/model"
	"gov2panel/internal/service"
	"gov2panel/internal/utils"
	"strconv"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type sPlan struct {
	Cornerstone cornerstone.Cornerstone
}

func init() {
	service.RegisterPlan(New())
}

func New() *sPlan {
	return &sPlan{
		Cornerstone: *cornerstone.NewCornerstoneWithTable(dao.V2Plan.Table()),
	}
}

// AE设置
func (s *sPlan) AEPlan(data *entity.V2Plan) (err error) {
	if data.Id != 0 {
		err = s.Cornerstone.UpdateById(data.Id, data)
		return
	}

	err = s.Cornerstone.Save(data)
	return
}

// 删除
func (s *sPlan) DelPlan(ids []int) error {
	serverCount, err := service.ProxyService().GetServiceCountByPlanId(ids)
	if err != nil {
		return err
	}

	if serverCount > 0 {
		return errors.New("订阅有节点在使用，无法删除！")
	}

	userCount, err := service.User().GetUserCountByGroupIds(ids)
	if err != nil {
		return err
	}

	if userCount > 0 {
		return errors.New("订阅有用户在使用，无法删除！")
	}

	return s.Cornerstone.DelByIds(ids)
}

// 获取所有
func (s *sPlan) GetPlanAllList(req entity.V2Plan) (m []*entity.V2Plan, err error) {
	m = make([]*entity.V2Plan, 0)
	err = s.Cornerstone.GetDB().
		OmitEmpty().
		Where(dao.V2Plan.Columns().Id, req.Id).
		WhereLike(dao.V2Plan.Columns().Name, "%"+req.Name+"%").
		OrderDesc("order_id").Scan(&m)
	return m, err
}

// 获取显示的订阅
func (s *sPlan) GetPlanShowList() (m []*entity.V2Plan, err error) {
	m = make([]*entity.V2Plan, 0)
	err = s.Cornerstone.GetDB().
		Where(dao.V2Plan.Columns().Show, 1).
		OrderDesc("order_id").Scan(&m)
	return m, err
}

// 获取显示的订阅 可覆盖的
func (s *sPlan) GetPlanShowAndResetTrafficMethod1List() (m []*entity.V2Plan, err error) {
	m = make([]*entity.V2Plan, 0)
	err = s.Cornerstone.GetDB().
		Where(dao.V2Plan.Columns().Show, 1).
		Where(dao.V2Plan.Columns().ResetTrafficMethod, 1).
		OrderDesc("order_id").Scan(&m)
	return m, err
}

// 获取可覆盖的订阅
func (s *sPlan) GetPlanResetTrafficMethod1List() (m []*entity.V2Plan, err error) {
	m = make([]*entity.V2Plan, 0)
	err = s.Cornerstone.GetDB().
		Where(dao.V2Plan.Columns().ResetTrafficMethod, 1).
		OrderDesc("order_id").Scan(&m)
	return m, err
}

// 根据id获取
func (s *sPlan) GetPlanById(id int) (d *entity.V2Plan, err error) {
	d = new(entity.V2Plan)
	err = s.Cornerstone.GetOneById(id, d)
	return
}

// 用户购买/续费套餐处理
func (s *sPlan) UserBuyAndRenew(code string, plan *entity.V2Plan, user *entity.V2User) (err error) {

	//套餐最终价格
	var price float64
	price = plan.Price
	var couponRes *userv1.CouponRes

	//检查用户是否有专享折扣
	if user.Discount > 0 {
		price = price - (price * user.Discount / 100)
	}

	//检查套餐当前用户数量
	if plan.CapacityLimit > 0 && user.GroupId != plan.Id {
		planUserCoun, err := service.User().GetUserCountByPlanID(plan.Id)
		if err != nil {
			return err
		}

		if planUserCoun >= plan.CapacityLimit {
			return errors.New("当前订阅人数达到上限！")
		}
	}

	if code != "" {
		//检查优惠码
		couponRes, err = service.Coupon().CheckCouponCanUseByCode(&userv1.CouponReq{Code: code, PlanId: plan.Id, TUserID: user.Id})
		if err != nil {
			return err
		}

		switch couponRes.Data.Type {
		case 1: //金额优惠
			price = (price - couponRes.Data.Value)
		case 2: //比例优惠
			price = price - (price * couponRes.Data.Value / 100)
		}
	}

	if price < 0 {
		price = 0
	}

	if user.Balance < price {
		return errors.New("余额不足，请去钱包充值")
	}

	//扣款 和 设置用户套餐
	g.DB().Transaction(context.TODO(), func(ctx context.Context, tx gdb.TX) error {

		// 扣款
		err = service.RechargeRecords().SaveRechargeRecords(
			&entity.V2RechargeRecords{
				Amount:          plan.Price,
				UserId:          user.Id,
				OperateType:     2,
				ConsumptionName: plan.Name,
			},
			"",
			price,
			plan.Id,
			code,
		)
		if err != nil {
			return err
		}

		//添加优惠码使用记录
		if code != "" {
			_, err := tx.Ctx(ctx).Insert(d.V2CouponUse.Table(), g.Map{
				d.V2CouponUse.Columns().CouponId: couponRes.Data.Id,
				d.V2CouponUse.Columns().UserId:   user.Id,
				d.V2CouponUse.Columns().PlanId:   plan.Id,
			})
			if err != nil {
				return err
			}
		}

		//为用户添加套餐 流量 过期时间 等
		var userUpData g.Map
		switch plan.ResetTrafficMethod { //套餐类型，1 覆盖、2 叠加
		case 1:
			userUpData = g.Map{
				d.V2User.Columns().GroupId:        plan.Id,
				d.V2User.Columns().U:              0,
				d.V2User.Columns().D:              0,
				d.V2User.Columns().TransferEnable: utils.GBToBytes(plan.TransferEnable),
				d.V2User.Columns().ExpiredAt:      time.Now().Add(time.Duration(plan.Expired) * 24 * time.Hour),
			}
		case 2:
			userUpData = g.Map{
				d.V2User.Columns().TransferEnable: gdb.Raw(d.V2User.Columns().TransferEnable + "+" + fmt.Sprintf("%v", utils.GBToBytes(plan.TransferEnable))),
				d.V2User.Columns().ExpiredAt:      gdb.Raw(fmt.Sprintf("DATE_ADD(%s, INTERVAL %s DAY)", d.V2User.Columns().ExpiredAt, strconv.Itoa(plan.Expired))),
			}
		}
		_, err = tx.Ctx(ctx).Model(d.V2User.Table()).Data(userUpData).Where(d.V2User.Columns().Id, user.Id).Update()
		if err != nil {
			return err
		}

		return nil
	})

	//查询用户更新到上报缓存
	user, _ = service.User().GetUserById(user.Id)
	service.User().MUpUserMap(model.UserToUserTraffic(user))

	return
}
