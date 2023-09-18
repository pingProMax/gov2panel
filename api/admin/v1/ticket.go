package v1

import (
	"gov2panel/internal/model/entity"
	"gov2panel/internal/model/model"

	"github.com/gogf/gf/v2/frame/g"
)

type TicketReq struct {
	g.Meta `path:"/ticket" tags:"Ticket" method:"get,post" summary:"工单管理"`
	SortOrder
	OffsetLimit
	entity.V2Ticket
	UserName string `json:"user_name"`
}
type TicketRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Data   []*model.TicketInfo `json:"data"`
	Totle  int                 `json:"totle"`
}

type TicketAEReq struct {
	g.Meta `path:"/ticket/ae" tags:"Ticket" method:"post" summary:"工单管理AE"`
	entity.V2Ticket
}
type TicketAERes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type TicketDelReq struct {
	g.Meta `path:"/ticket/del" tags:"Ticket" method:"post" summary:"工单管理删除"`
	Ids    []int `json:"ids"`
}
type TicketDelRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type TicketCloseReq struct {
	g.Meta `path:"/ticket/close" tags:"Ticket" method:"post" summary:"工单管理关闭"`
	Ids    []int `json:"ids"`
}
type TicketCloseRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
