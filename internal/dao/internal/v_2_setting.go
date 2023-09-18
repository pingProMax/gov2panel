// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// V2SettingDao is the data access object for table v2_setting.
type V2SettingDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns V2SettingColumns // columns contains all the column names of Table for convenient usage.
}

// V2SettingColumns defines and stores column names for table v2_setting.
type V2SettingColumns struct {
	Code      string //
	Value     string //
	OrderId   string // 顺序
	Remarks   string // 备注
	CreatedAt string //
	UpdatedAt string //
}

// v2SettingColumns holds the columns for table v2_setting.
var v2SettingColumns = V2SettingColumns{
	Code:      "code",
	Value:     "value",
	OrderId:   "order_id",
	Remarks:   "remarks",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewV2SettingDao creates and returns a new DAO object for table data access.
func NewV2SettingDao() *V2SettingDao {
	return &V2SettingDao{
		group:   "default",
		table:   "v2_setting",
		columns: v2SettingColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *V2SettingDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *V2SettingDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *V2SettingDao) Columns() V2SettingColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *V2SettingDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *V2SettingDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *V2SettingDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
