// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
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
		CalculateUserCommission(CType int, CRate int, fromUserId int, val float64) (commission float64, err error)
		// 更新
		UpUser(data *entity.V2User) (err error)
		// 更新过期用户的权限组和流量
		ClearExpiredUserGroupIdAndUDTransferEnable() (err error)
		// 用户注册
		RegisterUser(UserName string, Passwd string, CommissionCode string) error
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
		// 获取用户
		GetUserByTokenAndUDAndGTExpiredAt(token string) (u *entity.V2User, err error)
		// 根据 token 获取用户
		GetUserByToken(token string) (u *entity.V2User, err error)
		// 邀请码获取 邀请用户id
		GetUserByCommissionCode(commissionCode string) (u *entity.V2User, err error)
		// 获取用户并且检测用户状态
		GetUserByIdAndCheck(id int) (u *entity.V2User, err error)
		// 获取用户 订阅组下的用户数据
		GetUserListByGroupIds(groupIds []int) (u []*entity.V2User, err error)
		// 获取用户数量 订阅组下的用户数据
		GetUserCountByGroupIds(groupIds []int) (totle int, err error)
		// 更新用户 流量使用情况 直接更新数据库 u+值、d+值、t+值
		UpUserUAndDBy(data []*model.UserTraffic) (err error)
		// 更新用户 u、d、t
		UpUserDUTBy(data []*model.UserTraffic) (err error)
		// 更新用户 7天流量使用数据 用户id, 流量, 日期 20240901
		UpUserDay7Flow(ctx context.Context, userId int, flow int64, date string) (err error)
		// 从文件加载用户 7天流量使用数据到缓存
		LoadUserDay7FlowFromFile(ctx context.Context, filename string) error
		// 保存用户 7天流量使用数据 到文件
		SaveUserDay7FlowToFile(ctx context.Context, filename string) error
		// 用户登录
		Login(userName string, passwd string) (user *entity.V2User, err error)
		// 获取用户数据
		GetUserList(req *v1.UserReq, orderBy string, orderDirection string, offset int, limit int) (items []*model.UserInfo, total int, err error)
		// 修改密码
		UpUserPasswdById(ctx context.Context, req *userv1.UserUpPasswdReq) (res *userv1.UserUpPasswdRes, err error)
		// 获取当月注册量
		GetNowMonthCount() (count int, err error)
		// 重置用户的Token和uuid
		ResetTokenAndUuidById(id int) (err error)
		// 获取当月每一天注册量
		GetNowMonthDayCount() (count []int, err error)
		// 获取订阅组用户数量
		GetUserCountByPlanID(id int) (count int, err error)
		// 创建 token
		CreateToken(ctx context.Context, user *entity.V2User) (signedToken string, claims *model.JWTClaims, err error)
		// 从上下文获取用户信息
		GetCtxUser(ctx context.Context) *entity.V2User
		// 启动 把有效用户 存入到内存
		MSaveToRam() (err error)
		// 更新用户 流量使用情况2 直接更新缓存（原来有一个直接更新数据库UpUserUAndDBy）
		MUpUserUAndBy(ctx context.Context, data []*model.UserTraffic) (err error)
		// 所有数据持久化
		MSaveAllRam() (err error)
		// 更新/添加 缓存
		MUpUserMap(data *model.UserTraffic)
		// 查询数据库更新到缓存
		MGetDb2UserMap(uid int) (err error)
		// 删除 缓存
		MDelUserMap(id int)
		// 权限组获取用户
		MGetUserByGroupId(GroupId int) (d []*model.UserTraffic)
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
