package v1

import (
	"gov2panel/internal/model/entity"
	"gov2panel/internal/model/model"

	"github.com/gogf/gf/v2/frame/g"
)

type WalletReq struct {
	g.Meta  `path:"/wallet" tags:"Wallet" method:"get" summary:"钱包页面"`
	TUserID int
}
type WalletRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type V2RechargeRecordsReq struct {
	g.Meta  `path:"/recharge_records" tags:"Wallet" method:"post" summary:"充值消费记录"`
	TUserID int

	OffsetLimit
}
type V2RechargeRecordsRes struct {
	g.Meta                 `mime:"text/html" example:"string"`
	V2RechargeRecordsData  []*entity.V2RechargeRecords `json:"recharge_records"`
	V2RechargeRecordsTotle int                         `json:"recharge_records_totle"`
}

type InvitationRecordsReq struct {
	g.Meta  `path:"/invitation_records" tags:"InvitationRecords" method:"post" summary:"邀请收入记录"`
	TUserID int
	OffsetLimit
}

type InvitationRecordsRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Data   []*model.InvitationRecordsInfo `json:"data"`
	Totle  int                            `json:"totle"`
}

type RechargeReq struct {
	g.Meta  `path:"/recharge" tags:"Wallet" method:"get" summary:"充值页面"`
	TUserID int
}

type RechargeRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type GetPayListReq struct {
	g.Meta  `path:"/pay_list" tags:"Wallet" method:"post" summary:"获取支付列表 api"`
	TUserID int
}

type GetPayListRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Data   []*entity.V2Payment `json:"data"`
}

type PayRedirectionReq struct {
	g.Meta    `path:"/pay_redirection" tags:"Wallet" method:"post" summary:"支付重定向"`
	TUserID   int
	PaymentId int     `json:"payment_id"`
	Amount    float64 `json:"amount"`
}

type PayRedirectionRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Url    string `json:"url"`
}
