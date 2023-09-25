package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type IndexReq struct {
	g.Meta `path:"/" tags:"Index" method:"get,post" summary:"首页页面和api"`
}
type IndexRes struct {
	g.Meta          `mime:"text/html" example:"string"`
	MonthAmount     float64 `json:"month_amount"`      //当月的收入
	OpenTicketCount int     `json:"open_ticket_count"` //打开工单数量
	ServiceCount    int     `json:"service_count"`     //节点数量
	MonthUserCount  int     `json:"month_user_count"`  //当月用户注册量

	ChartLabels  string `json:"chart_labels"`   //当月的 月/日
	DayUserCount string `json:"day_user_count"` //当月每天的用户注册量
	DayIncomeSum string `json:"day_income_sum"` //当月每天的收益
}
