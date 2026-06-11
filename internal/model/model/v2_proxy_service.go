package model

import "gov2panel/internal/model/entity"

type ProxyServiceInfo struct {
	V2Plan         []*entity.V2Plan        `json:"plan"`
	V2Route        []*entity.V2ServerRoute `json:"route"`
	V2ProxyService *entity.V2ProxyService  `json:"service"`
}

// 流量排行榜用
type ProxyServiceFlow struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Flow int64  `json:"flow"`
}

// 带url，方便客户端导入
type ProxyServiceSubInfo struct {
	V2ProxyService *entity.V2ProxyService `json:"service"`
	Url            string                 `json:"url"`
}
