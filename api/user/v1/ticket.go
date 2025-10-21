package v1

import (
	"gov2panel/internal/model/entity"
	"gov2panel/internal/model/model"

	"github.com/gogf/gf/v2/frame/g"
)

type TicketReq struct {
	g.Meta `path:"/ticket" tags:"Ticket" method:"get,post" summary:"工单"`
	entity.V2Ticket
	OffsetLimit
}
type TicketRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Data   []*model.TicketInfo `json:"data"`
	Totle  int                 `json:"totle"`
}

type TicketCloseReq struct {
	g.Meta `path:"/ticket/close" tags:"Ticket" method:"post" summary:"工单关闭"`
	Ids    []int `json:"ids"`
}
type TicketCloseRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type TicketCreateReq struct {
	g.Meta `path:"/ticket/create" tags:"Ticket" method:"post" summary:"创建工单"`
	entity.V2Ticket
}
type TicketCreateRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type TicketMessageReq struct {
	g.Meta `path:"/ticket/ticket_msg" tags:"TicketMessage" method:"get,post" summary:"工单信息获取"`
	entity.V2TicketMessage
}
type TicketMessageRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Data   []*model.TicketMessageInfo `json:"data"`
}

type TicketMessageAddReq struct {
	g.Meta `path:"/ticket/ticket_msg/add" tags:"TicketMessageAdd" method:"post" summary:"工单信息回复获取"`
	entity.V2TicketMessage
}
type TicketMessageAddRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
