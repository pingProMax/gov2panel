package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	v1 "gov2panel/api/admin/v1"
	"gov2panel/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Index(ctx context.Context, req *v1.IndexReq) (res *v1.IndexRes, err error) {
	switch g.RequestFromCtx(ctx).Method {
	case "GET":
		res = &v1.IndexRes{}
		chartLabels := make([]string, 0)
		nowTime := time.Now()

		monthStr := ""
		if int(nowTime.Month()) < 10 {
			monthStr = "0" + strconv.Itoa(int(nowTime.Month()))
		} else {
			monthStr = strconv.Itoa(int(nowTime.Month()))
		}

		for i := nowTime.Day(); i > 0; i-- {

			iStr := ""
			if i < 10 {
				iStr = fmt.Sprintf("0%s", strconv.Itoa(i))
			} else {
				iStr = strconv.Itoa(i)
			}
			chartLabels = append([]string{fmt.Sprintf("%s/%s", monthStr, iStr)}, chartLabels...)
		}

		var chartLabelsBytes []byte
		chartLabelsBytes, err = json.Marshal(chartLabels)
		if err != nil {
			return
		}
		res.ChartLabels = string(chartLabelsBytes)

		//获取当月每天的用户注册量
		var dayUserCountList []int
		dayUserCountList, err = service.User().GetNowMonthDayCount()
		if err != nil {
			return
		}
		var dayUserCountBytes []byte
		dayUserCountBytes, err = json.Marshal(dayUserCountList)
		if err != nil {
			return
		}
		res.DayUserCount = string(dayUserCountBytes)

		//获取当个月每天的收益
		var dayIncomeSumList []int
		dayIncomeSumList, err = service.RechargeRecords().GetNowMonthDaySum()
		if err != nil {
			return
		}

		fmt.Println(dayIncomeSumList)
		var dayIncomeSumBytes []byte
		dayIncomeSumBytes, err = json.Marshal(dayIncomeSumList)
		if err != nil {
			return
		}
		res.DayIncomeSum = string(dayIncomeSumBytes)

		//本月收入
		res.MonthAmount, err = service.RechargeRecords().GetNowMonthSumAmount()
		if err != nil {
			return
		}

		//未处理工单数量
		res.OpenTicketCount, err = service.Ticket().GetOpenTicketCount()
		if err != nil {
			return
		}

		//节点数量
		res.ServiceCount, err = service.ProxyService().GetServiceCount()
		if err != nil {
			return
		}

		//本月注册量
		res.MonthUserCount, err = service.User().GetNowMonthCount()
		if err != nil {
			return
		}

		setTplAdmin(ctx, "index", g.Map{"data": res})
	default:
		return
	}

	return
}
