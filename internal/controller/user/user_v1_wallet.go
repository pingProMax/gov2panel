package user

import (
	"context"

	v1 "gov2panel/api/user/v1"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Wallet(ctx context.Context, req *v1.WalletReq) (res *v1.WalletRes, err error) {
	res = &v1.WalletRes{}
	inviteCount, err := service.User().GetInviteCountByUserId(req.TUserID)
	if err != nil {
		return res, err
	}
	var user entity.V2User

	g.RequestFromCtx(ctx).GetCtxVar("database_user").Struct(&user)

	cType, cRate := service.User().GetUserCTypeAndCRate(&user)
	setTplUser(ctx, "wallet", g.Map{
		"inviteCount": inviteCount,
		"cType":       cType,
		"cRate":       cRate,
	})

	return
}

func (c *ControllerV1) V2RechargeRecords(ctx context.Context, req *v1.V2RechargeRecordsReq) (res *v1.V2RechargeRecordsRes, err error) {

	res = &v1.V2RechargeRecordsRes{}
	//获取用户充值消费记录
	res.V2RechargeRecordsData, res.V2RechargeRecordsTotle, err = service.RechargeRecords().GetRechargeRecordsListByUserId(req.TUserID, "id", "desc", req.Offset, req.Limit)

	return
}

func (c *ControllerV1) InvitationRecords(ctx context.Context, req *v1.InvitationRecordsReq) (res *v1.InvitationRecordsRes, err error) {
	res = &v1.InvitationRecordsRes{}
	res.Data, res.Totle, err = service.InvitationRecords().GetInvitationRecordsListByUserId(req.TUserID, "id", "desc", req.Offset, req.Limit)
	return
}

// 充值 Recharge 页面
func (c *ControllerV1) Recharge(ctx context.Context, req *v1.RechargeReq) (res *v1.RechargeRes, err error) {
	setTplUser(ctx, "recharge", nil)
	return
}

// 获取支付列表 GetPayList
func (c *ControllerV1) GetPayList(ctx context.Context, req *v1.GetPayListReq) (res *v1.GetPayListRes, err error) {
	res = &v1.GetPayListRes{}
	res.Data, err = service.Payment().GetPaymentShowList()
	return
}

// 支付重定向
func (c *ControllerV1) PayRedirection(ctx context.Context, req *v1.PayRedirectionReq) (res *v1.PayRedirectionRes, err error) {
	res = &v1.PayRedirectionRes{}
	res.Url, err = service.Payment().GetPayUrl(req)
	return
}
