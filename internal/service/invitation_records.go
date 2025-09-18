// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "gov2panel/api/admin/v1"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/model/model"
)

type (
	IInvitationRecords interface {
		// 获取数据
		GetInvitationRecordsList(req *v1.InvitationRecordsReq, orderBy string, orderDirection string, offset int, limit int) (m []*model.InvitationRecordsInfo, total int, err error)
		// 获取数据根据用户id
		GetInvitationRecordsListByUserId(userId int, orderBy string, orderDirection string, offset int, limit int) (m []*model.InvitationRecordsInfo, total int, err error)
		// 获取数据根据id
		GetOneById(id int) (d *entity.V2InvitationRecords, err error)
		// 获取数据根据id
		GetOneByFromUserId(from_user_id int) (d *entity.V2InvitationRecords, err error)
		// 添加
		Insert(data *entity.V2InvitationRecords) (err error)
		// 更新
		UpInvitationRecordsStateById(id int, state int) (err error)
		// 审核状态
		AdminiUpStateById(id int, state int) (err error)
		// CommissionTransferBalance佣金转余额
		CommissionTransferBalance(userId int) (err error)
		// 佣金提现
		WithdrawalBalance(userId int) (err error)
	}
)

var (
	localInvitationRecords IInvitationRecords
)

func InvitationRecords() IInvitationRecords {
	if localInvitationRecords == nil {
		panic("implement not found for interface IInvitationRecords, forgot register?")
	}
	return localInvitationRecords
}

func RegisterInvitationRecords(i IInvitationRecords) {
	localInvitationRecords = i
}
