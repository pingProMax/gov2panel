// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"gov2panel/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type (
	ISetting interface {
		// AE设置
		AESetting(data *entity.V2Setting) error
		// 添加设置
		SaveSetting(data *entity.V2Setting) error
		// 删除设置
		DelSetting(codes []string) error
		// 获取所有设置
		GetSettingAllMap() (map[string]*g.Var, error)
		// 获取所有设置
		GetSettingAllList(req entity.V2Setting) ([]*entity.V2Setting, error)
	}
)

var (
	localSetting ISetting
)

func Setting() ISetting {
	if localSetting == nil {
		panic("implement not found for interface ISetting, forgot register?")
	}
	return localSetting
}

func RegisterSetting(i ISetting) {
	localSetting = i
}
