// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"gov2panel/internal/model/entity"
)

type (
	ICouponUse interface {
		// 根据id获取一条数据
		GetCouponUseById(id int) (d *entity.V2CouponUse, err error)
		// 根据user_id和coupon_id获取数据
		GetCouponUseByUserIdAndCouponId(userId int, couponId int) (d []*entity.V2CouponUse, err error)
		// 根据coupon_id获取数据
		GetCouponUseByCouponId(couponId int) (d []*entity.V2CouponUse, err error)
	}
)

var (
	localCouponUse ICouponUse
)

func CouponUse() ICouponUse {
	if localCouponUse == nil {
		panic("implement not found for interface ICouponUse, forgot register?")
	}
	return localCouponUse
}

func RegisterCouponUse(i ICouponUse) {
	localCouponUse = i
}
