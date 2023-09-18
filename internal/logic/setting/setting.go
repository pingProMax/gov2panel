package setting

import (
	"gov2panel/internal/dao/custom_dao"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type sSetting struct{}

func init() {
	service.RegisterSetting(New())
}

func New() *sSetting {
	return &sSetting{}
}

// AE设置
func (s *sSetting) AESetting(data *entity.V2Setting) error {
	v2Setting, err := custom_dao.GetSettingByCode(data.Code)
	if err != nil {
		return err
	}
	if v2Setting == nil {
		return custom_dao.SaveSetting(data)
	}

	return custom_dao.UpSetting(data)
}

// 添加设置
func (s *sSetting) SaveSetting(data *entity.V2Setting) error {
	return custom_dao.SaveSetting(data)
}

// 删除设置
func (s *sSetting) DelSetting(codes []string) error {
	return custom_dao.DelSetting(codes)
}

// 获取所有设置
func (s *sSetting) GetSettingAllMap() (map[string]*g.Var, error) {
	return custom_dao.GetSettingAllMap()
}

// 获取所有设置
func (s *sSetting) GetSettingAllList(req entity.V2Setting) ([]*entity.V2Setting, error) {
	return custom_dao.GetSettingAllList(req)
}
