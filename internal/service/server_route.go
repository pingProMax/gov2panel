// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "gov2panel/api/admin/v1"
	"gov2panel/internal/model/entity"
)

type (
	IServerRoute interface {
		// 获取所有
		GetServerRouteList(req *v1.ServerRouteReq, orderBy string, orderDirection string, offset int, limit int) (m []*entity.V2ServerRoute, total int, err error)
		// 获取所有
		ServerRouteAll() (m []*entity.V2ServerRoute, err error)
		// AE设置
		AEServerRoute(data *entity.V2ServerRoute) (err error)
		// 删除
		DelServerRoute(ids []int) error
	}
)

var (
	localServerRoute IServerRoute
)

func ServerRoute() IServerRoute {
	if localServerRoute == nil {
		panic("implement not found for interface IServerRoute, forgot register?")
	}
	return localServerRoute
}

func RegisterServerRoute(i IServerRoute) {
	localServerRoute = i
}
