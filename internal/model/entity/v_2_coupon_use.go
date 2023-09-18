// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// V2CouponUse is the golang structure for table v2_coupon_use.
type V2CouponUse struct {
	Id        int         `json:"id"         ` //
	CouponId  int         `json:"coupon_id"  ` //
	UserId    int         `json:"user_id"    ` //
	CreatedAt *gtime.Time `json:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updated_at" ` //
	PlanId    int         `json:"plan_id"    ` //
}
