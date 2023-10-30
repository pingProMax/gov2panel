package proxy_service

import (
	"encoding/json"
	"fmt"
	v1 "gov2panel/api/admin/v1"
	"gov2panel/internal/dao"
	"gov2panel/internal/logic/cornerstone"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/model/model"
	"gov2panel/internal/service"
	"gov2panel/internal/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
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

// 获取所有服务器信息
func (s *sProxyService) GetProxyServiceAllList() (m []*entity.V2ProxyService, err error) {

	m = make([]*entity.V2ProxyService, 0)
	err = s.Cornerstone.GetDB().Scan(&m)
	return m, err
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

// 更改服务器地址
func (s *sProxyService) UpProxyServiceIpById(id int, ip string) (err error) {
	_, err = s.Cornerstone.GetDB().Data(g.Map{dao.V2ProxyService.Columns().Host: ip}).Where(dao.V2ProxyService.Columns().Id, id).Update()
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
func (s *sProxyService) GetServiceAndRouteListById(id int) (data *entity.V2ProxyService, routeList []*entity.V2ServerRoute, err error) {
	data = new(entity.V2ProxyService)
	err = s.Cornerstone.GetOneById(id, data)
	if err != nil {
		return
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

// id 获取节点信息和订阅信息
func (s *sProxyService) GetServicePlanIdsById(id int) (data *entity.V2ProxyService, planIds []int, err error) {
	data = new(entity.V2ProxyService)
	err = s.Cornerstone.GetOneById(id, data)
	if err != nil {
		return
	}

	var planIdStrSlice []string
	err = json.Unmarshal([]byte(data.PlanId), &planIdStrSlice)

	planIds = make([]int, len(planIdStrSlice))

	// 遍历 []string 切片并将每个字符串转换为整数
	for i, str := range planIdStrSlice {
		num, err := strconv.Atoi(str)
		if err != nil {
			continue
		}
		planIds[i] = num
	}

	return
}

// id 获取节点信息 和订阅信息
func (s *sProxyService) GetServicePlanListById(id int) (data *entity.V2ProxyService, planList []*entity.V2Plan, err error) {
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

	return
}

// planId 获取节点信息
func (s *sProxyService) GetServiceListByPlanIdAndShow1(planId int) (data []*entity.V2ProxyService, err error) {
	whereSqlStr := dao.V2ProxyService.Columns().PlanId + " like " + "'%\"" + strconv.Itoa(planId) + "\"%'"
	data = make([]*entity.V2ProxyService, 0)

	err = s.Cornerstone.GetDB().Where(whereSqlStr).Where(dao.V2ProxyService.Columns().Show, 1).OrderDesc(dao.V2ProxyService.Columns().OrderId).Scan(&data)
	return
}

// 获取节点数量
func (s *sProxyService) GetServiceCount() (data int, err error) {

	data, err = s.Cornerstone.GetDB().Count()
	return
}

// 缓存 服务器当前用户数量
func (s *sProxyService) CacheServiceFlow(nodeId int, userTraffic []model.UserTraffic) (err error) {
	timeNow := time.Now()
	ctx := gctx.New()

	//服务器当前用户在线数量
	err = gcache.Set(ctx, fmt.Sprintf("SERVER_%s_ONLINE_USER", strconv.Itoa(nodeId)), len(userTraffic), 3600*time.Second)
	if err != nil {
		return
	}

	//服务器最后提交数据时间
	err = gcache.Set(ctx, fmt.Sprintf("SERVER_%s_LAST_PUSH_AT", strconv.Itoa(nodeId)), timeNow.Unix(), 0)
	if err != nil {
		return
	}

	var upload int64
	var download int64
	for _, v := range userTraffic {
		upload = upload + v.Upload
		download = download + v.Download
	}

	//服务器当天的流量使用情况 (记录两天的)
	cacheServerFlowKey := fmt.Sprintf("SERVER_%s_%s_FLOW", strconv.Itoa(nodeId), utils.GetDateNowStr())
	serverFlow, err := gcache.Get(ctx, cacheServerFlowKey)
	if err != nil {
		return
	}

	err = gcache.Set(ctx, cacheServerFlowKey, serverFlow.Int64()+upload+download, 49*time.Hour)
	if err != nil {
		return
	}

	return
}

// 获取所有服务器当前在线用户数量和服务器最后提交时间
// map[服务器id][type 1在线数量、2服务器最后提交时间]int
func (s *sProxyService) GetOnlineUserCountAndLastPushAt() (data map[int]map[int]int64, err error) {
	data = make(map[int]map[int]int64, 0)

	ctx := gctx.New()

	cacheKeyS, err := gcache.KeyStrings(ctx)
	if err != nil {
		return data, err
	}

	for _, v := range cacheKeyS {
		if strings.HasPrefix(v, "SERVER_") && strings.HasSuffix(v, "_ONLINE_USER") {
			idStr := strings.ReplaceAll(v, "SERVER_", "")
			idStr = strings.ReplaceAll(idStr, "_ONLINE_USER", "")
			id := gconv.Int(idStr)
			onlineUser, err := gcache.Get(ctx, v)
			if err != nil {
				return data, err
			}

			if data[id] == nil {
				data[id] = make(map[int]int64)
			}
			data[id][1] = onlineUser.Int64()
		} else if strings.HasPrefix(v, "SERVER_") && strings.HasSuffix(v, "_LAST_PUSH_AT") {
			idStr := strings.ReplaceAll(v, "SERVER_", "")
			idStr = strings.ReplaceAll(idStr, "_LAST_PUSH_AT", "")
			id := gconv.Int(idStr)
			lastPushAt, err := gcache.Get(ctx, v)
			if err != nil {
				return data, err
			}

			if data[id] == nil {
				data[id] = make(map[int]int64)
			}
			data[id][2] = lastPushAt.Int64()
		}
	}

	return data, nil
}
