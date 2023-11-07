package v1

import (
	"gov2panel/internal/model/model"

	"github.com/gogf/gf/v2/frame/g"
)

type EPayNotifyReq struct {
	g.Meta `path:"/pay/e_pay_notify" tags:"Pay" method:"get,post" summary:"易支付 异步通知"`
	model.Epay
}

type EPayNotifyRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type AlphaNotifyReq struct {
	g.Meta `path:"/pay/alpha_notify" tags:"Pay" method:"get,post" summary:"Alpha 异步通知"`
	model.Epay
}

type AlphaNotifyRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
