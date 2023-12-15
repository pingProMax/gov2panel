// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "gov2panel/api/user/v1"
	"gov2panel/internal/model/entity"
)

type (
	IPayment interface {
		// AE设置
		AEPayment(data *entity.V2Payment) (err error)
		// 删除
		DelPayment(ids []int) error
		// 获取id
		GetPaymentById(id int) (data *entity.V2Payment, err error)
		// 获取所有
		AdminGetPaymentAllList(req entity.V2Payment) (m []*entity.V2Payment, err error)
		// 获取显示的支付
		GetPaymentShowList() (m []*entity.V2Payment, err error)
		// 支付业务，获取支付url
		GetPayUrl(res *v1.PayRedirectionReq) (urlStr string, err error)
	}
)

var (
	localPayment IPayment
)

func Payment() IPayment {
	if localPayment == nil {
		panic("implement not found for interface IPayment, forgot register?")
	}
	return localPayment
}

func RegisterPayment(i IPayment) {
	localPayment = i
}
