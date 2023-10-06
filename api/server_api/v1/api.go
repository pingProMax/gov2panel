package v1

import (
	"gov2panel/internal/model/model"

	"github.com/gogf/gf/v2/frame/g"
)

type ConfigReq struct {
	g.Meta   `path:"/config" tags:"Api" method:"get,post" summary:"获取节点信息api"`
	NodeId   int    `json:"node_id"`
	NodeType string `json:"node_type"`
}

type ConfigRes map[string]interface{}

type UserReq struct {
	g.Meta   `path:"/user" tags:"Api" method:"get,post" summary:"获取节点用户"`
	NodeId   int    `json:"node_id"`
	NodeType string `json:"node_type"`
}

type UserRes struct {
	Users []map[string]interface{} `json:"users"`
}

type PushReq struct {
	g.Meta `path:"/push" tags:"Api" method:"get,post" summary:"报告用户流量"`
	Data   []model.UserTraffic
}

type PushRes struct {
}

type ChangeIPReq struct {
	g.Meta `path:"/change_id" tags:"Api" method:"get,post" summary:"更改服务器iip"`
	NodeId int    `json:"node_id"`
	Ip     string `json:"ip"`
}

type ChangeIPRes struct {
}
