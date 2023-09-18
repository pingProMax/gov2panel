package server_route

import (
	"errors"
	v1 "gov2panel/api/admin/v1"
	"gov2panel/internal/dao"
	"gov2panel/internal/logic/cornerstone"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/service"
)

type sServerRoute struct {
	Cornerstone cornerstone.Cornerstone
}

func init() {
	service.RegisterServerRoute(New())
}

func New() *sServerRoute {
	return &sServerRoute{
		Cornerstone: *cornerstone.NewCornerstoneWithTable(dao.V2ServerRoute.Table()),
	}
}

// 获取所有
func (s *sServerRoute) GetServerRouteList(req *v1.ServerRouteReq, orderBy, orderDirection string, offset, limit int) (m []*entity.V2ServerRoute, total int, err error) {

	m = make([]*entity.V2ServerRoute, 0)
	db := s.Cornerstone.GetDB()

	db.WhereLike(dao.V2ServerRoute.Columns().Remarks, "%"+req.Remarks+"%")
	if req.Id != 0 {
		db.Where(dao.V2ServerRoute.Columns().Id, req.Id)
	}

	if req.Enable != -1 {
		db.Where(dao.V2ServerRoute.Columns().Enable, req.Enable)
	}
	if req.Action != "" {
		db.Where(dao.V2ServerRoute.Columns().Action, req.Action)
	}

	dbC := *db
	dbCCount := &dbC

	err = db.Order(orderBy, orderDirection).Limit(offset, limit).Scan(&m)
	if err != nil {
		return m, 0, err
	}

	total, err = dbCCount.Count()
	if err != nil {
		return m, 0, err
	}

	return m, total, err
}

// 获取所有
func (s *sServerRoute) ServerRouteAll() (m []*entity.V2ServerRoute, err error) {

	m = make([]*entity.V2ServerRoute, 0)
	db := s.Cornerstone.GetDB()

	err = db.Order("id", "desc").Scan(&m)
	if err != nil {
		return m, err
	}

	return m, err
}

// AE设置
func (s *sServerRoute) AEServerRoute(data *entity.V2ServerRoute) (err error) {
	if data.Id != 0 {
		err = s.Cornerstone.UpdateById(data.Id, data)
		return
	}

	err = s.Cornerstone.Save(data)
	return
}

// 删除
func (s *sServerRoute) DelServerRoute(ids []int) error {
	serverCount, err := service.ProxyService().GetServiceCountByRouteId(ids)
	if err != nil {
		return err
	}

	if serverCount > 0 {
		return errors.New("路由有节点在使用，无法删除！")
	}

	return s.Cornerstone.DelByIds(ids)
}
