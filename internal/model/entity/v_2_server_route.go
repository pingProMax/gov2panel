// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// V2ServerRoute is the golang structure for table v2_server_route.
type V2ServerRoute struct {
	Id          int         `json:"id"           orm:"id"           ` //
	Remarks     string      `json:"remarks"      orm:"remarks"      ` // 备注
	Match       string      `json:"match"        orm:"match"        ` // 规则
	Action      string      `json:"action"       orm:"action"       ` // block|dns
	ActionValue string      `json:"action_value" orm:"action_value" ` //
	CreatedAt   *gtime.Time `json:"created_at"   orm:"created_at"   ` //
	UpdatedAt   *gtime.Time `json:"updated_at"   orm:"updated_at"   ` //
	Enable      int         `json:"enable"       orm:"enable"       ` // 是否启用
}
