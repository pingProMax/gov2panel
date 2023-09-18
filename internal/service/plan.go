// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	userv1 "gov2panel/api/user/v1"
	"gov2panel/internal/model/entity"
)

type (
	IPlan interface {
		// AE设置
		AEPlan(data *entity.V2Plan) (err error)
		// 删除
		DelPlan(ids []int) error
		// 获取所有
		GetPlanAllList(req entity.V2Plan) (m []*entity.V2Plan, err error)
		// 获取显示的订阅
		GetPlanShowList() (m []*entity.V2Plan, err error)
		// 获取显示的订阅 可覆盖的
		GetPlanShowAndResetTrafficMethod1List() (m []*entity.V2Plan, err error)
		// 删除
		GetPlanById(id int) (d *entity.V2Plan, err error)
		// 用户购买套餐处理
		UserBuy(req *userv1.BuyReq) (res *userv1.BuyRes, err error)
	}
)

var (
	localPlan IPlan
)

func Plan() IPlan {
	if localPlan == nil {
		panic("implement not found for interface IPlan, forgot register?")
	}
	return localPlan
}

func RegisterPlan(i IPlan) {
	localPlan = i
}
