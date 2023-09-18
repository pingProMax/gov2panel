package v1

import (
	"gov2panel/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type PaymentReq struct {
	g.Meta `path:"/payment" tags:"Payment" method:"get,post" summary:"支付管理"`
	entity.V2Payment
}
type PaymentRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Data   []*entity.V2Payment `json:"data"`
}

type PaymentAEReq struct {
	g.Meta `path:"/payment/ae" tags:"Payment" method:"post" summary:"支付管理AE"`
	entity.V2Payment
}
type PaymentAERes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type PaymentDelReq struct {
	g.Meta `path:"/payment/del" tags:"Payment" method:"post" summary:"支付管理删除"`
	Ids    []int `json:"ids"`
}
type PaymentDelRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type PaymentGetShowReq struct {
	g.Meta `path:"/payment/get_show" tags:"Payment" method:"post" summary:"获取显示的支付列表"`
}
type PaymentGetShowRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Data   []*entity.V2Payment `json:"data"`
}
