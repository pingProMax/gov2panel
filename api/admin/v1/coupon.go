package v1

import (
	"gov2panel/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type CouponReq struct {
	g.Meta `path:"/coupon" tags:"Coupon" method:"get,post" summary:"优惠卷管理"`
	entity.V2Coupon
}
type CouponRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Data   []*entity.V2Coupon `json:"data"`
}

type CouponAEReq struct {
	g.Meta `path:"/coupon/ae" tags:"Coupon" method:"post" summary:"优惠卷管理AE"`
	entity.V2Coupon
}
type CouponAERes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type CouponDelReq struct {
	g.Meta `path:"/coupon/del" tags:"Coupon" method:"post" summary:"优惠卷管理删除"`
	Ids    []int `json:"ids"`
}
type CouponDelRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
