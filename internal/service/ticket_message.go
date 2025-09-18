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
	ITicketMessage interface {
		// 获取工单信息数据
		GetTicketMessageArrByTicketId(ticketId int) (m []*model.TicketMessageInfo, err error)
		// 获取工单信息数据
		GetTicketMessageArrByTicketIdAndUserId(ticketId int, userId int) (m []*model.TicketMessageInfo, err error)
		// 管理员回复工单
		SaveTicketMessageAdmin(data *entity.V2TicketMessage) (err error)
		// 用户回复工单
		SaveTicketMessageUser(data *entity.V2TicketMessage) (err error)
	}
)

var (
	localTicketMessage ITicketMessage
)

func TicketMessage() ITicketMessage {
	if localTicketMessage == nil {
		panic("implement not found for interface ITicketMessage, forgot register?")
	}
	return localTicketMessage
}

func RegisterTicketMessage(i ITicketMessage) {
	localTicketMessage = i
}
