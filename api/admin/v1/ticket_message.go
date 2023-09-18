package v1

import (
	"gov2panel/internal/model/entity"
	"gov2panel/internal/model/model"

	"github.com/gogf/gf/v2/frame/g"
)

type TicketMessageReq struct {
	g.Meta `path:"/ticket_msg" tags:"TicketMessage" method:"get,post" summary:"工单信息获取"`
	entity.V2TicketMessage
}
type TicketMessageRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Data   []*model.TicketMessageInfo `json:"data"`
}

type TicketMessageAddReq struct {
	g.Meta `path:"/ticket_msg/add" tags:"TicketMessageAdd" method:"post" summary:"工单信息回复获取"`
	entity.V2TicketMessage
}
type TicketMessageAddRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
