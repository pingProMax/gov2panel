package v1

import (
	"gov2panel/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type SettingReq struct {
	g.Meta `path:"/setting" tags:"Setting" method:"get,post" summary:"系统设置"`
	entity.V2Setting
}
type SettingRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Data   []*entity.V2Setting `json:"data"`
}

type SettingAEReq struct {
	g.Meta `path:"/setting/ae" tags:"Setting" method:"post" summary:"系统设置AE"`
	entity.V2Setting
}
type SettingAERes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type SettingDelReq struct {
	g.Meta `path:"/setting/del" tags:"Setting" method:"post" summary:"系统设置删除"`
	Codes  []string `json:"codes"`
}
type SettingDelRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
