// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// V2Ticket is the golang structure for table v2_ticket.
type V2Ticket struct {
	Id          int         `json:"id"           ` //
	UserId      int         `json:"user_id"      ` //
	Subject     string      `json:"subject"      ` //
	Level       int         `json:"level"        ` //
	Status      int         `json:"status"       ` // 0:已开启 1:已关闭
	ReplyStatus int         `json:"reply_status" ` // 0:待回复 1:已回复
	CreatedAt   *gtime.Time `json:"created_at"   ` //
	UpdatedAt   *gtime.Time `json:"updated_at"   ` //
}
