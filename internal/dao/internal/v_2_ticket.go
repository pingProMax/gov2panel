// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// V2TicketDao is the data access object for the table v2_ticket.
type V2TicketDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  V2TicketColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// V2TicketColumns defines and stores column names for the table v2_ticket.
type V2TicketColumns struct {
	Id          string //
	UserId      string //
	Subject     string //
	Level       string // 1低 2中 3高
	Status      string // -1:已开启 1:已关闭
	ReplyStatus string // -1:待回复 1:已回复
	CreatedAt   string //
	UpdatedAt   string //
}

// v2TicketColumns holds the columns for the table v2_ticket.
var v2TicketColumns = V2TicketColumns{
	Id:          "id",
	UserId:      "user_id",
	Subject:     "subject",
	Level:       "level",
	Status:      "status",
	ReplyStatus: "reply_status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewV2TicketDao creates and returns a new DAO object for table data access.
func NewV2TicketDao(handlers ...gdb.ModelHandler) *V2TicketDao {
	return &V2TicketDao{
		group:    "default",
		table:    "v2_ticket",
		columns:  v2TicketColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *V2TicketDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *V2TicketDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *V2TicketDao) Columns() V2TicketColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *V2TicketDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *V2TicketDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *V2TicketDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
