package model

import "gov2panel/internal/model/entity"

type UserInfo struct {
	V2User *entity.V2User `json:"user"`
	V2Plan *entity.V2Plan `json:"plan"`
}
