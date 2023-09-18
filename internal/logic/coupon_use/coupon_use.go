package CouponUse

import (
	"gov2panel/internal/dao"
	"gov2panel/internal/logic/cornerstone"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/service"
)

type sCouponUse struct {
	Cornerstone cornerstone.Cornerstone
}

func init() {
	service.RegisterCouponUse(New())
}

func New() *sCouponUse {
	return &sCouponUse{
		Cornerstone: *cornerstone.NewCornerstoneWithTable(dao.V2CouponUse.Table()),
	}
}

// 根据id获取一条数据
func (s *sCouponUse) GetCouponUseById(id int) (d *entity.V2CouponUse, err error) {
	err = s.Cornerstone.GetDB().Where(dao.V2CouponUse.Columns().Id, id).Scan(&d)
	return
}

// 根据user_id和coupon_id获取数据
func (s *sCouponUse) GetCouponUseByUserIdAndCouponId(userId int, couponId int) (d []*entity.V2CouponUse, err error) {
	d = make([]*entity.V2CouponUse, 0)
	err = s.Cornerstone.GetDB().Where(dao.V2CouponUse.Columns().UserId, userId).Where(dao.V2CouponUse.Columns().CouponId, couponId).Scan(&d)
	return
}

// 根据coupon_id获取数据
func (s *sCouponUse) GetCouponUseByCouponId(couponId int) (d []*entity.V2CouponUse, err error) {
	d = make([]*entity.V2CouponUse, 0)
	err = s.Cornerstone.GetDB().Where(dao.V2CouponUse.Columns().CouponId, couponId).Scan(&d)
	return
}
