// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// V2Payment is the golang structure for table v2_payment.
type V2Payment struct {
	Id                 int         `json:"id"                   ` //
	Uuid               string      `json:"uuid"                 ` // uuid
	Payment            string      `json:"payment"              ` // 支付类型
	Name               string      `json:"name"                 ` // 名字
	Icon               string      `json:"icon"                 ` // 图标地址
	Config             string      `json:"config"               ` // 配置json数
	NotifyDomain       string      `json:"notify_domain"        ` // 回调域名
	HandlingFeeFixed   float64     `json:"handling_fee_fixed"   ` // 固定手续费
	HandlingFeePercent int         `json:"handling_fee_percent" ` // 百分比手续费
	Enable             int         `json:"enable"               ` // 是否启用
	OrderId            int         `json:"order_id"             ` // 顺序
	CreatedAt          *gtime.Time `json:"created_at"           ` //
	UpdatedAt          *gtime.Time `json:"updated_at"           ` //
	Remarks            string      `json:"remarks"              ` // 备注
}
