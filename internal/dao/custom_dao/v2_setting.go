package custom_dao

import (
	"gov2panel/internal/dao"
	"gov2panel/internal/model/entity"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

//系统数据库 带缓存.
//查询缓存 文档
//gdb支持对查询结果的缓存处理，常用于多读少写的查询缓存场景，并支持手动的缓存清理。需要注意的是，查询缓存仅支持链式操作，且在事务操作下不可用。

func getSettingDB() *gdb.Model {
	return dao.V2Setting.DB().Model(dao.V2Setting.Table())
}

// 添加设置
func SaveSetting(data *entity.V2Setting) (err error) {
	_, err = getSettingDB().Data(data).Cache(gdb.CacheOption{
		Duration: -1, //清理缓存
		Name:     "v2_setting_all",
		Force:    false,
	}).Insert()
	return err
}

// 更新设置
func UpSetting(data *entity.V2Setting) (err error) {
	_, err = getSettingDB().Data(data).FieldsEx(dao.V2Setting.Columns().CreatedAt).Cache(gdb.CacheOption{
		Duration: -1, //清理缓存
		Name:     "v2_setting_all",
		Force:    false,
	}).Where(dao.V2Setting.Columns().Code, data.Code).Update()
	return err
}

// 删除设置
func DelSetting(codes []string) (err error) {
	_, err = getSettingDB().WhereIn(dao.V2Setting.Columns().Code, codes).Cache(gdb.CacheOption{
		Duration: -1, //清理缓存
		Name:     "v2_setting_all",
		Force:    false,
	}).Delete()
	return err
}

// 获取所有设置
func GetSettingAllMap() (m map[string]*g.Var, err error) {
	m = make(map[string]*g.Var)

	result, err := getSettingDB().Cache(gdb.CacheOption{
		Duration: 0, //永不过期
		Name:     "v2_setting_all",
		Force:    false,
	}).All()
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		m[v["code"].String()] = gvar.New(v["value"])
	}

	return m, nil
}

// 获取所有设置
func GetSettingAllList(req entity.V2Setting) (m []*entity.V2Setting, err error) {
	m = make([]*entity.V2Setting, 0)

	err = getSettingDB().
		WhereLike(dao.V2Setting.Columns().Code, "%"+req.Code+"%").
		WhereLike(dao.V2Setting.Columns().Value, "%"+req.Value+"%").
		WhereLike(dao.V2Setting.Columns().Remarks, "%"+req.Remarks+"%").
		OrderDesc("order_id").Scan(&m)
	return m, err
}

// 获取指定code值
func GetSettingByCode(code string) (items *entity.V2Setting, err error) {
	err = getSettingDB().Where(dao.V2Setting.Columns().Code, code).Scan(&items)
	return items, err
}
