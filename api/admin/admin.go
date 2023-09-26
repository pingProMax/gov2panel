// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package admin

import (
	"context"
	
	"gov2panel/api/admin/v1"
)

type IAdminV1 interface {
	Coupon(ctx context.Context, req *v1.CouponReq) (res *v1.CouponRes, err error)
	CouponAE(ctx context.Context, req *v1.CouponAEReq) (res *v1.CouponAERes, err error)
	CouponDel(ctx context.Context, req *v1.CouponDelReq) (res *v1.CouponDelRes, err error)
	Index(ctx context.Context, req *v1.IndexReq) (res *v1.IndexRes, err error)
	InvitationRecords(ctx context.Context, req *v1.InvitationRecordsReq) (res *v1.InvitationRecordsRes, err error)
	InvitationRecordsUpState(ctx context.Context, req *v1.InvitationRecordsUpStateReq) (res *v1.InvitationRecordsUpStateRes, err error)
	Knowledge(ctx context.Context, req *v1.KnowledgeReq) (res *v1.KnowledgeRes, err error)
	KnowledgeAE(ctx context.Context, req *v1.KnowledgeAEReq) (res *v1.KnowledgeAERes, err error)
	KnowledgeDel(ctx context.Context, req *v1.KnowledgeDelReq) (res *v1.KnowledgeDelRes, err error)
	Payment(ctx context.Context, req *v1.PaymentReq) (res *v1.PaymentRes, err error)
	PaymentAE(ctx context.Context, req *v1.PaymentAEReq) (res *v1.PaymentAERes, err error)
	PaymentDel(ctx context.Context, req *v1.PaymentDelReq) (res *v1.PaymentDelRes, err error)
	PaymentGetShow(ctx context.Context, req *v1.PaymentGetShowReq) (res *v1.PaymentGetShowRes, err error)
	Plan(ctx context.Context, req *v1.PlanReq) (res *v1.PlanRes, err error)
	PlanAE(ctx context.Context, req *v1.PlanAEReq) (res *v1.PlanAERes, err error)
	PlanDel(ctx context.Context, req *v1.PlanDelReq) (res *v1.PlanDelRes, err error)
	PlanGetShow(ctx context.Context, req *v1.PlanGetShowReq) (res *v1.PlanGetShowRes, err error)
	PlanGetShowAndResetTrafficMethod1(ctx context.Context, req *v1.PlanGetShowAndResetTrafficMethod1Req) (res *v1.PlanGetShowAndResetTrafficMethod1Res, err error)
	ProxyService(ctx context.Context, req *v1.ProxyServiceReq) (res *v1.ProxyServiceRes, err error)
	ProxyServiceAE(ctx context.Context, req *v1.ProxyServiceAEReq) (res *v1.ProxyServiceAERes, err error)
	ProxyServiceDel(ctx context.Context, req *v1.ProxyServiceDelReq) (res *v1.ProxyServiceDelRes, err error)
	ProxyServiceFlow(ctx context.Context, req *v1.ProxyServiceFlowReq) (res *v1.ProxyServiceFlowRes, err error)
	OnlineUserCountAndLastPushAt(ctx context.Context, req *v1.OnlineUserCountAndLastPushAtReq) (res *v1.OnlineUserCountAndLastPushAtRes, err error)
	RechargeRecords(ctx context.Context, req *v1.RechargeRecordsReq) (res *v1.RechargeRecordsRes, err error)
	RechargeRecordsAdd(ctx context.Context, req *v1.RechargeRecordsAddReq) (res *v1.RechargeRecordsAddRes, err error)
	RechargeRecordsUpRemarks(ctx context.Context, req *v1.RechargeRecordsUpRemarksReq) (res *v1.RechargeRecordsUpRemarksRes, err error)
	ServerRoute(ctx context.Context, req *v1.ServerRouteReq) (res *v1.ServerRouteRes, err error)
	ServerRouteAE(ctx context.Context, req *v1.ServerRouteAEReq) (res *v1.ServerRouteAERes, err error)
	ServerRouteDel(ctx context.Context, req *v1.ServerRouteDelReq) (res *v1.ServerRouteDelRes, err error)
	ServerRouteAll(ctx context.Context, req *v1.ServerRouteAllReq) (res *v1.ServerRouteAllRes, err error)
	Setting(ctx context.Context, req *v1.SettingReq) (res *v1.SettingRes, err error)
	SettingAE(ctx context.Context, req *v1.SettingAEReq) (res *v1.SettingAERes, err error)
	SettingDel(ctx context.Context, req *v1.SettingDelReq) (res *v1.SettingDelRes, err error)
	Ticket(ctx context.Context, req *v1.TicketReq) (res *v1.TicketRes, err error)
	TicketAE(ctx context.Context, req *v1.TicketAEReq) (res *v1.TicketAERes, err error)
	TicketDel(ctx context.Context, req *v1.TicketDelReq) (res *v1.TicketDelRes, err error)
	TicketClose(ctx context.Context, req *v1.TicketCloseReq) (res *v1.TicketCloseRes, err error)
	TicketMessage(ctx context.Context, req *v1.TicketMessageReq) (res *v1.TicketMessageRes, err error)
	TicketMessageAdd(ctx context.Context, req *v1.TicketMessageAddReq) (res *v1.TicketMessageAddRes, err error)
	User(ctx context.Context, req *v1.UserReq) (res *v1.UserRes, err error)
	UserAE(ctx context.Context, req *v1.UserAEReq) (res *v1.UserAERes, err error)
	UserDel(ctx context.Context, req *v1.UserDelReq) (res *v1.UserDelRes, err error)
	UserUpBanned1(ctx context.Context, req *v1.UserUpBanned1Req) (res *v1.UserUpBanned1Res, err error)
}


