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
	IRechargeRecords interface {
		// 保存数据
		// payCode 充值代码，填的充值通道，人工填admin
		// val 充值金额|消费金额
		// id 充值的支付id|消费订阅id
		// couponCode 消费的优惠码，填的优惠码
		SaveRechargeRecords(data *entity.V2RechargeRecords, payCode string, val float64, id int, couponCode string) (err error)
		// 获取数据
		GetRechargeRecordsList(req *v1.RechargeRecordsReq, orderBy, orderDirection string, offset, limit int) (m []*model.RechargeRecordsInfo, total int, err error)
		// 获取数据根据用户id
		GetRechargeRecordsListByUserId(userId int, orderBy, orderDirection string, offset, limit int) (m []*entity.V2RechargeRecords, total int, err error)
		// 更新备注
		UpRechargeRecordsRemarksById(id int, remarks string) (err error)
		// 获取当月收入
		GetNowMonthSumAmount() (amount float64, err error)
		// 获取当月每一天的收入
		GetNowMonthDaySum() (data []int, err error)
	}
)

var (
	localRechargeRecords IRechargeRecords
)

func RechargeRecords() IRechargeRecords {
	if localRechargeRecords == nil {
		panic("implement not found for interface IRechargeRecords, forgot register?")
	}
	return localRechargeRecords
}

func RegisterRechargeRecords(i IRechargeRecords) {
	localRechargeRecords = i
}
