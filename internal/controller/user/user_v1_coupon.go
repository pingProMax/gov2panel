package user

import (
	"context"

	v1 "gov2panel/api/user/v1"
	"gov2panel/internal/service"
)

func (c *ControllerV1) Coupon(ctx context.Context, req *v1.CouponReq) (res *v1.CouponRes, err error) {
	return service.Coupon().CheckCouponCanUseByCode(ctx, req)
}
