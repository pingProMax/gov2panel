package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ApiReq struct {
	g.Meta   `path:"/config" tags:"Api" method:"get" summary:"获取节点信息api"`
	NodeId   int    `json:"node_id"`
	NodeType string `json:"node_type"`
}

type ApiRes map[string]interface{}
