package Knowledge

import (
	"gov2panel/internal/dao"
	"gov2panel/internal/logic/cornerstone"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/model/model"
	"gov2panel/internal/service"
)

type sKnowledge struct {
	Cornerstone cornerstone.Cornerstone
}

func init() {
	service.RegisterKnowledge(New())
}

func New() *sKnowledge {
	return &sKnowledge{
		Cornerstone: *cornerstone.NewCornerstoneWithTable(dao.V2Knowledge.Table()),
	}
}

// AE设置
func (s *sKnowledge) AEKnowledge(data *entity.V2Knowledge) (err error) {
	if data.Id != 0 {
		err = s.Cornerstone.UpdateById(data.Id, data)
		return
	}

	err = s.Cornerstone.Save(data)
	return
}

// 删除
func (s *sKnowledge) DelKnowledge(ids []int) error {
	return s.Cornerstone.DelByIds(ids)
}

// 获取所有
func (s *sKnowledge) GetKnowledgeAllList(req entity.V2Knowledge) (m []*entity.V2Knowledge, err error) {
	m = make([]*entity.V2Knowledge, 0)
	db := s.Cornerstone.GetDB().
		WhereLike(dao.V2Knowledge.Columns().Title, "%"+req.Title+"%").
		WhereLike(dao.V2Knowledge.Columns().Category, "%"+req.Category+"%")
	if req.Id != 0 {
		db.Where(dao.V2Knowledge.Columns().Id, req.Id)
	}
	if req.Show != -1 {
		db.Where(dao.V2Knowledge.Columns().Show, req.Show)
	}
	err = db.OrderDesc("order_id").Scan(&m)
	return m, err
}

// 获取显示的数据
func (s *sKnowledge) GetKnowledgeShowList(req entity.V2Knowledge) (data []*model.KnowledgeInfo, err error) {
	data = make([]*model.KnowledgeInfo, 0)
	req.Show = 1

	m := make([]*entity.V2Knowledge, 0)

	db := s.Cornerstone.GetDB().
		WhereLike(dao.V2Knowledge.Columns().Title, "%"+req.Title+"%").
		WhereLike(dao.V2Knowledge.Columns().Category, "%"+req.Category+"%")
	if req.Id != 0 {
		db.Where(dao.V2Knowledge.Columns().Id, req.Id)
	}
	if req.Show != -1 {
		db.Where(dao.V2Knowledge.Columns().Show, req.Show)
	}
	err = db.OrderDesc("order_id").Scan(&m)
	if err != nil {
		return data, err
	}

	for i := 0; i < len(m); i++ {
		repeat := false
		for j := i + 1; j < len(m); j++ {
			if m[i].Category == m[j].Category {
				repeat = true
				break
			}
		}
		if !repeat {
			data = append(
				data,
				&model.KnowledgeInfo{
					Category: m[i].Category,
					Data:     make([]*entity.V2Knowledge, 0),
				},
			)
		}
	}

	for i := 0; i < len(data); i++ {

		for _, v := range m {
			if v.Category == data[i].Category {
				data[i].Data = append(data[i].Data, v)
			}
		}

	}

	return data, err
}
