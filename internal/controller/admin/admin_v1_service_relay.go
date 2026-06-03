package admin

import (
	"context"
	v1 "gov2panel/api/admin/v1"
	"gov2panel/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) ServiceRelay(ctx context.Context, req *v1.ServiceRelayReq) (res *v1.ServiceRelayRes, err error) {
	switch g.RequestFromCtx(ctx).Method {
	case "GET":
		setTplAdmin(ctx, "service_relay", nil)
	case "POST":
		res = &v1.ServiceRelayRes{}
		res.Data, res.Totle, err = service.ServerRelay().GetServiceRelayList(req, req.Sort, req.Order, req.Offset, req.Limit)
		return

	default:
		return
	}
	return
}

func (c *ControllerV1) ServiceRelayAE(ctx context.Context, req *v1.ServiceRelayAEReq) (res *v1.ServiceRelayAERes, err error) {
	err = service.ServerRelay().AEServiceRelay(&req.V2ServiceRelay)
	if err != nil {
		return res, err
	}
	return res, err
}

func (c *ControllerV1) ServiceRelayDel(ctx context.Context, req *v1.ServiceRelayDelReq) (res *v1.ServiceRelayDelRes, err error) {
	err = service.ServerRelay().DelServiceRelay(req.Ids)
	if err != nil {
		return res, err
	}
	return res, err
}

func (c *ControllerV1) ServiceRelayShow(ctx context.Context, req *v1.ServiceRelayShowReq) (res *v1.ServiceRelayShowRes, err error) {
	res = &v1.ServiceRelayShowRes{}
	err = service.ServerRelay().UpServiceShow(req.Ids, req.Show)
	return res, err
}
