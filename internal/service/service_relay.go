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
	IServerRelay interface {
		// 获取所有
		GetServiceRelayList(req *v1.ServiceRelayReq, orderBy string, orderDirection string, offset int, limit int) (m []*entity.V2ServiceRelay, total int, err error)
		// AE设置
		AEServiceRelay(data *entity.V2ServiceRelay) (err error)
		// 删除
		DelServiceRelay(ids []int) error
		// 批量设置节点显示隐藏状态
		UpServiceShow(ids []int, show int) (err error)
		// GetServiceRelayListByShow 获取
		GetServiceRelayListByShow(show int) (m []*entity.V2ServiceRelay, err error)
	}
)

var (
	localServerRelay IServerRelay
)

func ServerRelay() IServerRelay {
	if localServerRelay == nil {
		panic("implement not found for interface IServerRelay, forgot register?")
	}
	return localServerRelay
}

func RegisterServerRelay(i IServerRelay) {
	localServerRelay = i
}
