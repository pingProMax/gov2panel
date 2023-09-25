package user

import (
	"context"

	v1 "gov2panel/api/user/v1"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Node(ctx context.Context, req *v1.NodeReq) (res *v1.NodeRes, err error) {
	var user entity.V2User
	err = g.RequestFromCtx(ctx).GetCtxVar("database_user").Struct(&user)
	if err != nil {
		g.RequestFromCtx(ctx).Response.Write(err.Error())
		return
	}

	switch g.RequestFromCtx(ctx).Method {
	case "GET":
		setTplUser(ctx, "node", nil)
	case "POST":
		res = &v1.NodeRes{}

		res.Data, err = service.ProxyService().GetServiceListByPlanIdAndShow1(user.GroupId)

		return
	}

	return
}
