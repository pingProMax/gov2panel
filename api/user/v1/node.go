package v1

import (
	"gov2panel/internal/model/model"

	"github.com/gogf/gf/v2/frame/g"
)

type NodeReq struct {
	g.Meta `path:"/node" tags:"Node" method:"get,post" summary:"节点列表页面和api"`
}
type NodeRes struct {
	Data []*model.ProxyServiceSubInfo `json:"data"`
}

type OnlineUserCountAndLastPushAtReq struct {
	g.Meta `path:"/node/online_user_count_and_last_push_at" tags:"Node" method:"post" summary:"获取所有服务器当前在线用户数量和服务器最后提交时间"`
}
type OnlineUserCountAndLastPushAtRes struct {
	Data map[int]map[int]int64 `json:"data"`
}
