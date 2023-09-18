package proxy_service

import (
	"encoding/json"
	v1 "gov2panel/api/admin/v1"
	"gov2panel/internal/dao"
	"gov2panel/internal/logic/cornerstone"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/model/model"
	"gov2panel/internal/service"
	"strconv"
)

type sProxyService struct {
	Cornerstone cornerstone.Cornerstone
}

func init() {
	service.RegisterProxyService(New())
}

func New() *sProxyService {
	return &sProxyService{
		Cornerstone: *cornerstone.NewCornerstoneWithTable(dao.V2ProxyService.Table()),
	}
}

// 获取所有
func (s *sProxyService) GetProxyServiceList(req *v1.ProxyServiceReq, orderBy, orderDirection string, offset, limit int) (m []*model.ProxyServiceInfo, total int, err error) {

	m = make([]*model.ProxyServiceInfo, 0)
	db := s.Cornerstone.GetDB()

	db.WhereLike(dao.V2ProxyService.Columns().Agreement, "%"+req.Agreement+"%")
	db.WhereLike(dao.V2ProxyService.Columns().ServiceJson, "%"+req.ServiceJson+"%")
	db.WhereLike(dao.V2ProxyService.Columns().Name, "%"+req.Name+"%")
	db.WhereLike(dao.V2ProxyService.Columns().Host, "%"+req.Host+"%")
	db.WhereLike(dao.V2ProxyService.Columns().Port, "%"+req.Port+"%")
	if req.Id != 0 {
		db.Where(dao.V2ProxyService.Columns().Id, req.Id)
	}
	if req.Show != -1 {
		db.Where(dao.V2ProxyService.Columns().Show, req.Show)
	}

	if req.PlanId != "-1" {
		db.WhereLike(dao.V2ProxyService.Columns().PlanId, "%\""+req.PlanId+"\"%")
	}

	dbC := *db
	dbCCount := &dbC

	err = db.Order(orderBy, orderDirection).Limit(offset, limit).ScanList(&m, "V2ProxyService")
	if err != nil {
		return m, 0, err
	}

	total, err = dbCCount.Count()
	if err != nil {
		return m, 0, err
	}

	if total > 0 {

		planList, err := service.Plan().GetPlanShowAndResetTrafficMethod1List()
		if err != nil {
			return m, 0, err
		}

		for i := 0; i < len(m); i++ {

			// 定义一个切片来存储解析后的数据
			var stringSlice []string

			json.Unmarshal([]byte(m[i].V2ProxyService.PlanId), &stringSlice)
			m[i].V2Plan = make([]*entity.V2Plan, 0)

			for _, v := range stringSlice {

				for _, vv := range planList {
					if v == strconv.Itoa(vv.Id) {
						m[i].V2Plan = append(m[i].V2Plan, vv)
						break
					}
				}
			}

		}

		routeList, err := service.ServerRoute().ServerRouteAll()
		if err != nil {
			return m, 0, err
		}

		for i := 0; i < len(m); i++ {

			// 定义一个切片来存储解析后的数据
			var stringSlice []string

			json.Unmarshal([]byte(m[i].V2ProxyService.RouteId), &stringSlice)
			m[i].V2Route = make([]*entity.V2ServerRoute, 0)

			for _, v := range stringSlice {

				for _, vv := range routeList {
					if v == strconv.Itoa(vv.Id) {
						m[i].V2Route = append(m[i].V2Route, vv)
						break
					}
				}
			}

		}

	}

	return m, total, err
}

// AE设置
func (s *sProxyService) AEProxyService(data *entity.V2ProxyService) (err error) {
	if data.Id != 0 {
		err = s.Cornerstone.UpdateById(data.Id, data)
		return
	}

	err = s.Cornerstone.Save(data)
	return
}

// 删除
func (s *sProxyService) DelProxyService(ids []int) error {
	return s.Cornerstone.DelByIds(ids)
}

// 查询服务器中的路由数量 根据路由id
func (s *sProxyService) GetServiceCountByRouteId(routeId []int) (int, error) {
	sqlStr := ""

	for i, v := range routeId {
		if i != 0 {
			sqlStr = sqlStr + " or "
		}

		sqlStr = sqlStr + dao.V2ProxyService.Columns().RouteId + " like " + "'%\"" + strconv.Itoa(v) + "\"%'"
	}

	return s.Cornerstone.GetDB().Where(sqlStr).Count()
}

// 查询服务器中的订阅数量 根据订阅id
func (s *sProxyService) GetServiceCountByPlanId(PlanId []int) (int, error) {
	sqlStr := ""

	for i, v := range PlanId {
		if i != 0 {
			sqlStr = sqlStr + " or "
		}

		sqlStr = sqlStr + dao.V2ProxyService.Columns().PlanId + " like " + "'%\"" + strconv.Itoa(v) + "\"%'"
	}

	return s.Cornerstone.GetDB().Where(sqlStr).Count()
}

// id和type 获取节点信息
func (s *sProxyService) GetServiceById(id int) (data *entity.V2ProxyService, planList []*entity.V2Plan, routeList []*entity.V2ServerRoute, err error) {
	data = new(entity.V2ProxyService)
	err = s.Cornerstone.GetOneById(id, data)
	if err != nil {
		return
	}

	var planIdStrSlice []string
	err = json.Unmarshal([]byte(data.PlanId), &planIdStrSlice)

	if len(planIdStrSlice) > 0 {
		planList = make([]*entity.V2Plan, 0)
		err = s.Cornerstone.GetDBT(dao.V2Plan.Table()).Where(dao.V2Plan.Columns().Id, planIdStrSlice).Scan(&planList)
		if err != nil {
			return
		}
	}

	var routeIdStrSlice []string
	json.Unmarshal([]byte(data.RouteId), &routeIdStrSlice)

	if len(routeIdStrSlice) > 0 {
		routeList = make([]*entity.V2ServerRoute, 0)
		err = s.Cornerstone.GetDBT(dao.V2ServerRoute.Table()).Where(dao.V2ServerRoute.Columns().Id, routeIdStrSlice).Scan(&routeList)
		if err != nil {
			return
		}
	}

	return
}
