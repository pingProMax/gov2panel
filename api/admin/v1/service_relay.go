package v1

import (
	"gov2panel/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type ServiceRelayReq struct {
	g.Meta `path:"/service_relay" tags:"ServiceRelay" method:"get,post" summary:"中继服务器管理"`
	SortOrder
	OffsetLimit
	entity.V2ServiceRelay
}
type ServiceRelayRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Data   []*entity.V2ServiceRelay `json:"data"`
	Totle  int                      `json:"totle"`
}

type ServiceRelayAEReq struct {
	g.Meta `path:"/service_relay/ae" tags:"ServiceRelay" method:"post" summary:"中继服务器管理AE"`
	entity.V2ServiceRelay
}
type ServiceRelayAERes struct {
}

type ServiceRelayDelReq struct {
	g.Meta `path:"/service_relay/del" tags:"ServiceRelay" method:"post" summary:"中继服务器管理删除"`
	Ids    []int `json:"ids"`
}
type ServiceRelayDelRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type ServiceRelayShowReq struct {
	g.Meta `path:"/service_relay/show" tags:"ServiceRelay" method:"post" summary:"更新服务器显示隐藏"`
	Ids    []int `json:"ids"`
	Show   int   `json:"show"` // 1启用 -1停用
}
type ServiceRelayShowRes struct {
}
