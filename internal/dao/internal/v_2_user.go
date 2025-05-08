// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// V2UserDao is the data access object for the table v2_user.
type V2UserDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  V2UserColumns      // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// V2UserColumns defines and stores column names for the table v2_user.
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
	CommissionType    string // 3: system 1: period 2: onetime
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
	LastLoginIp       string //
	Uuid              string // uuid
	GroupId           string // 权限组
	Token             string // token 订阅用
	Remarks           string // 备注
	ExpiredAt         string //
	CreatedAt         string //
	UpdatedAt         string //
}

// v2UserColumns holds the columns for the table v2_user.
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
func NewV2UserDao(handlers ...gdb.ModelHandler) *V2UserDao {
	return &V2UserDao{
		group:    "default",
		table:    "v2_user",
		columns:  v2UserColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *V2UserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *V2UserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *V2UserDao) Columns() V2UserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *V2UserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *V2UserDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *V2UserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
