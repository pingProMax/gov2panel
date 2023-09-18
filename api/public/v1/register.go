package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type RegisterReq struct {
	g.Meta         `path:"/register" tags:"Register" method:"get,post" summary:"注册"`
	UserName       string
	Passwd         string
	CommissionCode string //邀请码
	Code           string
	VerifyCaptcha
}
type RegisterRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}
