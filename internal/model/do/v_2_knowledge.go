// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// V2Knowledge is the golang structure of table v2_knowledge for DAO operations like Where/Data.
type V2Knowledge struct {
	g.Meta    `orm:"table:v2_knowledge, do:true"`
	Id        interface{} //
	Category  interface{} // 分類名
	Title     interface{} // 標題
	Body      interface{} // 內容
	OrderId   interface{} // 排序
	Show      interface{} // 顯示
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
