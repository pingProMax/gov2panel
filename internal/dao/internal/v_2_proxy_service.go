// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// V2ProxyServiceDao is the data access object for table v2_proxy_service.
type V2ProxyServiceDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns V2ProxyServiceColumns // columns contains all the column names of Table for convenient usage.
}

// V2ProxyServiceColumns defines and stores column names for table v2_proxy_service.
type V2ProxyServiceColumns struct {
	Id          string //
	Agreement   string // 协议
	ServiceJson string // 服务器json数据
	Name        string // 显示名称
	PlanId      string // 所属订阅组
	Show        string // 是否显示
	Host        string // 服务器地址
	Port        string // 服务器端口
	Rate        string // 倍率
	OrderId     string // 顺序
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	RouteId     string // 所属路由组
}

// v2ProxyServiceColumns holds the columns for table v2_proxy_service.
var v2ProxyServiceColumns = V2ProxyServiceColumns{
	Id:          "id",
	Agreement:   "agreement",
	ServiceJson: "service_json",
	Name:        "name",
	PlanId:      "plan_id",
	Show:        "show",
	Host:        "host",
	Port:        "port",
	Rate:        "rate",
	OrderId:     "order_id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	RouteId:     "route_id",
}

// NewV2ProxyServiceDao creates and returns a new DAO object for table data access.
func NewV2ProxyServiceDao() *V2ProxyServiceDao {
	return &V2ProxyServiceDao{
		group:   "default",
		table:   "v2_proxy_service",
		columns: v2ProxyServiceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *V2ProxyServiceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *V2ProxyServiceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *V2ProxyServiceDao) Columns() V2ProxyServiceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *V2ProxyServiceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *V2ProxyServiceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *V2ProxyServiceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
