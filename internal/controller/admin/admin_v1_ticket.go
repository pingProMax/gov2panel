package admin

import (
	"context"
	v1 "gov2panel/api/admin/v1"
	"gov2panel/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Ticket(ctx context.Context, req *v1.TicketReq) (res *v1.TicketRes, err error) {
	switch g.RequestFromCtx(ctx).Method {
	case "GET":
		setTplAdmin(ctx, "ticket", nil)
	case "POST":
		res = &v1.TicketRes{}
		res.Data, res.Totle, err = service.Ticket().GetTicketList(&req.V2Ticket, req.UserName, req.Sort, req.Order, req.Offset, req.Limit)
		return

	default:
		return
	}
	return
}

func (c *ControllerV1) TicketAE(ctx context.Context, req *v1.TicketAEReq) (res *v1.TicketAERes, err error) {
	err = service.Ticket().AETicket(&req.V2Ticket)
	if err != nil {
		return res, err
	}
	return res, err
}

func (c *ControllerV1) TicketDel(ctx context.Context, req *v1.TicketDelReq) (res *v1.TicketDelRes, err error) {
	err = service.Ticket().DelTicket(req.Ids)
	if err != nil {
		return res, err
	}
	return res, err
}

func (c *ControllerV1) TicketClose(ctx context.Context, req *v1.TicketCloseReq) (res *v1.TicketCloseRes, err error) {
	err = service.Ticket().CloseTicket(req.Ids)
	if err != nil {
		return res, err
	}
	return res, err
}
