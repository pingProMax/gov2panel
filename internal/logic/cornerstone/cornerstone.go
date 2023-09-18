package cornerstone

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type Cornerstone struct {
	Table string //表
}

// 通用 增删改查
func NewCornerstoneWithTable(table string) *Cornerstone {
	return &Cornerstone{
		Table: table,
	}
}

// 账号查询用户
func (s *Cornerstone) GetDB() *gdb.Model {
	return g.Model(s.Table)
}

// 账号查询用户
func (s *Cornerstone) GetDBT(table string) *gdb.Model {
	return g.Model(table)
}

func (s *Cornerstone) Check() {
	if s.Table == "" {
		panic("table is empty")
	}
}

// 获取一条数据根据id
func (s *Cornerstone) GetOneById(id int, resultData interface{}) (err error) {
	s.Check()
	err = g.Model(s.Table).Where("id", id).Scan(resultData)
	return
}

// 删除
func (s *Cornerstone) DelByIds(ids []int) (err error) {
	s.Check()
	_, err = g.Model(s.Table).WhereIn("id", ids).Delete()
	return err
}

// 添加
func (s *Cornerstone) Save(data interface{}) (err error) {
	_, err = g.Model(s.Table).Data(data).Insert()
	return err
}

// 更新
func (s *Cornerstone) UpdateById(id int, data interface{}) (err error) {
	_, err = g.Model(s.Table).Data(data).FieldsEx("created_at").Where("id", id).Update()
	return err
}
