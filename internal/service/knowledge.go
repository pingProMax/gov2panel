// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"gov2panel/internal/model/entity"
	"gov2panel/internal/model/model"
)

type (
	IKnowledge interface {
		// AE设置
		AEKnowledge(data *entity.V2Knowledge) (err error)
		// 删除
		DelKnowledge(ids []int) error
		// 获取所有
		GetKnowledgeAllList(req entity.V2Knowledge) (m []*entity.V2Knowledge, err error)
		// 获取显示的数据
		GetKnowledgeShowList(req entity.V2Knowledge) (data []*model.KnowledgeInfo, err error)
	}
)

var (
	localKnowledge IKnowledge
)

func Knowledge() IKnowledge {
	if localKnowledge == nil {
		panic("implement not found for interface IKnowledge, forgot register?")
	}
	return localKnowledge
}

func RegisterKnowledge(i IKnowledge) {
	localKnowledge = i
}
