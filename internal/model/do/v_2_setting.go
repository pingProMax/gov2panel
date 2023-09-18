// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// V2Setting is the golang structure of table v2_setting for DAO operations like Where/Data.
type V2Setting struct {
	g.Meta    `orm:"table:v2_setting, do:true"`
	Code      interface{} //
	Value     interface{} //
	OrderId   interface{} // 顺序
	Remarks   interface{} // 备注
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
