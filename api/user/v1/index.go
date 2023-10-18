package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type IndexReq struct {
	g.Meta `path:"/" tags:"Index" method:"get,post" summary:"用户中心"`
}
type IndexRes struct {
	g.Meta         `mime:"text/html" example:"string"`
	Token          string      `json:"token"`               //订阅token
	TransferEnable int64       `json:"transfer_enable"    ` //总流量
	U              int64       `json:"u"`
	D              int64       `json:"d"`
	ExpiredAt      *gtime.Time `json:"expired_at"         ` //到期时间
	PlanName       string      `json:"plan_name"`           //订阅名
	UserName       string      `json:"user_name"`           //用户名
}

type AppBulletinReq struct {
	g.Meta `path:"/app_bulletin" tags:"Index" method:"get,post" summary:"app 公告"`
}
type AppBulletinRes struct {
	Data string `json:"data"` //App 公告信息
}
