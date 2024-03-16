// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// V2TicketDao is the data access object for table v2_ticket.
type V2TicketDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns V2TicketColumns // columns contains all the column names of Table for convenient usage.
}

// V2TicketColumns defines and stores column names for table v2_ticket.
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

// v2TicketColumns holds the columns for table v2_ticket.
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
func NewV2TicketDao() *V2TicketDao {
	return &V2TicketDao{
		group:   "default",
		table:   "v2_ticket",
		columns: v2TicketColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *V2TicketDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *V2TicketDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *V2TicketDao) Columns() V2TicketColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *V2TicketDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *V2TicketDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *V2TicketDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
