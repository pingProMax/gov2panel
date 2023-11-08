package v1

import (
	"gov2panel/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type PlanReq struct {
	g.Meta `path:"/plan" tags:"Plan" method:"get,post" summary:"订阅权限管理"`
	entity.V2Plan
}
type PlanRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Data   []*entity.V2Plan `json:"data"`
}

type PlanAEReq struct {
	g.Meta `path:"/plan/ae" tags:"Plan" method:"post" summary:"订阅权限管理AE"`
	entity.V2Plan
}
type PlanAERes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type PlanDelReq struct {
	g.Meta `path:"/plan/del" tags:"Plan" method:"post" summary:"订阅权限管理删除"`
	Ids    []int `json:"ids"`
}
type PlanDelRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type PlanGetShowReq struct {
	g.Meta `path:"/plan/get_show" tags:"Plan" method:"post" summary:"获取显示的订阅列表"`
}
type PlanGetShowRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Data   []*entity.V2Plan `json:"data"`
}

type PlanGetResetTrafficMethod1Req struct {
	g.Meta `path:"/plan/get_rtm1" tags:"Plan" method:"post" summary:"获取 可覆盖的订阅列表"`
}
type PlanGetResetTrafficMethod1Res struct {
	g.Meta `mime:"text/html" example:"string"`
	Data   []*entity.V2Plan `json:"data"`
}
