// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// V2InvitationRecordsDao is the data access object for the table v2_invitation_records.
type V2InvitationRecordsDao struct {
	table    string                     // table is the underlying table name of the DAO.
	group    string                     // group is the database configuration group name of the current DAO.
	columns  V2InvitationRecordsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler         // handlers for customized model modification.
}

// V2InvitationRecordsColumns defines and stores column names for the table v2_invitation_records.
type V2InvitationRecordsColumns struct {
	Id                string //
	Amount            string // 金额
	UserId            string // 邀请者
	FromUserId        string // 被邀请者
	CommissionRate    string // 佣金比例
	RechargeRecordsId string // 订单id
	CreatedAt         string //
	UpdatedAt         string //
	OperateType       string // 1邀请 2提现
	State             string // 状态 -1未审核 1审核 2拒绝
}

// v2InvitationRecordsColumns holds the columns for the table v2_invitation_records.
var v2InvitationRecordsColumns = V2InvitationRecordsColumns{
	Id:                "id",
	Amount:            "amount",
	UserId:            "user_id",
	FromUserId:        "from_user_id",
	CommissionRate:    "commission_rate",
	RechargeRecordsId: "recharge_records_id",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
	OperateType:       "operate_type",
	State:             "state",
}

// NewV2InvitationRecordsDao creates and returns a new DAO object for table data access.
func NewV2InvitationRecordsDao(handlers ...gdb.ModelHandler) *V2InvitationRecordsDao {
	return &V2InvitationRecordsDao{
		group:    "default",
		table:    "v2_invitation_records",
		columns:  v2InvitationRecordsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *V2InvitationRecordsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *V2InvitationRecordsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *V2InvitationRecordsDao) Columns() V2InvitationRecordsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *V2InvitationRecordsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *V2InvitationRecordsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *V2InvitationRecordsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
