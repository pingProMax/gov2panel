package admin

import (
	"context"

	v1 "gov2panel/api/admin/v1"
	"gov2panel/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Plan(ctx context.Context, req *v1.PlanReq) (res *v1.PlanRes, err error) {
	switch g.RequestFromCtx(ctx).Method {
	case "GET":
		setTplAdmin(ctx, "plan", nil)
	case "POST":
		res = &v1.PlanRes{}
		res.Data, err = service.Plan().GetPlanAllList(req.V2Plan)
		return

	default:
		return
	}
	return
}

func (c *ControllerV1) PlanAE(ctx context.Context, req *v1.PlanAEReq) (res *v1.PlanAERes, err error) {
	err = service.Plan().AEPlan(&req.V2Plan)
	if err != nil {
		return res, err
	}
	return res, err
}

func (c *ControllerV1) PlanDel(ctx context.Context, req *v1.PlanDelReq) (res *v1.PlanDelRes, err error) {
	err = service.Plan().DelPlan(req.Ids)
	if err != nil {
		return res, err
	}
	return res, err
}

func (c *ControllerV1) PlanGetShow(ctx context.Context, req *v1.PlanGetShowReq) (res *v1.PlanGetShowRes, err error) {
	res = &v1.PlanGetShowRes{}
	res.Data, err = service.Plan().GetPlanShowList()

	return
}

func (c *ControllerV1) PlanGetResetTrafficMethod1(ctx context.Context, req *v1.PlanGetResetTrafficMethod1Req) (res *v1.PlanGetResetTrafficMethod1Res, err error) {
	res = &v1.PlanGetResetTrafficMethod1Res{}
	res.Data, err = service.Plan().GetPlanResetTrafficMethod1List()

	return
}
