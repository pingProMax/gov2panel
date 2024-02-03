package user

import (
	"context"
	"errors"
	"strings"

	v1 "gov2panel/api/user/v1"
	"gov2panel/internal/service"

	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Ticket(ctx context.Context, req *v1.TicketReq) (res *v1.TicketRes, err error) {
	switch g.RequestFromCtx(ctx).Method {
	case "GET":
		setTplUser(ctx, "ticket", nil)
	case "POST":
		res = &v1.TicketRes{}
		req.UserId = req.TUserID
		req.ReplyStatus = 0
		req.Status = 0
		req.Level = 0
		res.Data, res.Totle, err = service.Ticket().GetTicketList(&req.V2Ticket, "", "id", "desc", req.Offset, req.Limit)

		return
	}

	return
}

func (c *ControllerV1) TicketClose(ctx context.Context, req *v1.TicketCloseReq) (res *v1.TicketCloseRes, err error) {
	err = service.Ticket().CloseTicketByUserIdAndId(req.Ids, req.TUserID)
	if err != nil {
		return res, err
	}
	return res, err
}

func (c *ControllerV1) TicketCreate(ctx context.Context, req *v1.TicketCreateReq) (res *v1.TicketCreateRes, err error) {
	req.V2Ticket.UserId = req.TUserID
	req.V2Ticket.Status = -1
	req.V2Ticket.ReplyStatus = -1
	if strings.TrimSpace(req.V2Ticket.Subject) == "" {
		err = errors.New("主题不能为空！")
		return
	}
	req.V2Ticket.Subject = ghtml.Entities(req.V2Ticket.Subject)
	err = service.Ticket().AETicket(&req.V2Ticket)
	return
}

func (c *ControllerV1) TicketMessage(ctx context.Context, req *v1.TicketMessageReq) (res *v1.TicketMessageRes, err error) {
	res = &v1.TicketMessageRes{}
	res.Data, err = service.TicketMessage().GetTicketMessageArrByTicketIdAndUserId(req.TicketId, req.TUserID)
	return
}

func (c *ControllerV1) TicketMessageAdd(ctx context.Context, req *v1.TicketMessageAddReq) (res *v1.TicketMessageAddRes, err error) {
	res = &v1.TicketMessageAddRes{}
	req.V2TicketMessage.UserId = g.RequestFromCtx(ctx).Get("TUserID").Int()
	req.V2TicketMessage.Message = ghtml.Entities(req.V2TicketMessage.Message)
	err = service.TicketMessage().SaveTicketMessageUser(&req.V2TicketMessage)
	return
}
