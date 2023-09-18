package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type SubscribeReq struct {
	g.Meta `path:"/subscribe" tags:"Pay" method:"get" summary:"订阅"`
}

type SubscribeRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
