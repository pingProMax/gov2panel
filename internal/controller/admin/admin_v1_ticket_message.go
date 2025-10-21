package admin

import (
	"context"

	v1 "gov2panel/api/admin/v1"
	"gov2panel/internal/service"

	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) TicketMessage(ctx context.Context, req *v1.TicketMessageReq) (res *v1.TicketMessageRes, err error) {
	switch g.RequestFromCtx(ctx).Method {
	case "GET":

	case "POST":
		res = &v1.TicketMessageRes{}
		res.Data, err = service.TicketMessage().GetTicketMessageArrByTicketId(req.TicketId)
		return

	default:
		return
	}
	return
}

func (c *ControllerV1) TicketMessageAdd(ctx context.Context, req *v1.TicketMessageAddReq) (res *v1.TicketMessageAddRes, err error) {
	res = &v1.TicketMessageAddRes{}
	req.V2TicketMessage.UserId = c.getUser(ctx).Id
	req.V2TicketMessage.Message = ghtml.Entities(req.V2TicketMessage.Message)
	err = service.TicketMessage().SaveTicketMessageAdmin(&req.V2TicketMessage)
	return
}
