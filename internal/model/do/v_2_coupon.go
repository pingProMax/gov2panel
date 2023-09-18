// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// V2Coupon is the golang structure of table v2_coupon for DAO operations like Where/Data.
type V2Coupon struct {
	g.Meta           `orm:"table:v2_coupon, do:true"`
	Id               interface{} //
	Code             interface{} // 优惠码
	Name             interface{} // 名称
	Type             interface{} // 类型 1金额优惠 2百分比优惠
	Value            interface{} // 优惠多少
	Enable           interface{} // 是否启用
	LimitUse         interface{} // 每个用户可使用次数
	LimitUseWithUser interface{} // 最大使用次数
	LimitPlanId      interface{} // 指定订阅
	StartedAt        *gtime.Time // 有效期开始时间
	EndedAt          *gtime.Time // 有效期结束时间
	CreatedAt        *gtime.Time //
	UpdatedAt        *gtime.Time //
	Remarks          interface{} // 备注
}
