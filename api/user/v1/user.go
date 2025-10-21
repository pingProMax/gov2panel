package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserReq struct {
	g.Meta `path:"/user" tags:"User" method:"get,post" summary:"个人中心"`
}
type UserRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type UserUpPasswdReq struct {
	g.Meta    `path:"/up_passwd" tags:"User" method:"post" summary:"修改密码"`
	OldPasswd string `json:"old_passwd"`
	NewPasswd string `json:"new_passwd"`
}
type UserUpPasswdRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type ResetTokenAndUuidReq struct {
	g.Meta `path:"/reset_token_uuid" tags:"User" method:"post" summary:"重置token和uuid"`
}
type ResetTokenAndUuidRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type LogoutReq struct {
	g.Meta `path:"/logout" tags:"User" method:"get" summary:"退出登录"`
}
type LogoutRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
