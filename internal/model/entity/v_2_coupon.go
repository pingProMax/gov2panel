// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// V2Coupon is the golang structure for table v2_coupon.
type V2Coupon struct {
	Id               int         `json:"id"                  ` //
	Code             string      `json:"code"                ` // 优惠码
	Name             string      `json:"name"                ` // 名称
	Type             int         `json:"type"                ` // 类型 1金额优惠 2百分比优惠
	Value            float64     `json:"value"               ` // 优惠多少
	Enable           int         `json:"enable"              ` // 是否启用
	LimitUse         int         `json:"limit_use"           ` // 每个用户可使用次数
	LimitUseWithUser int         `json:"limit_use_with_user" ` // 最大使用次数
	LimitPlanId      int         `json:"limit_plan_id"       ` // 指定订阅
	StartedAt        *gtime.Time `json:"started_at"          ` // 有效期开始时间
	EndedAt          *gtime.Time `json:"ended_at"            ` // 有效期结束时间
	CreatedAt        *gtime.Time `json:"created_at"          ` //
	UpdatedAt        *gtime.Time `json:"updated_at"          ` //
	Remarks          string      `json:"remarks"             ` // 备注
}
