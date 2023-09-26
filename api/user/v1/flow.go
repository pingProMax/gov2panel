package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type FlowReq struct {
	g.Meta  `path:"/flow" tags:"Flow" method:"get" summary:"流量"`
	TUserID int
}
type FlowRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Data   []map[string]interface{} `json:"data"`
}
