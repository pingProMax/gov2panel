package user

import (
	"context"

	v1 "gov2panel/api/user/v1"
	"gov2panel/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Plan(ctx context.Context, req *v1.PlanReq) (res *v1.PlanRes, err error) {
	res = &v1.PlanRes{}
	res.Data, err = service.Plan().GetPlanShowList()
	setTplUser(ctx, "plan", g.Map{"data": res.Data})
	return
}

func (c *ControllerV1) Plan2(ctx context.Context, req *v1.Plan2Req) (res *v1.Plan2Res, err error) {
	res = &v1.Plan2Res{}
	res.Data, err = service.Plan().GetPlanById(req.Id)
	if res.Data == nil {
		return res, gerror.NewCode(gcode.CodeNotFound)
	}
	setTplUser(ctx, "plan2", g.Map{"data": res.Data})
	return
}

// 购买api
func (c *ControllerV1) Buy(ctx context.Context, req *v1.BuyReq) (res *v1.BuyRes, err error) {
	res = &v1.BuyRes{}
	return service.Plan().UserBuy(req)
}
