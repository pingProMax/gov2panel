// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// V2RechargeRecords is the golang structure of table v2_recharge_records for DAO operations like Where/Data.
type V2RechargeRecords struct {
	g.Meta          `orm:"table:v2_recharge_records, do:true"`
	Id              interface{} //
	Amount          interface{} // 金额
	UserId          interface{} // 用户id
	OperateType     interface{} // 1充值 2消费
	RechargeName    interface{} // 充值类型 operate_type=1才有
	ConsumptionName interface{} // 消费类型 operate_type=2才有
	Remarks         interface{} // 备注
	TransactionId   interface{} // 订单号 规则看程序注释
	CreatedAt       *gtime.Time //
	UpdatedAt       *gtime.Time //
}
