package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type SubscribeReq struct {
	g.Meta       `path:"/api/subscribe" tags:"Pay" method:"get" summary:"订阅"`
	Token        string `json:"token"`
	Flag         string `json:"flag"`
	FlagInfoHide bool   `json:"flag_info_hide"`
}

type SubscribeRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
