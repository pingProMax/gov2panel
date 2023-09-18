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
	ITicket interface {
		// 获取工单数据
		GetUserList(req *entity.V2Ticket, userName string, orderBy, orderDirection string, offset, limit int) (data []*model.TicketInfo, totle int, err error)
		// AE设置
		AETicket(data *entity.V2Ticket) (err error)
		// 删除
		DelTicket(ids []int) error
		// 关闭
		CloseTicket(ids []int) (err error)
		// 用户关闭工单
		CloseTicketByUserIdAndId(ids []int, userId int) (err error)
		// 获取工单 根据id和用户id
		GetTicketByIdAndUserId(ticketId, userId int) (data *entity.V2Ticket, err error)
		// 更新工单 根据id
		UpTicketStatusAndReplyStatusById(ticketId, status, reply_status int) (err error)
		// 获取所有
		GetTicketList(req *entity.V2Ticket, userName string, orderBy, orderDirection string, offset, limit int) (m []*model.TicketInfo, total int, err error)
	}
)

var (
	localTicket ITicket
)

func Ticket() ITicket {
	if localTicket == nil {
		panic("implement not found for interface ITicket, forgot register?")
	}
	return localTicket
}

func RegisterTicket(i ITicket) {
	localTicket = i
}
