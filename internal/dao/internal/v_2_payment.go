// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// V2PaymentDao is the data access object for the table v2_payment.
type V2PaymentDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns V2PaymentColumns // columns contains all the column names of Table for convenient usage.
}

// V2PaymentColumns defines and stores column names for the table v2_payment.
type V2PaymentColumns struct {
	Id                 string //
	Uuid               string // uuid
	Payment            string // 支付类型
	Name               string // 名字
	Icon               string // 图标地址
	Config             string // 配置json数
	NotifyDomain       string // 回调域名
	HandlingFeeFixed   string // 固定手续费
	HandlingFeePercent string // 百分比手续费
	Enable             string // 是否启用
	OrderId            string // 顺序
	CreatedAt          string //
	UpdatedAt          string //
	Remarks            string // 备注
}

// v2PaymentColumns holds the columns for the table v2_payment.
var v2PaymentColumns = V2PaymentColumns{
	Id:                 "id",
	Uuid:               "uuid",
	Payment:            "payment",
	Name:               "name",
	Icon:               "icon",
	Config:             "config",
	NotifyDomain:       "notify_domain",
	HandlingFeeFixed:   "handling_fee_fixed",
	HandlingFeePercent: "handling_fee_percent",
	Enable:             "enable",
	OrderId:            "order_id",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
	Remarks:            "remarks",
}

// NewV2PaymentDao creates and returns a new DAO object for table data access.
func NewV2PaymentDao() *V2PaymentDao {
	return &V2PaymentDao{
		group:   "default",
		table:   "v2_payment",
		columns: v2PaymentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *V2PaymentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *V2PaymentDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *V2PaymentDao) Columns() V2PaymentColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *V2PaymentDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *V2PaymentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *V2PaymentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
