// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// V2ServerRouteDao is the data access object for table v2_server_route.
type V2ServerRouteDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns V2ServerRouteColumns // columns contains all the column names of Table for convenient usage.
}

// V2ServerRouteColumns defines and stores column names for table v2_server_route.
type V2ServerRouteColumns struct {
	Id          string //
	Remarks     string // 备注
	Match       string // 规则
	Action      string // block|dns
	ActionValue string //
	CreatedAt   string //
	UpdatedAt   string //
	Enable      string // 是否启用
}

// v2ServerRouteColumns holds the columns for table v2_server_route.
var v2ServerRouteColumns = V2ServerRouteColumns{
	Id:          "id",
	Remarks:     "remarks",
	Match:       "match",
	Action:      "action",
	ActionValue: "action_value",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	Enable:      "enable",
}

// NewV2ServerRouteDao creates and returns a new DAO object for table data access.
func NewV2ServerRouteDao() *V2ServerRouteDao {
	return &V2ServerRouteDao{
		group:   "default",
		table:   "v2_server_route",
		columns: v2ServerRouteColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *V2ServerRouteDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *V2ServerRouteDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *V2ServerRouteDao) Columns() V2ServerRouteColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *V2ServerRouteDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *V2ServerRouteDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *V2ServerRouteDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
