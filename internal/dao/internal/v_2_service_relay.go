// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// V2ServiceRelayDao is the data access object for the table v2_service_relay.
type V2ServiceRelayDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  V2ServiceRelayColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// V2ServiceRelayColumns defines and stores column names for the table v2_service_relay.
type V2ServiceRelayColumns struct {
	Id        string //
	Ip        string // ip
	NameGroup string // 组名字
	Asn       string // asn分组 AS9808|AS4134|AS4837
	Show      string // 是否启用
	OrderId   string // 顺序
	Remarks   string // 备注
	CreatedAt string //
	UpdatedAt string //
}

// v2ServiceRelayColumns holds the columns for the table v2_service_relay.
var v2ServiceRelayColumns = V2ServiceRelayColumns{
	Id:        "id",
	Ip:        "ip",
	NameGroup: "name_group",
	Asn:       "asn",
	Show:      "show",
	OrderId:   "order_id",
	Remarks:   "remarks",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewV2ServiceRelayDao creates and returns a new DAO object for table data access.
func NewV2ServiceRelayDao(handlers ...gdb.ModelHandler) *V2ServiceRelayDao {
	return &V2ServiceRelayDao{
		group:    "default",
		table:    "v2_service_relay",
		columns:  v2ServiceRelayColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *V2ServiceRelayDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *V2ServiceRelayDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *V2ServiceRelayDao) Columns() V2ServiceRelayColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *V2ServiceRelayDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *V2ServiceRelayDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *V2ServiceRelayDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
