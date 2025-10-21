package model

import (
	"gov2panel/internal/model/entity"

	"github.com/golang-jwt/jwt/v5"
)

type UserInfo struct {
	V2User *entity.V2User `json:"user"`
	V2Plan *entity.V2Plan `json:"plan"`
}

// JWT 令牌的自定义声明
type JWTClaims struct {
	UserName string `json:"user_name"`
	TUserID  int    `json:"t_user_id"`
	TPasswd  string `json:"t_passwd"`
	jwt.RegisteredClaims
}
