// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// V2ServerRoute is the golang structure of table v2_server_route for DAO operations like Where/Data.
type V2ServerRoute struct {
	g.Meta      `orm:"table:v2_server_route, do:true"`
	Id          interface{} //
	Remarks     interface{} // 备注
	Match       interface{} // 规则
	Action      interface{} // block|dns
	ActionValue interface{} //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	Enable      interface{} // 是否启用
}
