// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// V2Setting is the golang structure for table v2_setting.
type V2Setting struct {
	Code      string      `json:"code"       ` //
	Value     string      `json:"value"      ` //
	OrderId   int         `json:"order_id"   ` // 顺序
	Remarks   string      `json:"remarks"    ` // 备注
	CreatedAt *gtime.Time `json:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updated_at" ` //
}
