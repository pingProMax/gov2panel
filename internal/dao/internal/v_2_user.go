// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// V2UserDao is the data access object for table v2_user.
type V2UserDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns V2UserColumns // columns contains all the column names of Table for convenient usage.
}

// V2UserColumns defines and stores column names for table v2_user.
type V2UserColumns struct {
	Id                string //
	InviteUserId      string // 邀请id
	TelegramId        string // 电报id
	UserName          string // 账号
	Password          string // 密码
	PasswordAlgo      string // 加密方式
	PasswordSalt      string // 加密盐
	Balance           string // 账户余额
	Discount          string // 专享折扣
	CommissionType    string // 0: system 1: period 2: onetime
	CommissionRate    string // 返利比例
	CommissionBalance string // aff余额
	CommissionCode    string // 邀请码
	T                 string // 最后在线时间戳
	U                 string // 上传
	D                 string // 下载
	TransferEnable    string // 流量
	Banned            string // 是否禁用
	IsAdmin           string // 是否管理员
	IsStaff           string // 是否员工
	LastLoginAt       string // 最后登入时间
	LastLoginIp       string // 最后登入ip
	Uuid              string // uuid
	GroupId           string // 权限组
	Token             string // token 订阅用
	Remarks           string // 备注
	ExpiredAt         string // 到期时间
	CreatedAt         string // 创建时间
	UpdatedAt         string // 更新时间
}

// v2UserColumns holds the columns for table v2_user.
var v2UserColumns = V2UserColumns{
	Id:                "id",
	InviteUserId:      "invite_user_id",
	TelegramId:        "telegram_id",
	UserName:          "user_name",
	Password:          "password",
	PasswordAlgo:      "password_algo",
	PasswordSalt:      "password_salt",
	Balance:           "balance",
	Discount:          "discount",
	CommissionType:    "commission_type",
	CommissionRate:    "commission_rate",
	CommissionBalance: "commission_balance",
	CommissionCode:    "commission_code",
	T:                 "t",
	U:                 "u",
	D:                 "d",
	TransferEnable:    "transfer_enable",
	Banned:            "banned",
	IsAdmin:           "is_admin",
	IsStaff:           "is_staff",
	LastLoginAt:       "last_login_at",
	LastLoginIp:       "last_login_ip",
	Uuid:              "uuid",
	GroupId:           "group_id",
	Token:             "token",
	Remarks:           "remarks",
	ExpiredAt:         "expired_at",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
}

// NewV2UserDao creates and returns a new DAO object for table data access.
func NewV2UserDao() *V2UserDao {
	return &V2UserDao{
		group:   "default",
		table:   "v2_user",
		columns: v2UserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *V2UserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *V2UserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *V2UserDao) Columns() V2UserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *V2UserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *V2UserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *V2UserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
