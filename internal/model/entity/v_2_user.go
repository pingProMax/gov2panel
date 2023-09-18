// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// V2User is the golang structure for table v2_user.
type V2User struct {
	Id                int         `json:"id"                 ` //
	InviteUserId      int         `json:"invite_user_id"     ` // 邀请id
	TelegramId        int64       `json:"telegram_id"        ` // 电报id
	UserName          string      `json:"user_name"          ` // 账号
	Password          string      `json:"password"           ` // 密码
	PasswordAlgo      string      `json:"password_algo"      ` // 加密方式
	PasswordSalt      string      `json:"password_salt"      ` // 加密盐
	Balance           float64     `json:"balance"            ` // 账户余额
	Discount          float64     `json:"discount"           ` // 专享折扣
	CommissionType    int         `json:"commission_type"    ` // 0: system 1: period 2: onetime
	CommissionRate    int         `json:"commission_rate"    ` // 返利比例
	CommissionBalance float64     `json:"commission_balance" ` // aff余额
	CommissionCode    string      `json:"commission_code"    ` // 邀请码
	T                 int         `json:"t"                  ` // 最后在线时间戳
	U                 int64       `json:"u"                  ` // 上传
	D                 int64       `json:"d"                  ` // 下载
	TransferEnable    int64       `json:"transfer_enable"    ` // 流量
	Banned            int         `json:"banned"             ` // 是否禁用
	IsAdmin           int         `json:"is_admin"           ` // 是否管理员
	IsStaff           int         `json:"is_staff"           ` // 是否员工
	LastLoginAt       int         `json:"last_login_at"      ` // 最后登入时间
	LastLoginIp       int         `json:"last_login_ip"      ` // 最后登入ip
	Uuid              string      `json:"uuid"               ` // uuid
	GroupId           int         `json:"group_id"           ` // 权限组
	Token             string      `json:"token"              ` // token 订阅用
	Remarks           string      `json:"remarks"            ` // 备注
	ExpiredAt         *gtime.Time `json:"expired_at"         ` // 到期时间
	CreatedAt         *gtime.Time `json:"created_at"         ` // 创建时间
	UpdatedAt         *gtime.Time `json:"updated_at"         ` // 更新时间
}
