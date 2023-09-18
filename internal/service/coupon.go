// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	userv1 "gov2panel/api/user/v1"
	"gov2panel/internal/model/entity"
)

type (
	ICoupon interface {
		// AE设置
		AECoupon(data *entity.V2Coupon) (err error)
		// 删除
		DelCoupon(ids []int) error
		// 获取所有
		GetCouponAllList(req entity.V2Coupon) (m []*entity.V2Coupon, err error)
		// 根据code 获取
		GetCouponByCode(code string) (d *entity.V2Coupon, err error)
		// 优惠码是否可用
		CheckCouponCanUseByCode(req *userv1.CouponReq) (res *userv1.CouponRes, err error)
	}
)

var (
	localCoupon ICoupon
)

func Coupon() ICoupon {
	if localCoupon == nil {
		panic("implement not found for interface ICoupon, forgot register?")
	}
	return localCoupon
}

func RegisterCoupon(i ICoupon) {
	localCoupon = i
}
