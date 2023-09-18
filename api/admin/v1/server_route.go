package v1

import (
	"gov2panel/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type ServerRouteReq struct {
	g.Meta `path:"/service_route" tags:"ServerRoute" method:"get,post" summary:"服务器路由"`
	SortOrder
	OffsetLimit
	entity.V2ServerRoute
}
type ServerRouteRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Data   []*entity.V2ServerRoute `json:"data"`
	Totle  int                     `json:"totle"`
}
type ServerRouteAEReq struct {
	g.Meta `path:"/service_route/ae" tags:"ServerRoute" method:"post" summary:"AE"`
	entity.V2ServerRoute
}
type ServerRouteAERes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type ServerRouteDelReq struct {
	g.Meta `path:"/service_route/del" tags:"ServerRoute" method:"post" summary:"删除"`
	Ids    []int `json:"ids"`
}
type ServerRouteDelRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type ServerRouteAllReq struct {
	g.Meta `path:"/service_route/all" tags:"ServerRoute" method:"post" summary:"获取所有"`
}
type ServerRouteAllRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Data   []*entity.V2ServerRoute `json:"data"`
}
