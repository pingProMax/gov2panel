package user

import (
	"context"
	"errors"
	"time"

	v1 "gov2panel/api/user/v1"
	"gov2panel/internal/service"

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
	return service.Plan().UserBuy(req)
}
