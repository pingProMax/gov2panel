package model

import (
	"gov2panel/internal/model/entity"

	"github.com/golang-jwt/jwt/v5"
)

type UserInfo struct {
	V2User *entity.V2User `json:"user"`
	V2Plan *entity.V2Plan `json:"plan"`
}

type UserDay7Flow struct {
	UserId int    `json:"user_id"`
	Date   string `json:"date"`
	Flow   int64  `json:"flow"`
}

// JWT 令牌的自定义声明
type JWTClaims struct {
	UserName string `json:"user_name"`
	TUserID  int    `json:"t_user_id"`
	TPasswd  string `json:"t_passwd"`
	jwt.RegisteredClaims
}
