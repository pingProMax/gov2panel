package admin

import (
	"context"

	v1 "gov2panel/api/admin/v1"
	"gov2panel/internal/service"
	"gov2panel/internal/utils"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) RechargeRecords(ctx context.Context, req *v1.RechargeRecordsReq) (res *v1.RechargeRecordsRes, err error) {
	switch g.RequestFromCtx(ctx).Method {
	case "GET":
		setTplAdmin(ctx, "recharge_records", nil)
	case "POST":
		res = &v1.RechargeRecordsRes{}
		res.Data, res.Totle, err = service.RechargeRecords().GetRechargeRecordsList(req, req.Sort, req.Order, req.Offset, req.Limit)
		return

	default:
		return
	}
	return
}

func (c *ControllerV1) RechargeRecordsAdd(ctx context.Context, req *v1.RechargeRecordsAddReq) (res *v1.RechargeRecordsAddRes, err error) {
	req.V2RechargeRecords.TransactionId = utils.RechargeOrderNo(req.V2RechargeRecords.Amount, 0)
	err = service.RechargeRecords().SaveRechargeRecords(&req.V2RechargeRecords, "admin", req.Amount, 0, "")
	return res, err
}

func (c *ControllerV1) RechargeRecordsUpRemarks(ctx context.Context, req *v1.RechargeRecordsUpRemarksReq) (res *v1.RechargeRecordsUpRemarksRes, err error) {
	err = service.RechargeRecords().UpRechargeRecordsRemarksById(req.Id, req.Remarks)
	res = &v1.RechargeRecordsUpRemarksRes{}
	return res, err
}
