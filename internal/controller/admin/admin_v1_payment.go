package admin

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gov2panel/api/admin/v1"
	"gov2panel/internal/service"
)

func (c *ControllerV1) Payment(ctx context.Context, req *v1.PaymentReq) (res *v1.PaymentRes, err error) {
	switch g.RequestFromCtx(ctx).Method {
	case "GET":
		setTplAdmin(ctx, "payment", nil)
	case "POST":
		res = &v1.PaymentRes{}
		res.Data, err = service.Payment().AdminGetPaymentAllList(req.V2Payment)
		return

	default:
		return
	}
	return
}

func (c *ControllerV1) PaymentAE(ctx context.Context, req *v1.PaymentAEReq) (res *v1.PaymentAERes, err error) {
	err = service.Payment().AEPayment(&req.V2Payment)
	if err != nil {
		return res, err
	}
	return res, err
}

func (c *ControllerV1) PaymentDel(ctx context.Context, req *v1.PaymentDelReq) (res *v1.PaymentDelRes, err error) {
	err = service.Payment().DelPayment(req.Ids)
	if err != nil {
		return res, err
	}
	return res, err
}

func (c *ControllerV1) PaymentGetShow(ctx context.Context, req *v1.PaymentGetShowReq) (res *v1.PaymentGetShowRes, err error) {
	res = &v1.PaymentGetShowRes{}
	res.Data, err = service.Payment().GetPaymentShowList()

	return
}
