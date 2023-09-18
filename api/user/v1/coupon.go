package v1

import (
	"gov2panel/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type CouponReq struct {
	g.Meta  `path:"/coupon" tags:"Coupon" method:"post" summary:"优惠卷信息获取"`
	Code    string `json:"code"`
	PlanId  int    `json:"plan_id"`
	TUserID int
}
type CouponRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Data   *entity.V2Coupon `json:"data"`
}
