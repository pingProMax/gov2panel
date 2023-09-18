package v1

import (
	"gov2panel/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type UserReq struct {
	g.Meta `path:"/user" tags:"User" method:"get,post" summary:"个人中心"`
	entity.V2User
}
type UserRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Data   []*entity.V2User `json:"data"`
}

type UserUpPasswdReq struct {
	g.Meta    `path:"/up_passwd" tags:"User" method:"post" summary:"修改密码"`
	TUserID   int
	OldPasswd string `json:"old_passwd"`
	NewPasswd string `json:"new_passwd"`
}
type UserUpPasswdRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
