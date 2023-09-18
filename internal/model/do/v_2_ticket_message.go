// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// V2TicketMessage is the golang structure of table v2_ticket_message for DAO operations like Where/Data.
type V2TicketMessage struct {
	g.Meta    `orm:"table:v2_ticket_message, do:true"`
	Id        interface{} //
	UserId    interface{} //
	TicketId  interface{} //
	Message   interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
