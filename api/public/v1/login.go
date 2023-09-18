package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta `path:"/login" tags:"Login" method:"get,post" summary:"登录"`
}
type LoginRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}
