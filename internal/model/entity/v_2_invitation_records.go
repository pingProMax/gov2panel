// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// V2InvitationRecords is the golang structure for table v2_invitation_records.
type V2InvitationRecords struct {
	Id                int         `json:"id"                  ` //
	Amount            float64     `json:"amount"              ` // 金额
	UserId            int         `json:"user_id"             ` // 邀请者
	FromUserId        int         `json:"from_user_id"        ` // 被邀请者
	CommissionRate    int         `json:"commission_rate"     ` // 佣金比例
	RechargeRecordsId int         `json:"recharge_records_id" ` // 订单id
	CreatedAt         *gtime.Time `json:"created_at"          ` //
	UpdatedAt         *gtime.Time `json:"updated_at"          ` //
	OperateType       int         `json:"operate_type"        ` // 1邀请 2提现
	State             int         `json:"state"               ` // 状态 -1未审核 1审核 2拒绝
}
