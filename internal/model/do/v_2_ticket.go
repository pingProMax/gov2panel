// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// V2Ticket is the golang structure of table v2_ticket for DAO operations like Where/Data.
type V2Ticket struct {
	g.Meta      `orm:"table:v2_ticket, do:true"`
	Id          interface{} //
	UserId      interface{} //
	Subject     interface{} //
	Level       interface{} // 1低 2中 3高
	Status      interface{} // -1:已开启 1:已关闭
	ReplyStatus interface{} // -1:待回复 1:已回复
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
}
