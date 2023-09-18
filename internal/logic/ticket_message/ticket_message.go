package ticket_message

import (
	"errors"
	"gov2panel/internal/dao"
	"gov2panel/internal/logic/cornerstone"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/model/model"
	"gov2panel/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
)

type sTicketMessage struct {
	Cornerstone cornerstone.Cornerstone
}

func init() {
	service.RegisterTicketMessage(New())
}

func New() *sTicketMessage {
	return &sTicketMessage{
		Cornerstone: *cornerstone.NewCornerstoneWithTable(dao.V2TicketMessage.Table()),
	}
}

// 获取工单信息数据
func (s *sTicketMessage) GetTicketMessageArrByTicketId(ticketId int) (m []*model.TicketMessageInfo, err error) {
	m = make([]*model.TicketMessageInfo, 0)
	err = s.Cornerstone.GetDB().
		Where(dao.V2TicketMessage.Columns().TicketId, ticketId).
		ScanList(&m, "V2TicketMessage")
	dao.V2User.Ctx(dao.V2User.DB().GetCtx()).
		Where("id", gdb.ListItemValuesUnique(m, "V2TicketMessage", "UserId")).
		ScanList(&m, "V2User", "V2TicketMessage", "id:UserId")
	return m, err
}

// 获取工单信息数据
func (s *sTicketMessage) GetTicketMessageArrByTicketIdAndUserId(ticketId, userId int) (m []*model.TicketMessageInfo, err error) {
	ticket, err := service.Ticket().GetTicketByIdAndUserId(ticketId, userId)
	if err != nil {
		return m, err
	}
	if ticket == nil {
		return m, errors.New("You are a big sb !")
	}
	m, err = s.GetTicketMessageArrByTicketId(ticketId)
	return
}

// 管理员回复工单
func (s *sTicketMessage) SaveTicketMessageAdmin(data *entity.V2TicketMessage) (err error) {
	err = s.Cornerstone.Save(data)
	if err != nil {
		return
	}

	err = service.Ticket().UpTicketStatusAndReplyStatusById(data.TicketId, 0, 1)

	return
}

// 用户回复工单
func (s *sTicketMessage) SaveTicketMessageUser(data *entity.V2TicketMessage) (err error) {
	ticket, err := service.Ticket().GetTicketByIdAndUserId(data.TicketId, data.UserId)
	if err != nil {
		return err
	}
	if ticket == nil {
		return errors.New("You are a big sb !")
	}

	err = s.Cornerstone.Save(data)
	if err != nil {
		return
	}

	err = service.Ticket().UpTicketStatusAndReplyStatusById(data.TicketId, 0, 0)

	return
}
