// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// V2CouponUseDao is the data access object for the table v2_coupon_use.
type V2CouponUseDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns V2CouponUseColumns // columns contains all the column names of Table for convenient usage.
}

// V2CouponUseColumns defines and stores column names for the table v2_coupon_use.
type V2CouponUseColumns struct {
	Id        string //
	CouponId  string //
	UserId    string //
	CreatedAt string //
	UpdatedAt string //
	PlanId    string //
}

// v2CouponUseColumns holds the columns for the table v2_coupon_use.
var v2CouponUseColumns = V2CouponUseColumns{
	Id:        "id",
	CouponId:  "coupon_id",
	UserId:    "user_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	PlanId:    "plan_id",
}

// NewV2CouponUseDao creates and returns a new DAO object for table data access.
func NewV2CouponUseDao() *V2CouponUseDao {
	return &V2CouponUseDao{
		group:   "default",
		table:   "v2_coupon_use",
		columns: v2CouponUseColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *V2CouponUseDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *V2CouponUseDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *V2CouponUseDao) Columns() V2CouponUseColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *V2CouponUseDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *V2CouponUseDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *V2CouponUseDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
