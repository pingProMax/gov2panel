package ticket

import (
	"fmt"
	"gov2panel/internal/dao"
	"gov2panel/internal/logic/cornerstone"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/model/model"
	"gov2panel/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type sTicket struct {
	Cornerstone cornerstone.Cornerstone
}

func init() {
	service.RegisterTicket(New())
}

func New() *sTicket {
	return &sTicket{
		Cornerstone: *cornerstone.NewCornerstoneWithTable(dao.V2Ticket.Table()),
	}
}

// 获取工单数据
func (s *sTicket) GetUserList(req *entity.V2Ticket, userName string, orderBy, orderDirection string, offset, limit int) (data []*model.TicketInfo, totle int, err error) {

	//查询用户名
	data, totle, err = s.GetTicketList(req, userName, orderBy, orderDirection, offset, limit)
	return data, totle, err
}

// AE设置
func (s *sTicket) AETicket(data *entity.V2Ticket) (err error) {
	if data.Id != 0 {
		err = s.Cornerstone.UpdateById(data.Id, data)
		return
	}

	err = s.Cornerstone.Save(data)
	return
}

// 删除
func (s *sTicket) DelTicket(ids []int) error {
	return s.Cornerstone.DelByIds(ids)
}

// 关闭
func (s *sTicket) CloseTicket(ids []int) (err error) {
	_, err = s.Cornerstone.GetDB().Data(dao.V2Ticket.Columns().Status, 1).WhereIn(dao.V2Ticket.Columns().Id, ids).Update()
	return err
}

// 用户关闭工单
func (s *sTicket) CloseTicketByUserIdAndId(ids []int, userId int) (err error) {
	_, err = s.Cornerstone.GetDB().Data(dao.V2Ticket.Columns().Status, 1).WhereIn(dao.V2Ticket.Columns().Id, ids).Where(dao.V2Ticket.Columns().UserId, userId).Update()
	return err
}

// 获取工单 根据id和用户id
func (s *sTicket) GetTicketByIdAndUserId(ticketId, userId int) (data *entity.V2Ticket, err error) {
	data = new(entity.V2Ticket)
	err = s.Cornerstone.GetOneById(ticketId, data)
	return
}

// 更新工单 根据id
func (s *sTicket) UpTicketStatusAndReplyStatusById(ticketId, status, reply_status int) (err error) {
	_, err = s.Cornerstone.GetDB().
		Data(g.Map{
			dao.V2Ticket.Columns().Status:      status,
			dao.V2Ticket.Columns().ReplyStatus: reply_status,
		}).
		Where(dao.V2Ticket.Columns().Id, ticketId).Update()

	return
}

// 获取 打开工单的数量
func (s *sTicket) GetOpenTicketCount() (totle int, err error) {
	totle, err = s.Cornerstone.GetDB().Where(dao.V2Ticket.Columns().Status, -1).Count()
	return
}

// 获取所有
func (s *sTicket) GetTicketList(req *entity.V2Ticket, userName string, orderBy, orderDirection string, offset, limit int) (m []*model.TicketInfo, total int, err error) {
	m = make([]*model.TicketInfo, 0)
	db := s.Cornerstone.GetDB()
	orderBy = dao.V2Ticket.Table() + "." + orderBy
	db.LeftJoin(
		dao.V2User.Table(),
		fmt.Sprintf("%s.%s=%s.%s",
			dao.V2Ticket.Table(),
			dao.V2Ticket.Columns().UserId,
			dao.V2User.Table(),
			dao.V2User.Columns().Id,
		))

	db.WhereLike(dao.V2Ticket.Columns().Subject, "%"+req.Subject+"%")
	db.WhereLike(dao.V2User.Columns().UserName, "%"+userName+"%")
	if req.Id != 0 {
		db.Where(dao.V2Ticket.Columns().Id, req.Id)
	}
	if req.UserId != 0 {
		db.Where(dao.V2Ticket.Columns().UserId, req.UserId)
	}
	if req.Level != 0 {
		db.Where(dao.V2Ticket.Columns().Level, req.Level)
	}
	if req.Status != 0 {
		db.Where(dao.V2Ticket.Columns().Status, req.Status)
	}
	if req.ReplyStatus != 0 {
		db.Where(dao.V2Ticket.Columns().ReplyStatus, req.ReplyStatus)
	}

	dbC := *db
	dbCCount := &dbC

	db.Fields(fmt.Sprintf("%s.*", dao.V2Ticket.Table()))
	err = db.Order(orderBy, orderDirection).Limit(offset, limit).ScanList(&m, "V2Ticket")
	if err != nil {
		return m, 0, err
	}

	db.Fields("*")
	total, err = dbCCount.Count()
	if err != nil {
		return m, 0, err
	}

	if total > 0 {
		err = dao.V2User.Ctx(dao.V2User.DB().GetCtx()).
			Where("id", gdb.ListItemValuesUnique(m, "V2Ticket", "UserId")).
			ScanList(&m, "V2User", "V2Ticket", "id:UserId")
	}

	return m, total, err
}
