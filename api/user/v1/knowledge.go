package v1

import (
	"gov2panel/internal/model/entity"
	"gov2panel/internal/model/model"

	"github.com/gogf/gf/v2/frame/g"
)

type KnowledgeReq struct {
	g.Meta `path:"/knowledge" tags:"Knowledge" method:"get,post" summary:"使用文档"`
	entity.V2Knowledge
}
type KnowledgeRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Data   []*model.KnowledgeInfo `json:"data"`
}
