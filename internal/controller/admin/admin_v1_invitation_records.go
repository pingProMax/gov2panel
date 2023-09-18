package admin

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gov2panel/api/admin/v1"
	"gov2panel/internal/service"
)

func (c *ControllerV1) InvitationRecords(ctx context.Context, req *v1.InvitationRecordsReq) (res *v1.InvitationRecordsRes, err error) {
	switch g.RequestFromCtx(ctx).Method {
	case "GET":
		setTplAdmin(ctx, "invitation_records", nil)
	case "POST":
		res = &v1.InvitationRecordsRes{}
		res.Data, res.Totle, err = service.InvitationRecords().GetInvitationRecordsList(req, req.Sort, req.Order, req.Offset, req.Limit)
		return

	default:
		return
	}
	return
}

func (c *ControllerV1) InvitationRecordsUpState(ctx context.Context, req *v1.InvitationRecordsUpStateReq) (res *v1.InvitationRecordsUpStateRes, err error) {
	err = service.InvitationRecords().AdminiUpStateById(req.Id, req.State)
	return
}
