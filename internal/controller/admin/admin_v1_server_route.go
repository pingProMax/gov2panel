package admin

import (
	"context"
	"encoding/json"
	"strings"

	v1 "gov2panel/api/admin/v1"
	"gov2panel/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) ServerRoute(ctx context.Context, req *v1.ServerRouteReq) (res *v1.ServerRouteRes, err error) {
	switch g.RequestFromCtx(ctx).Method {
	case "GET":
		setTplAdmin(ctx, "server_route", nil)
	case "POST":
		res = &v1.ServerRouteRes{}
		res.Data, res.Totle, err = service.ServerRoute().GetServerRouteList(req, req.Sort, req.Order, req.Offset, req.Limit)
		return

	default:
		return
	}
	return
}

// 待测试
func (c *ControllerV1) ServerRouteAE(ctx context.Context, req *v1.ServerRouteAEReq) (res *v1.ServerRouteAERes, err error) {
	match, _ := json.Marshal(strings.Split(req.Match, "\r\n"))
	req.Match = string(match)
	err = service.ServerRoute().AEServerRoute(&req.V2ServerRoute)
	if err != nil {
		return res, err
	}
	return res, err
}

func (c *ControllerV1) ServerRouteDel(ctx context.Context, req *v1.ServerRouteDelReq) (res *v1.ServerRouteDelRes, err error) {
	err = service.ServerRoute().DelServerRoute(req.Ids)
	if err != nil {
		return res, err
	}
	return res, err
}

func (c *ControllerV1) ServerRouteAll(ctx context.Context, req *v1.ServerRouteAllReq) (res *v1.ServerRouteAllRes, err error) {
	res = &v1.ServerRouteAllRes{}
	res.Data, err = service.ServerRoute().ServerRouteAll()
	if err != nil {
		return res, err
	}
	return res, err
}
