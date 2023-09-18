package v1

import (
	"gov2panel/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type PlanReq struct {
	g.Meta `path:"/plan" tags:"Plan" method:"get,post" summary:"订阅购买"`
	entity.V2Plan
}
type PlanRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Data   []*entity.V2Plan `json:"data"`
}

type Plan2Req struct {
	g.Meta `path:"/plan2" tags:"Plan" method:"get" summary:"订阅购买页面"`
	Id     int `json:"id"`
}
type Plan2Res struct {
	g.Meta `mime:"text/html" example:"string"`
	Data   *entity.V2Plan `json:"data"`
}

type BuyReq struct {
	g.Meta  `path:"/buy" tags:"Buy" method:"post" summary:"购买"`
	PlanId  int    `json:"plan_id"` //订阅id
	Code    string `json:"code"`    //优惠码
	TUserID int    //用户id
}
type BuyRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
