// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// V2InvitationRecordsDao is the data access object for table v2_invitation_records.
type V2InvitationRecordsDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns V2InvitationRecordsColumns // columns contains all the column names of Table for convenient usage.
}

// V2InvitationRecordsColumns defines and stores column names for table v2_invitation_records.
type V2InvitationRecordsColumns struct {
	Id                string //
	Amount            string // 金额
	UserId            string // 邀请者
	FromUserId        string // 被邀请者
	CommissionRate    string // 佣金比例
	RechargeRecordsId string // 订单id
	CreatedAt         string // 创建时间
	UpdatedAt         string // 更新时间
	OperateType       string // 1邀请 2提现
	State             string // 状态 0未审核 1审核 2拒绝
}

// v2InvitationRecordsColumns holds the columns for table v2_invitation_records.
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
func NewV2InvitationRecordsDao() *V2InvitationRecordsDao {
	return &V2InvitationRecordsDao{
		group:   "default",
		table:   "v2_invitation_records",
		columns: v2InvitationRecordsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *V2InvitationRecordsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *V2InvitationRecordsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *V2InvitationRecordsDao) Columns() V2InvitationRecordsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *V2InvitationRecordsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *V2InvitationRecordsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *V2InvitationRecordsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
