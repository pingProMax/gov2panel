// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// V2ServiceRelay is the golang structure for table v2_service_relay.
type V2ServiceRelay struct {
	Id        int         `json:"id"         orm:"id"         ` //
	Ip        string      `json:"ip"         orm:"ip"         ` // ip
	NameGroup string      `json:"name_group" orm:"name_group" ` // 组名字
	Asn       string      `json:"asn"        orm:"asn"        ` // asn分组 AS9808|AS4134|AS4837
	Show      int         `json:"show"       orm:"show"       ` // 是否启用
	OrderId   int         `json:"order_id"   orm:"order_id"   ` // 顺序
	Remarks   string      `json:"remarks"    orm:"remarks"    ` // 备注
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" ` //
}
