// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// V2PlanDao is the data access object for the table v2_plan.
type V2PlanDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  V2PlanColumns      // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// V2PlanColumns defines and stores column names for the table v2_plan.
type V2PlanColumns struct {
	Id                 string //
	TransferEnable     string // 流量(GB)
	SpeedLimit         string // 速度限制
	Name               string // 名称
	Show               string // 是否显示
	OrderId            string // 顺序
	Renew              string // 是否允许续购
	Content            string // 描述
	Expired            string // 有效期 day
	Price              string // 价格
	ResetTrafficMethod string // 套餐类型，1 覆盖、2 叠加
	CapacityLimit      string // 最大用户
	CreatedAt          string //
	UpdatedAt          string //
	Remarks            string // 备注
}

// v2PlanColumns holds the columns for the table v2_plan.
var v2PlanColumns = V2PlanColumns{
	Id:                 "id",
	TransferEnable:     "transfer_enable",
	SpeedLimit:         "speed_limit",
	Name:               "name",
	Show:               "show",
	OrderId:            "order_id",
	Renew:              "renew",
	Content:            "content",
	Expired:            "expired",
	Price:              "price",
	ResetTrafficMethod: "reset_traffic_method",
	CapacityLimit:      "capacity_limit",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
	Remarks:            "remarks",
}

// NewV2PlanDao creates and returns a new DAO object for table data access.
func NewV2PlanDao(handlers ...gdb.ModelHandler) *V2PlanDao {
	return &V2PlanDao{
		group:    "default",
		table:    "v2_plan",
		columns:  v2PlanColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *V2PlanDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *V2PlanDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *V2PlanDao) Columns() V2PlanColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *V2PlanDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *V2PlanDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *V2PlanDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
