package user

import (
	"context"
	"errors"
	"net/http"
	"time"

	v1 "gov2panel/api/user/v1"
	"gov2panel/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/google/uuid"
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
	uuid := uuid.New().String()
	gcache.Set(ctx, uuid, 1, time.Hour)
	setTplUser(ctx, "plan2", g.Map{"data": res.Data, "uuid": uuid})
	return
}

func (c *ControllerV1) PlanRenew(ctx context.Context, req *v1.PlanRenewReq) (res *v1.PlanRenewRes, err error) {
	res = &v1.PlanRenewRes{}
	res.Data, err = service.Plan().GetPlanById(req.Id)
	if res.Data == nil {
		return res, gerror.NewCode(gcode.CodeNotFound)
	}

	user := c.getUser(ctx)

	if res.Data.Id != user.GroupId {
		ghttp.RequestFromCtx(ctx).Response.RedirectTo("/user", http.StatusFound)
		ghttp.RequestFromCtx(ctx).ExitAll()
		return
	}
	uuid := uuid.New().String()
	gcache.Set(ctx, uuid, 1, time.Hour)
	setTplUser(ctx, "plan_renew", g.Map{"data": res.Data, "uuid": uuid})
	return
}

// 购买api
func (c *ControllerV1) Buy(ctx context.Context, req *v1.BuyReq) (res *v1.BuyRes, err error) {
	res = &v1.BuyRes{}
	uuid, err := gcache.Get(ctx, req.Uuid)
	if err != nil {
		return
	}
	if uuid.Int() != 1 {
		err = errors.New("数据验证错误，请刷新页面")
		return
	} else {
		gcache.Remove(ctx, req.Uuid)
	}

	//检查要购买的套餐
	plan, err := service.Plan().GetPlanById(req.PlanId)
	if err != nil {
		return
	}
	if plan == nil {
		return res, errors.New("套餐不存在")
	}
	if plan.Show != 1 {
		return res, errors.New("套餐未开启")
	}
	if plan.Price < 0 || plan.Expired < 0 {
		return res, errors.New("套餐设置不对请联系管理员")
	}

	err = service.Plan().UserBuyAndRenew(ctx, req.Code, plan)
	return
}

// 续费api
func (c *ControllerV1) Renew(ctx context.Context, req *v1.RenewReq) (res *v1.RenewRes, err error) {
	res = &v1.RenewRes{}
	uuid, err := gcache.Get(ctx, req.Uuid)
	if err != nil {
		return
	}
	if uuid.Int() != 1 {
		err = errors.New("数据验证错误，请刷新页面")
		return
	} else {
		gcache.Remove(ctx, req.Uuid)
	}

	user := c.getUser(ctx)

	//检查套餐
	plan, err := service.Plan().GetPlanById(user.GroupId)
	if err != nil {
		return
	}
	if plan == nil {
		return res, errors.New("套餐不存在")
	}
	if plan.Price < 0 || plan.Expired < 0 {
		return res, errors.New("套餐设置不对请联系管理员")
	}
	if plan.Renew != 1 {
		return res, errors.New("当前套餐不允许续费！")
	}

	err = service.Plan().UserBuyAndRenew(ctx, req.Code, plan)
	return
}
