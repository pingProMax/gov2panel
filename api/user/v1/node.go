package v1

import (
	"gov2panel/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type NodeReq struct {
	g.Meta `path:"/node" tags:"Node" method:"get,post" summary:"节点列表页面和api"`
}
type NodeRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Data   []*entity.V2ProxyService `json:"data"`
}
