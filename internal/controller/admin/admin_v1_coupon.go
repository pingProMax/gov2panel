package admin

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gov2panel/api/admin/v1"
	"gov2panel/internal/service"
)

func (c *ControllerV1) Coupon(ctx context.Context, req *v1.CouponReq) (res *v1.CouponRes, err error) {
	switch g.RequestFromCtx(ctx).Method {
	case "GET":
		setTplAdmin(ctx, "coupon", nil)
	case "POST":
		res = &v1.CouponRes{}
		res.Data, err = service.Coupon().GetCouponAllList(req.V2Coupon)
		return

	default:
		return
	}
	return
}

func (c *ControllerV1) CouponAE(ctx context.Context, req *v1.CouponAEReq) (res *v1.CouponAERes, err error) {
	err = service.Coupon().AECoupon(&req.V2Coupon)
	if err != nil {
		return res, err
	}
	return res, err
}

func (c *ControllerV1) CouponDel(ctx context.Context, req *v1.CouponDelReq) (res *v1.CouponDelRes, err error) {
	err = service.Coupon().DelCoupon(req.Ids)
	if err != nil {
		return res, err
	}
	return res, err
}
