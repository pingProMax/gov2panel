// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// V2ProxyService is the golang structure for table v2_proxy_service.
type V2ProxyService struct {
	Id          int         `json:"id"           ` //
	Agreement   string      `json:"agreement"    ` // 协议
	ServiceJson string      `json:"service_json" ` // 服务器json数据
	Name        string      `json:"name"         ` // 显示名称
	PlanId      string      `json:"plan_id"      ` // 所属订阅组
	Show        int         `json:"show"         ` // 是否显示
	Host        string      `json:"host"         ` // 服务器地址
	Port        string      `json:"port"         ` // 服务器端口
	Rate        int         `json:"rate"         ` // 倍率
	OrderId     int         `json:"order_id"     ` // 顺序
	CreatedAt   *gtime.Time `json:"created_at"   ` //
	UpdatedAt   *gtime.Time `json:"updated_at"   ` //
	RouteId     string      `json:"route_id"     ` // 所属路由组
}
