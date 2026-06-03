// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// V2ServiceRelay is the golang structure of table v2_service_relay for DAO operations like Where/Data.
type V2ServiceRelay struct {
	g.Meta    `orm:"table:v2_service_relay, do:true"`
	Id        interface{} //
	Ip        interface{} // ip
	NameGroup interface{} // 组名字
	Asn       interface{} // asn分组 AS9808|AS4134|AS4837
	Show      interface{} // 是否启用
	OrderId   interface{} // 顺序
	Remarks   interface{} // 备注
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
