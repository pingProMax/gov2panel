// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// V2CouponUse is the golang structure of table v2_coupon_use for DAO operations like Where/Data.
type V2CouponUse struct {
	g.Meta    `orm:"table:v2_coupon_use, do:true"`
	Id        interface{} //
	CouponId  interface{} //
	UserId    interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	PlanId    interface{} //
}
