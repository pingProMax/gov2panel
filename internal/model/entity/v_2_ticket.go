// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// V2Ticket is the golang structure for table v2_ticket.
type V2Ticket struct {
	Id          int         `json:"id"           orm:"id"           ` //
	UserId      int         `json:"user_id"      orm:"user_id"      ` //
	Subject     string      `json:"subject"      orm:"subject"      ` //
	Level       int         `json:"level"        orm:"level"        ` // 1低 2中 3高
	Status      int         `json:"status"       orm:"status"       ` // -1:已开启 1:已关闭
	ReplyStatus int         `json:"reply_status" orm:"reply_status" ` // -1:待回复 1:已回复
	CreatedAt   *gtime.Time `json:"created_at"   orm:"created_at"   ` //
	UpdatedAt   *gtime.Time `json:"updated_at"   orm:"updated_at"   ` //
}
