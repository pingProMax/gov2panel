package user

import (
	"context"

	v1 "gov2panel/api/user/v1"
	"gov2panel/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Node(ctx context.Context, req *v1.NodeReq) (res *v1.NodeRes, err error) {

	switch g.RequestFromCtx(ctx).Method {
	case "GET":
		setTplUser(ctx, "node", nil)
	case "POST":
		res = &v1.NodeRes{}

		user := c.getUser(ctx)
		res.Data, err = service.ProxyService().GetServiceListByPlanIdAndShow1(user.GroupId)

		return
	}

	return
}

// 获取所有服务器当前在线用户数量和服务器最后提交时间
// map[服务器id][type 1在线数量、2服务器最后提交时间]int
func (c *ControllerV1) OnlineUserCountAndLastPushAt(ctx context.Context, req *v1.OnlineUserCountAndLastPushAtReq) (res *v1.OnlineUserCountAndLastPushAtRes, err error) {
	res = &v1.OnlineUserCountAndLastPushAtRes{}
	res.Data, err = service.ProxyService().GetOnlineUserCountAndLastPushAt()

	return res, err
}
