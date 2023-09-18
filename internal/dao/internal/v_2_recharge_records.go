// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// V2RechargeRecordsDao is the data access object for table v2_recharge_records.
type V2RechargeRecordsDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns V2RechargeRecordsColumns // columns contains all the column names of Table for convenient usage.
}

// V2RechargeRecordsColumns defines and stores column names for table v2_recharge_records.
type V2RechargeRecordsColumns struct {
	Id              string //
	Amount          string // 金额
	UserId          string // 用户id
	OperateType     string // 1充值 2消费
	RechargeName    string // 充值类型 operate_type=1才有
	ConsumptionName string // 消费类型 operate_type=2才有
	Remarks         string // 备注
	TransactionId   string // 订单号 规则看程序注释
	CreatedAt       string // 创建时间
	UpdatedAt       string // 更新时间
}

// v2RechargeRecordsColumns holds the columns for table v2_recharge_records.
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
func NewV2RechargeRecordsDao() *V2RechargeRecordsDao {
	return &V2RechargeRecordsDao{
		group:   "default",
		table:   "v2_recharge_records",
		columns: v2RechargeRecordsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *V2RechargeRecordsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *V2RechargeRecordsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *V2RechargeRecordsDao) Columns() V2RechargeRecordsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *V2RechargeRecordsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *V2RechargeRecordsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *V2RechargeRecordsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
