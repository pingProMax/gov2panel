// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// V2ServerRoute is the golang structure for table v2_server_route.
type V2ServerRoute struct {
	Id          int         `json:"id"           ` //
	Remarks     string      `json:"remarks"      ` // 备注
	Match       string      `json:"match"        ` // 规则
	Action      string      `json:"action"       ` // block|dns
	ActionValue string      `json:"action_value" ` //
	CreatedAt   *gtime.Time `json:"created_at"   ` //
	UpdatedAt   *gtime.Time `json:"updated_at"   ` //
	Enable      int         `json:"enable"       ` // 是否启用
}
