// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	v1 "gov2panel/api/user/v1"
)

type IUserV1 interface {
	Coupon(ctx context.Context, req *v1.CouponReq) (res *v1.CouponRes, err error)
	Flow(ctx context.Context, req *v1.FlowReq) (res *v1.FlowRes, err error)
	Index(ctx context.Context, req *v1.IndexReq) (res *v1.IndexRes, err error)
	AppBulletin(ctx context.Context, req *v1.AppBulletinReq) (res *v1.AppBulletinRes, err error)
	Knowledge(ctx context.Context, req *v1.KnowledgeReq) (res *v1.KnowledgeRes, err error)
	Node(ctx context.Context, req *v1.NodeReq) (res *v1.NodeRes, err error)
	OnlineUserCountAndLastPushAt(ctx context.Context, req *v1.OnlineUserCountAndLastPushAtReq) (res *v1.OnlineUserCountAndLastPushAtRes, err error)
	Plan(ctx context.Context, req *v1.PlanReq) (res *v1.PlanRes, err error)
	Plan2(ctx context.Context, req *v1.Plan2Req) (res *v1.Plan2Res, err error)
	PlanRenew(ctx context.Context, req *v1.PlanRenewReq) (res *v1.PlanRenewRes, err error)
	Buy(ctx context.Context, req *v1.BuyReq) (res *v1.BuyRes, err error)
	Renew(ctx context.Context, req *v1.RenewReq) (res *v1.RenewRes, err error)
	Ticket(ctx context.Context, req *v1.TicketReq) (res *v1.TicketRes, err error)
	TicketClose(ctx context.Context, req *v1.TicketCloseReq) (res *v1.TicketCloseRes, err error)
	TicketCreate(ctx context.Context, req *v1.TicketCreateReq) (res *v1.TicketCreateRes, err error)
	TicketMessage(ctx context.Context, req *v1.TicketMessageReq) (res *v1.TicketMessageRes, err error)
	TicketMessageAdd(ctx context.Context, req *v1.TicketMessageAddReq) (res *v1.TicketMessageAddRes, err error)
	User(ctx context.Context, req *v1.UserReq) (res *v1.UserRes, err error)
	UserUpPasswd(ctx context.Context, req *v1.UserUpPasswdReq) (res *v1.UserUpPasswdRes, err error)
	ResetTokenAndUuid(ctx context.Context, req *v1.ResetTokenAndUuidReq) (res *v1.ResetTokenAndUuidRes, err error)
	Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error)
	Wallet(ctx context.Context, req *v1.WalletReq) (res *v1.WalletRes, err error)
	V2RechargeRecords(ctx context.Context, req *v1.V2RechargeRecordsReq) (res *v1.V2RechargeRecordsRes, err error)
	InvitationRecords(ctx context.Context, req *v1.InvitationRecordsReq) (res *v1.InvitationRecordsRes, err error)
	Recharge(ctx context.Context, req *v1.RechargeReq) (res *v1.RechargeRes, err error)
	GetPayList(ctx context.Context, req *v1.GetPayListReq) (res *v1.GetPayListRes, err error)
	PayRedirection(ctx context.Context, req *v1.PayRedirectionReq) (res *v1.PayRedirectionRes, err error)
	CommissionTransferBalance(ctx context.Context, req *v1.CommissionTransferBalanceReq) (res *v1.CommissionTransferBalanceRes, err error)
	CWithdrawalBalance(ctx context.Context, req *v1.CWithdrawalBalanceReq) (res *v1.CWithdrawalBalanceRes, err error)
}
