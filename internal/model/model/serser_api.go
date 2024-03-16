package model

import "github.com/gogf/gf/v2/os/gtime"

//后端 上报用户流量用
type UserTraffic struct {
	UID            int
	Email          string
	Upload         int64
	Download       int64
	TransferEnable int64       `json:"transfer_enable"    ` // 流量
	ExpiredAt      *gtime.Time `json:"expired_at"         ` // 到期时间
	GroupId        int         `json:"group_id"           ` // 权限组
	Banned         int         `json:"banned"             ` // 是否禁用
}
