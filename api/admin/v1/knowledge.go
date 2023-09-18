package v1

import (
	"gov2panel/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type KnowledgeReq struct {
	g.Meta `path:"/knowledge" tags:"Knowledge" method:"get,post" summary:"知识库管理"`
	entity.V2Knowledge
}
type KnowledgeRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Data   []*entity.V2Knowledge `json:"data"`
}

type KnowledgeAEReq struct {
	g.Meta `path:"/knowledge/ae" tags:"Knowledge" method:"post" summary:"知识库管理AE"`
	entity.V2Knowledge
}
type KnowledgeAERes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type KnowledgeDelReq struct {
	g.Meta `path:"/knowledge/del" tags:"Knowledge" method:"post" summary:"知识库管理删除"`
	Ids    []int `json:"ids"`
}
type KnowledgeDelRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
