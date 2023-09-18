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
		GetProxyServiceList(req *v1.ProxyServiceReq, orderBy, orderDirection string, offset, limit int) (m []*model.ProxyServiceInfo, total int, err error)
		// AE设置
		AEProxyService(data *entity.V2ProxyService) (err error)
		// 删除
		DelProxyService(ids []int) error
		// 查询服务器中的路由数量 根据路由id
		GetServiceCountByRouteId(routeId []int) (int, error)
		// 查询服务器中的订阅数量 根据订阅id
		GetServiceCountByPlanId(PlanId []int) (int, error)
		// id和type 获取节点信息
		GetServiceById(id int) (data *entity.V2ProxyService, planList []*entity.V2Plan, routeList []*entity.V2ServerRoute, err error)
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
