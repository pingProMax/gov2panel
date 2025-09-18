// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
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
		// 获取可覆盖的订阅
		GetPlanResetTrafficMethod1List() (m []*entity.V2Plan, err error)
		// 根据id获取
		GetPlanById(id int) (d *entity.V2Plan, err error)
		// 用户购买/续费套餐处理
		UserBuyAndRenew(code string, plan *entity.V2Plan, user *entity.V2User) (err error)
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
