// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// V2User is the golang structure for table v2_user.
type V2User struct {
	Id                int         `json:"id"                 orm:"id"                 ` //
	InviteUserId      int         `json:"invite_user_id"     orm:"invite_user_id"     ` // 邀请id
	TelegramId        int64       `json:"telegram_id"        orm:"telegram_id"        ` // 电报id
	UserName          string      `json:"user_name"          orm:"user_name"          ` // 账号
	Password          string      `json:"password"           orm:"password"           ` // 密码
	PasswordAlgo      string      `json:"password_algo"      orm:"password_algo"      ` // 加密方式
	PasswordSalt      string      `json:"password_salt"      orm:"password_salt"      ` // 加密盐
	Balance           float64     `json:"balance"            orm:"balance"            ` // 账户余额
	Discount          float64     `json:"discount"           orm:"discount"           ` // 专享折扣
	CommissionType    int         `json:"commission_type"    orm:"commission_type"    ` // 3: system 1: period 2: onetime
	CommissionRate    int         `json:"commission_rate"    orm:"commission_rate"    ` // 返利比例
	CommissionBalance float64     `json:"commission_balance" orm:"commission_balance" ` // aff余额
	CommissionCode    string      `json:"commission_code"    orm:"commission_code"    ` // 邀请码
	T                 int64       `json:"t"                  orm:"t"                  ` // 最后在线时间戳
	U                 int64       `json:"u"                  orm:"u"                  ` // 上传
	D                 int64       `json:"d"                  orm:"d"                  ` // 下载
	TransferEnable    int64       `json:"transfer_enable"    orm:"transfer_enable"    ` // 流量
	Banned            int         `json:"banned"             orm:"banned"             ` // 是否禁用
	IsAdmin           int         `json:"is_admin"           orm:"is_admin"           ` // 是否管理员
	IsStaff           int         `json:"is_staff"           orm:"is_staff"           ` // 是否员工
	LastLoginAt       int         `json:"last_login_at"      orm:"last_login_at"      ` // 最后登入时间
	LastLoginIp       string      `json:"last_login_ip"      orm:"last_login_ip"      ` //
	Uuid              string      `json:"uuid"               orm:"uuid"               ` // uuid
	GroupId           int         `json:"group_id"           orm:"group_id"           ` // 权限组
	Token             string      `json:"token"              orm:"token"              ` // token 订阅用
	Remarks           string      `json:"remarks"            orm:"remarks"            ` // 备注
	ExpiredAt         *gtime.Time `json:"expired_at"         orm:"expired_at"         ` //
	CreatedAt         *gtime.Time `json:"created_at"         orm:"created_at"         ` //
	UpdatedAt         *gtime.Time `json:"updated_at"         orm:"updated_at"         ` //
}
