package server_api

import (
	"context"
	"encoding/json"

	v1 "gov2panel/api/server_api/v1"
	"gov2panel/internal/model/model"
	"gov2panel/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) Api(ctx context.Context, req *v1.ApiReq) (res *v1.ApiRes, err error) {
	res = &v1.ApiRes{}
	server, _, routeList, err := service.ProxyService().GetServiceById(req.NodeId)
	if err != nil {
		return
	}

	routeArr := make([]*model.Route, 0)

	for i := 0; i < len(routeList); i++ {

		var strSlice []string
		err = json.Unmarshal([]byte(routeList[i].Match), &strSlice)

		routeArr = append(routeArr, &model.Route{
			Id:          routeList[i].Id,
			Action:      routeList[i].Action,
			Match:       strSlice,
			ActionValue: routeList[i].ActionValue,
		})
	}

	json.Unmarshal([]byte(server.ServiceJson), &res)
	ress := map[string]interface{}(*res)
	ress["routes"] = routeArr
	// ress["plan"] = planList
	ghttp.RequestFromCtx(ctx).Response.WriteJsonExit(ress)
	return
}
