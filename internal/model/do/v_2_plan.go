// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// V2Plan is the golang structure of table v2_plan for DAO operations like Where/Data.
type V2Plan struct {
	g.Meta             `orm:"table:v2_plan, do:true"`
	Id                 interface{} //
	TransferEnable     interface{} // 流量(GB)
	SpeedLimit         interface{} // 速度限制
	Name               interface{} // 名称
	Show               interface{} // 是否显示
	OrderId            interface{} // 顺序
	Renew              interface{} // 是否允许续购
	Content            interface{} // 描述
	Expired            interface{} // 有效期 day
	Price              interface{} // 价格
	ResetTrafficMethod interface{} // 套餐类型，1 覆盖、2 叠加
	CapacityLimit      interface{} // 最大用户
	CreatedAt          *gtime.Time //
	UpdatedAt          *gtime.Time //
	Remarks            interface{} // 备注
}
