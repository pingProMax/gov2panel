package model

import "gov2panel/internal/model/entity"

type ProxyServiceInfo struct {
	V2Plan         []*entity.V2Plan        `json:"plan"`
	V2Route        []*entity.V2ServerRoute `json:"route"`
	V2ProxyService *entity.V2ProxyService  `json:"service"`
}
