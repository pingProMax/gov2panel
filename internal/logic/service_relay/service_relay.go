package service_relay

import (
	v1 "gov2panel/api/admin/v1"
	"gov2panel/internal/dao"
	"gov2panel/internal/logic/cornerstone"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/service"
	"math/rand"
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

// GetServiceRelayListByShow 根据启用状态获取列表
func (s *sServerRelay) GetRandomRelayByFilter(m []*entity.V2ServiceRelay, targetNameGroup string, targetAsn string) string {

	if len(m) == 0 {
		return ""
	}

	// 1. 定义一个临时切片，用来存放所有符合条件的数据指针
	var filtered []*entity.V2ServiceRelay

	// 2. 遍历原始切片，进行条件筛选
	for _, item := range m {
		if item == nil {
			continue
		}
		// 核心逻辑：只有当 NameGroup 和 Asn 都匹配时，才加入候选池
		if item.NameGroup == targetNameGroup && strings.Contains(item.Asn, targetAsn) {
			filtered = append(filtered, item)
		}
	}

	// 3. 如果没有找到任何符合条件的数据，直接返回 nil
	if len(filtered) == 0 {
		return ""
	}

	// 4. 从符合条件的数据池中，随机抽取一条
	// (Go 1.22+ 推荐直接使用 rand.Intn，老版本需要先 rand.Seed)
	randomIndex := rand.Intn(len(filtered))
	return filtered[randomIndex].Ip

}
