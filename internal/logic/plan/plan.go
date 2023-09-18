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

// 删除
func (s *sPlan) GetPlanById(id int) (d *entity.V2Plan, err error) {
	d = new(entity.V2Plan)
	err = s.Cornerstone.GetOneById(id, d)
	return
}

// 用户购买套餐处理
func (s *sPlan) UserBuy(req *userv1.BuyReq) (res *userv1.BuyRes, err error) {
	res = &userv1.BuyRes{}

	//检查套餐
	plan, err := s.GetPlanById(req.PlanId)
	if err != nil {
		return
	}
	if plan == nil {
		return res, errors.New("套餐不存在")
	}
	if plan.Show != 1 {
		return res, errors.New("套餐未开启")
	}
	if plan.Price < 0 || plan.Expired < 0 {
		return res, errors.New("套餐设置不对请联系管理员")
	}

	//套餐最终价格
	var price float64
	price = plan.Price
	var couponRes *userv1.CouponRes

	u, err := service.User().GetUserByIdAndCheck(req.TUserID)
	if err != nil {
		return res, err
	}

	//检查用户是否有专享折扣
	if u.Discount > 0 {
		price = price - (price * u.Discount / 100)
	}

	if req.Code != "" {
		//检查优惠码
		couponRes, err = service.Coupon().CheckCouponCanUseByCode(&userv1.CouponReq{Code: req.Code, PlanId: req.PlanId, TUserID: req.TUserID})
		if err != nil {
			return res, err
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

	if u.Balance < price {
		return res, errors.New("余额不足")
	}

	//扣款 和 设置用户套餐
	g.DB().Transaction(context.TODO(), func(ctx context.Context, tx gdb.TX) error {

		// 扣款
		err = service.RechargeRecords().SaveRechargeRecords(
			&entity.V2RechargeRecords{
				Amount:          plan.Price,
				UserId:          u.Id,
				OperateType:     2,
				ConsumptionName: plan.Name,
			},
			"",
			price,
			plan.Id,
			req.Code,
		)
		if err != nil {
			return err
		}

		//添加优惠码使用记录
		if req.Code != "" {
			_, err := tx.Ctx(ctx).Insert(d.V2CouponUse.Table(), g.Map{
				d.V2CouponUse.Columns().CouponId: couponRes.Data.Id,
				d.V2CouponUse.Columns().UserId:   u.Id,
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
		_, err = tx.Ctx(ctx).Model(d.V2User.Table()).Data(userUpData).Where(d.V2User.Columns().Id, u.Id).Update()
		if err != nil {
			return err
		}

		return nil
	})

	return
}
