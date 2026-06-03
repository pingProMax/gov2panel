package service_relay

import (
	v1 "gov2panel/api/admin/v1"
	"gov2panel/internal/dao"
	"gov2panel/internal/logic/cornerstone"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/service"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
)

type sServerRelay struct {
	Cornerstone cornerstone.Cornerstone
}

func init() {
	service.RegisterServerRelay(New())
}

func New() *sServerRelay {
	return &sServerRelay{
		Cornerstone: *cornerstone.NewCornerstoneWithTable(dao.V2ServiceRelay.Table()),
	}
}

// 获取所有
func (s *sServerRelay) GetServiceRelayList(req *v1.ServiceRelayReq, orderBy, orderDirection string, offset, limit int) (m []*entity.V2ServiceRelay, total int, err error) {

	m = make([]*entity.V2ServiceRelay, 0)
	db := s.Cornerstone.GetDB()

	// 1. 拼接基础查询条件
	db = db.WhereLike(dao.V2ServiceRelay.Columns().NameGroup, "%"+req.NameGroup+"%")
	db = db.WhereLike(dao.V2ServiceRelay.Columns().Ip, "%"+req.Ip+"%")
	db = db.WhereLike(dao.V2ServiceRelay.Columns().Asn, "%"+req.Asn+"%")
	db = db.WhereLike(dao.V2ServiceRelay.Columns().Remarks, "%"+req.Remarks+"%")
	if req.Id != 0 {
		db = db.Where(dao.V2ServiceRelay.Columns().Id, req.Id)
	}
	if req.Show != 0 {
		db = db.Where(dao.V2ServiceRelay.Columns().Show, req.Show)
	}

	// 2. 使用 Clone() 复制一个用于 Count 的对象
	dbCount := db.Clone()

	// 3. 执行列表查询（注意：GoFrame 链式操作建议覆盖赋值，如 db = db.Order...）
	err = db.Order(orderBy, orderDirection).Limit(offset, limit).Scan(&m)
	if err != nil {
		return m, 0, err
	}

	// 4. 执行总数统计
	total, err = dbCount.Count()
	if err != nil {
		return m, 0, err
	}

	return m, total, nil
}

// AE设置
func (s *sServerRelay) AEServiceRelay(data *entity.V2ServiceRelay) (err error) {
	if data.Id != 0 {
		err = s.Cornerstone.UpdateById(data.Id, data)
		return
	}

	// 1. 先把 Windows 的 \r\n 统一替换为 \n
	fixedData := strings.ReplaceAll(data.Ip, "\r\n", "\n")

	// 2. 根据 \n 切割成切片 (Slice)
	ips := strings.Split(fixedData, "\n")

	for _, line := range ips {
		// 建议加上 TrimSpace，防止用户不小心打了空格或由于特殊换行留下的空白
		line = strings.TrimSpace(line)

		// 过滤掉空行
		if line == "" {
			continue
		}

		newData := *data
		newData.Ip = line // 修改新数据的 IP

		err = s.Cornerstone.Save(newData)
		if err != nil {
			return err
		}

	}

	return

}

// 删除
func (s *sServerRelay) DelServiceRelay(ids []int) error {
	return s.Cornerstone.DelByIds(ids)
}

// 批量设置节点显示隐藏状态
func (s *sServerRelay) UpServiceShow(ids []int, show int) (err error) {
	db := s.Cornerstone.GetDB().Data(
		g.Map{
			dao.V2ServiceRelay.Columns().Show: show,
		},
	)
	_, err = db.WhereIn(dao.V2ServiceRelay.Columns().Id, ids).Update()
	return
}

// GetServiceRelayListByShow 根据启用状态获取列表
func (s *sServerRelay) GetServiceRelayListByShow(show int) (m []*entity.V2ServiceRelay, err error) {

	m = make([]*entity.V2ServiceRelay, 0)
	db := s.Cornerstone.GetDB()

	db = db.Where(dao.V2ServiceRelay.Columns().Show, show)

	err = db.Scan(&m)
	if err != nil {
		return m, err
	}

	return m, nil

}
