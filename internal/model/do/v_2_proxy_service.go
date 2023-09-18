// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// V2ProxyService is the golang structure of table v2_proxy_service for DAO operations like Where/Data.
type V2ProxyService struct {
	g.Meta      `orm:"table:v2_proxy_service, do:true"`
	Id          interface{} //
	Agreement   interface{} // 协议
	ServiceJson interface{} // 服务器json数据
	Name        interface{} // 显示名称
	PlanId      interface{} // 所属订阅组
	Show        interface{} // 是否显示
	Host        interface{} // 服务器地址
	Port        interface{} // 服务器端口
	Rate        interface{} // 倍率
	OrderId     interface{} // 顺序
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	RouteId     interface{} // 所属路由组
}
