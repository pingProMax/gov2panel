package admin

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"gov2panel/internal/model/model"
	"gov2panel/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"

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

// 服务器流量
func (c *ControllerV1) ProxyServiceFlow(ctx context.Context, req *v1.ProxyServiceFlowReq) (res *v1.ProxyServiceFlowRes, err error) {
	res = &v1.ProxyServiceFlowRes{}
	res.ServiceFlowTop = make([]*model.ProxyServiceFlow, 0)
	// res.ServiceFlowTop = append(res.ServiceFlowTop, &model.ProxyServiceFlow{
	// 	Id:   999,
	// 	Name: "test1",
	// 	Flow: 3564444444000,
	// })
	// res.ServiceFlowTop = append(res.ServiceFlowTop, &model.ProxyServiceFlow{
	// 	Id:   998,
	// 	Name: "test2",
	// 	Flow: 6564444444000,
	// })
	// res.ServiceFlowTop = append(res.ServiceFlowTop, &model.ProxyServiceFlow{
	// 	Id:   998,
	// 	Name: "test3",
	// 	Flow: 5564444444000,
	// })

	cacheKeyS, err := gcache.KeyStrings(ctx)
	if err != nil {
		return res, err
	}

	serviceList, err := service.ProxyService().GetProxyServiceAllList()
	if err != nil {
		return res, err
	}

	for _, v := range cacheKeyS {
		if strings.HasPrefix(v, "SERVER_") && strings.HasSuffix(v, fmt.Sprintf("_%s_FLOW", req.Date)) {
			idStr := strings.ReplaceAll(v, "SERVER_", "")
			idStr = strings.ReplaceAll(idStr, fmt.Sprintf("_%s_FLOW", req.Date), "")
			id := gconv.Int(idStr)

			flow, err := gcache.Get(ctx, v)
			if err != nil {
				return res, err
			}

			for _, service := range serviceList {
				if service.Id == id {
					f := &model.ProxyServiceFlow{
						Id:   id,
						Name: service.Name,
						Flow: flow.Int64(),
					}
					res.ServiceFlowTop = append(res.ServiceFlowTop, f)
					break
				}

			}
		}
	}

	// 使用 sort.Slice 函数进行排序
	sort.Slice(res.ServiceFlowTop, func(i, j int) bool {
		return res.ServiceFlowTop[i].Flow > res.ServiceFlowTop[j].Flow
	})

	return res, err
}

// 获取所有服务器当前在线用户数量和服务器最后提交时间
// map[服务器id][type 1在线数量、2服务器最后提交时间]int
func (c *ControllerV1) OnlineUserCountAndLastPushAt(ctx context.Context, req *v1.OnlineUserCountAndLastPushAtReq) (res *v1.OnlineUserCountAndLastPushAtRes, err error) {
	res = &v1.OnlineUserCountAndLastPushAtRes{}
	res.Data, err = service.ProxyService().GetOnlineUserCountAndLastPushAt()

	return res, err
}
