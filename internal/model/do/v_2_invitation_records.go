// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// V2InvitationRecords is the golang structure of table v2_invitation_records for DAO operations like Where/Data.
type V2InvitationRecords struct {
	g.Meta            `orm:"table:v2_invitation_records, do:true"`
	Id                interface{} //
	Amount            interface{} // 金额
	UserId            interface{} // 邀请者
	FromUserId        interface{} // 被邀请者
	CommissionRate    interface{} // 佣金比例
	RechargeRecordsId interface{} // 订单id
	CreatedAt         *gtime.Time // 创建时间
	UpdatedAt         *gtime.Time // 更新时间
	OperateType       interface{} // 1邀请 2提现
	State             interface{} // 状态 0未审核 1审核 2拒绝
}
