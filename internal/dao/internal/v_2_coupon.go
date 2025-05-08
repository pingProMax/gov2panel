// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// V2CouponDao is the data access object for the table v2_coupon.
type V2CouponDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  V2CouponColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// V2CouponColumns defines and stores column names for the table v2_coupon.
type V2CouponColumns struct {
	Id               string //
	Code             string // 优惠码
	Name             string // 名称
	Type             string // 类型 1金额优惠 2百分比优惠
	Value            string // 优惠多少
	Enable           string // 是否启用
	LimitUse         string // 每个用户可使用次数
	LimitUseWithUser string // 最大使用次数
	LimitPlanId      string // 指定订阅
	StartedAt        string //
	EndedAt          string //
	CreatedAt        string //
	UpdatedAt        string //
	Remarks          string // 备注
}

// v2CouponColumns holds the columns for the table v2_coupon.
var v2CouponColumns = V2CouponColumns{
	Id:               "id",
	Code:             "code",
	Name:             "name",
	Type:             "type",
	Value:            "value",
	Enable:           "enable",
	LimitUse:         "limit_use",
	LimitUseWithUser: "limit_use_with_user",
	LimitPlanId:      "limit_plan_id",
	StartedAt:        "started_at",
	EndedAt:          "ended_at",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	Remarks:          "remarks",
}

// NewV2CouponDao creates and returns a new DAO object for table data access.
func NewV2CouponDao(handlers ...gdb.ModelHandler) *V2CouponDao {
	return &V2CouponDao{
		group:    "default",
		table:    "v2_coupon",
		columns:  v2CouponColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *V2CouponDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *V2CouponDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *V2CouponDao) Columns() V2CouponColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *V2CouponDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *V2CouponDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *V2CouponDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
