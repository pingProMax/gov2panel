// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// V2TicketMessageDao is the data access object for table v2_ticket_message.
type V2TicketMessageDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns V2TicketMessageColumns // columns contains all the column names of Table for convenient usage.
}

// V2TicketMessageColumns defines and stores column names for table v2_ticket_message.
type V2TicketMessageColumns struct {
	Id        string //
	UserId    string //
	TicketId  string //
	Message   string //
	CreatedAt string //
	UpdatedAt string //
}

// v2TicketMessageColumns holds the columns for table v2_ticket_message.
var v2TicketMessageColumns = V2TicketMessageColumns{
	Id:        "id",
	UserId:    "user_id",
	TicketId:  "ticket_id",
	Message:   "message",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewV2TicketMessageDao creates and returns a new DAO object for table data access.
func NewV2TicketMessageDao() *V2TicketMessageDao {
	return &V2TicketMessageDao{
		group:   "default",
		table:   "v2_ticket_message",
		columns: v2TicketMessageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *V2TicketMessageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *V2TicketMessageDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *V2TicketMessageDao) Columns() V2TicketMessageColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *V2TicketMessageDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *V2TicketMessageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *V2TicketMessageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
