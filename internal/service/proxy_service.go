// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "gov2panel/api/admin/v1"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/model/model"
)

type (
	IProxyService interface {
		// 获取所有
		GetProxyServiceList(req *v1.ProxyServiceReq, orderBy string, orderDirection string, offset int, limit int) (m []*model.ProxyServiceInfo, total int, err error)
		// 获取所有服务器信息
		GetProxyServiceAllList() (m []*entity.V2ProxyService, err error)
		// AE设置
		AEProxyService(data *entity.V2ProxyService) (err error)
		// 更改服务器地址
		UpProxyServiceIpById(id int, ip string) (err error)
		// 删除
		DelProxyService(ids []int) error
		// 查询服务器中的路由数量 根据路由id
		GetServiceCountByRouteId(routeId []int) (int, error)
		// 查询服务器中的订阅数量 根据订阅id
		GetServiceCountByPlanId(PlanId []int) (int, error)
		// id和type 获取节点信息
		GetServiceAndRouteListById(id int) (data *entity.V2ProxyService, routeList []*entity.V2ServerRoute, err error)
		// id 获取节点信息和订阅信息
		GetServicePlanIdsById(id int) (data *entity.V2ProxyService, planIds []int, err error)
		// id 获取节点信息 和订阅信息
		GetServicePlanListById(id int) (data *entity.V2ProxyService, planList []*entity.V2Plan, err error)
		// planId 获取节点信息
		GetServiceListByPlanIdAndShow1(planId int) (data []*entity.V2ProxyService, err error)
		// 获取节点数量
		GetServiceCount() (data int, err error)
		// 缓存 服务器当前用户数量
		CacheServiceFlow(nodeId int, userTraffic []*model.UserTraffic) (err error)
		// 获取所有服务器当前在线用户数量和服务器最后提交时间
		// map[服务器id][type 1在线数量、2服务器最后提交时间]int
		GetOnlineUserCountAndLastPushAt() (data map[int]map[int]int64, err error)
		// 批量更新节点订阅
		UpBatchPlan(ids []int, planIds string) (err error)
		// 批量更新节点路由
		UpBatchRoute(ids []int, routeIds string) (err error)
	}
)

var (
	localProxyService IProxyService
)

func ProxyService() IProxyService {
	if localProxyService == nil {
		panic("implement not found for interface IProxyService, forgot register?")
	}
	return localProxyService
}

func RegisterProxyService(i IProxyService) {
	localProxyService = i
}
