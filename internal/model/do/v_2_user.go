// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// V2User is the golang structure of table v2_user for DAO operations like Where/Data.
type V2User struct {
	g.Meta            `orm:"table:v2_user, do:true"`
	Id                interface{} //
	InviteUserId      interface{} // 邀请id
	TelegramId        interface{} // 电报id
	UserName          interface{} // 账号
	Password          interface{} // 密码
	PasswordAlgo      interface{} // 加密方式
	PasswordSalt      interface{} // 加密盐
	Balance           interface{} // 账户余额
	Discount          interface{} // 专享折扣
	CommissionType    interface{} // 0: system 1: period 2: onetime
	CommissionRate    interface{} // 返利比例
	CommissionBalance interface{} // aff余额
	CommissionCode    interface{} // 邀请码
	T                 interface{} // 最后在线时间戳
	U                 interface{} // 上传
	D                 interface{} // 下载
	TransferEnable    interface{} // 流量
	Banned            interface{} // 是否禁用
	IsAdmin           interface{} // 是否管理员
	IsStaff           interface{} // 是否员工
	LastLoginAt       interface{} // 最后登入时间
	LastLoginIp       interface{} // 最后登入ip
	Uuid              interface{} // uuid
	GroupId           interface{} // 权限组
	Token             interface{} // token 订阅用
	Remarks           interface{} // 备注
	ExpiredAt         *gtime.Time // 到期时间
	CreatedAt         *gtime.Time // 创建时间
	UpdatedAt         *gtime.Time // 更新时间
}
