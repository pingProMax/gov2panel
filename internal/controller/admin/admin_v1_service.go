package admin

import (
	"context"

	"gov2panel/internal/service"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gov2panel/api/admin/v1"
)

func (c *ControllerV1) ProxyService(ctx context.Context, req *v1.ProxyServiceReq) (res *v1.ProxyServiceRes, err error) {
	switch g.RequestFromCtx(ctx).Method {
	case "GET":
		setTplAdmin(ctx, "service", nil)
	case "POST":
		res = &v1.ProxyServiceRes{}
		res.Data, res.Totle, err = service.ProxyService().GetProxyServiceList(req, req.Sort, req.Order, req.Offset, req.Limit)
		return

	default:
		return
	}
	return
}

func (c *ControllerV1) ProxyServiceAE(ctx context.Context, req *v1.ProxyServiceAEReq) (res *v1.ProxyServiceAERes, err error) {
	err = service.ProxyService().AEProxyService(&req.V2ProxyService)
	if err != nil {
		return res, err
	}
	return res, err
}

func (c *ControllerV1) ProxyServiceDel(ctx context.Context, req *v1.ProxyServiceDelReq) (res *v1.ProxyServiceDelRes, err error) {
	err = service.ProxyService().DelProxyService(req.Ids)
	if err != nil {
		return res, err
	}
	return res, err
}
