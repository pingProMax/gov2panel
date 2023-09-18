// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "gov2panel/api/admin/v1"
	userv1 "gov2panel/api/user/v1"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/model/model"
)

type (
	IUser interface {
		// 账号查询用户
		GetUserByUserName(userName string) (user *entity.V2User, err error)
		// 获取 user_id 的邀请数量
		GetInviteCountByUserId(user_id int) (inviteCount int, err error)
		// 获取 用户 的佣金比例\佣金类型
		GetUserCTypeAndCRate(user *entity.V2User) (commissionType int, commissionRate int)
		// 计算用户佣金
		// 用户
		// 金额
		CalculateUserCommission(CType, CRate int, fromUserId int, val float64) (commission float64, err error)
		// 更新
		UpUser(data *entity.V2User) (err error)
		// 用户注册
		RegisterUser(UserName, Passwd, CommissionCode string) error
		// 添加用户
		AddUser(user *entity.V2User) error
		// 删除
		DelUser(ids []int) error
		// 冻结
		UpUserBanned1(ids []int) (err error)
		// AE设置
		AEUser(data *entity.V2User) (err error)
		// 获取用户
		GetUserById(id int) (u *entity.V2User, err error)
		// 邀请码获取 邀请用户id
		GetUserByCommissionCode(commissionCode string) (u *entity.V2User, err error)
		// 获取用户并且检测用户装
		GetUserByIdAndCheck(id int) (u *entity.V2User, err error)
		// 用户登录
		Login(userName, passwd string) (user *entity.V2User, err error)
		// 获取用户数据
		GetUserList(req *v1.UserReq, orderBy, orderDirection string, offset, limit int) (items []*model.UserInfo, total int, err error)
		// 修改密码
		UpUserPasswdById(req *userv1.UserUpPasswdReq) (res *userv1.UserUpPasswdRes, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
