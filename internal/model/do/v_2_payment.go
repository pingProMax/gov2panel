// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// V2Payment is the golang structure of table v2_payment for DAO operations like Where/Data.
type V2Payment struct {
	g.Meta             `orm:"table:v2_payment, do:true"`
	Id                 interface{} //
	Uuid               interface{} // uuid
	Payment            interface{} // 支付类型
	Name               interface{} // 名字
	Icon               interface{} // 图标地址
	Config             interface{} // 配置json数
	NotifyDomain       interface{} // 回调域名
	HandlingFeeFixed   interface{} // 固定手续费
	HandlingFeePercent interface{} // 百分比手续费
	Enable             interface{} // 是否启用
	OrderId            interface{} // 顺序
	CreatedAt          *gtime.Time //
	UpdatedAt          *gtime.Time //
	Remarks            interface{} // 备注
}
