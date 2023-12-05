package coupon

import (
	"errors"
	userv1 "gov2panel/api/user/v1"
	"gov2panel/internal/dao"
	"gov2panel/internal/logic/cornerstone"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/service"
	"time"
)

type sCoupon struct {
	Cornerstone cornerstone.Cornerstone
}

func init() {
	service.RegisterCoupon(New())
}

func New() *sCoupon {
	return &sCoupon{
		Cornerstone: *cornerstone.NewCornerstoneWithTable(dao.V2Coupon.Table()),
	}
}

// AE设置
func (s *sCoupon) AECoupon(data *entity.V2Coupon) (err error) {
	if data.Id != 0 {
		err = s.Cornerstone.UpdateById(data.Id, data)
		return
	}

	err = s.Cornerstone.Save(data)
	return
}

// 删除
func (s *sCoupon) DelCoupon(ids []int) error {
	return s.Cornerstone.DelByIds(ids)
}

// 获取所有
func (s *sCoupon) GetCouponAllList(req entity.V2Coupon) (m []*entity.V2Coupon, err error) {
	m = make([]*entity.V2Coupon, 0)

	err = s.Cornerstone.GetDB().
		OmitEmpty().
		Where(dao.V2Coupon.Columns().Id, req.Id).
		WhereLike(dao.V2Coupon.Columns().Name, "%"+req.Name+"%").
		WhereLike(dao.V2Coupon.Columns().Code, "%"+req.Code+"%").
		Scan(&m)
	return m, err
}

// 根据code 获取
func (s *sCoupon) GetCouponByCode(code string) (d *entity.V2Coupon, err error) {
	err = s.Cornerstone.GetDB().Where(dao.V2Coupon.Columns().Code, code).Scan(&d)
	return
}

// 优惠码是否可用
func (s *sCoupon) CheckCouponCanUseByCode(req *userv1.CouponReq) (res *userv1.CouponRes, err error) {
	res = &userv1.CouponRes{}

	coupon, err := s.GetCouponByCode(req.Code)
	if err != nil {
		return res, err
	}
	if coupon == nil {
		return res, errors.New("优惠码不存在")
	}
	nowTimeUnix := time.Now().Unix()

	if coupon.StartedAt.Timestamp() <= nowTimeUnix && coupon.EndedAt.Timestamp() >= nowTimeUnix && coupon.Enable == 1 {

		//判断是否指定订阅
		if coupon.LimitPlanId != 0 {
			if coupon.LimitPlanId != req.PlanId {
				return res, errors.New("优惠码不能应用到此套餐")
			}
		}

		//判断每个用户可以使用次数
		if coupon.LimitUse != -1 {
			couponUseList, err := service.CouponUse().GetCouponUseByUserIdAndCouponId(req.TUserID, coupon.Id)
			if err != nil {
				return res, err
			}
			if len(couponUseList) >= coupon.LimitUse {
				return res, errors.New("此优惠码您已经使用过")
			}
		}

		//判断最大使用情况
		if coupon.LimitUseWithUser != -1 {
			couponWithUseList, err := service.CouponUse().GetCouponUseByCouponId(coupon.Id)
			if err != nil {
				return res, err
			}
			if len(couponWithUseList) >= coupon.LimitUseWithUser {
				return res, errors.New("此优惠码已被用完")
			}
		}

		res.Data = coupon
		return
	} else {
		return res, errors.New("优惠码未开始或已过期")
	}
}
