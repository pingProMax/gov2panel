// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// V2RechargeRecordsDao is the data access object for the table v2_recharge_records.
type V2RechargeRecordsDao struct {
	table    string                   // table is the underlying table name of the DAO.
	group    string                   // group is the database configuration group name of the current DAO.
	columns  V2RechargeRecordsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler       // handlers for customized model modification.
}

// V2RechargeRecordsColumns defines and stores column names for the table v2_recharge_records.
type V2RechargeRecordsColumns struct {
	Id              string //
	Amount          string // 金额
	UserId          string // 用户id
	OperateType     string // 1充值 2消费
	RechargeName    string // 充值类型 operate_type=1才有
	ConsumptionName string // 消费类型 operate_type=2才有
	Remarks         string // 备注
	TransactionId   string // 订单号 规则看程序注释
	CreatedAt       string //
	UpdatedAt       string //
}

// v2RechargeRecordsColumns holds the columns for the table v2_recharge_records.
var v2RechargeRecordsColumns = V2RechargeRecordsColumns{
	Id:              "id",
	Amount:          "amount",
	UserId:          "user_id",
	OperateType:     "operate_type",
	RechargeName:    "recharge_name",
	ConsumptionName: "consumption_name",
	Remarks:         "remarks",
	TransactionId:   "transaction_id",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

// NewV2RechargeRecordsDao creates and returns a new DAO object for table data access.
func NewV2RechargeRecordsDao(handlers ...gdb.ModelHandler) *V2RechargeRecordsDao {
	return &V2RechargeRecordsDao{
		group:    "default",
		table:    "v2_recharge_records",
		columns:  v2RechargeRecordsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *V2RechargeRecordsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *V2RechargeRecordsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *V2RechargeRecordsDao) Columns() V2RechargeRecordsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *V2RechargeRecordsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *V2RechargeRecordsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *V2RechargeRecordsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
