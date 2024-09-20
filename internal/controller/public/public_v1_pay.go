package public

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	v1 "gov2panel/api/public/v1"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/model/model"
	"gov2panel/internal/service"
	"gov2panel/internal/utils"
)

func (c *ControllerV1) EPayNotify(ctx context.Context, req *v1.EPayNotifyReq) (res *v1.EPayNotifyRes, err error) {

	// err = service.RechargeRecords().SaveRechargeRecords(
	// 	&entity.V2RechargeRecords{
	// 		Amount:        100,
	// 		UserId:        0,
	// 		OperateType:   1,
	// 		RechargeName:  "系统测试",
	// 		Remarks:       "",
	// 		TransactionId: "test",
	// 	},
	// 	"系统测试",
	// 	110,
	// 	0,
	// 	"",
	// )

	// fmt.Println(err, "test")
	// return

	r := g.RequestFromCtx(ctx)
	// rw := r.Response.RawWriter()

	paramMap := r.GetMapStrStr()
	sign := paramMap["sign"]
	signType := paramMap["sign_type"]

	// 删除 sign 和 sign_type
	delete(paramMap, "sign")
	delete(paramMap, "sign_type")

	urlData := notify(paramMap)

	paramList := strings.Split(paramMap["param"], "|") //用户实际得到的金额|支付方式的id|用户id|订单号
	if len(paramList) != 4 {
		ghttp.RequestFromCtx(ctx).Response.WriteExit("你你妈妈呢呢1？")
	}
	paymentId, _ := strconv.Atoi(paramList[1])
	payment, err := service.Payment().GetPaymentById(paymentId)
	if err != nil {
		ghttp.RequestFromCtx(ctx).Response.WriteExit("你你妈妈呢呢2？")
	}

	usetId, _ := strconv.Atoi(paramList[2])

	//{1040 2023091415451087211 1694677507984 wxpay product 2.1 TRADE_SUCCESS 1.00|6|3787 e5615bcd45e87e4ad8173ad60ec6a620 MD5}

	epayConfig := model.EpayConfig{}
	err = json.Unmarshal([]byte(payment.Config), &epayConfig)
	if err != nil {
		return
	}

	//验证签名
	if signType != "MD5" {
		ghttp.RequestFromCtx(ctx).Response.WriteExit("error")
	}
	fmt.Println(sign, utils.MD5V(urlData, epayConfig.Key.String()), paramList)
	if sign != utils.MD5V(urlData, epayConfig.Key.String()) {
		ghttp.RequestFromCtx(ctx).Response.WriteExit("你你妈妈呢呢3？")
	}

	//充值
	num, _ := strconv.ParseFloat(paramList[0], 64)         //用户实际到账金额
	payNum, _ := strconv.ParseFloat(paramMap["money"], 64) //支付金额
	err = service.RechargeRecords().SaveRechargeRecords(
		&entity.V2RechargeRecords{
			Amount:        num,
			UserId:        usetId,
			OperateType:   1,
			RechargeName:  payment.Name,
			Remarks:       "",
			TransactionId: paramList[3],
		},
		payment.Name,
		payNum,
		payment.Id,
		"",
	)
	if err != nil {
		return
	}

	ghttp.RequestFromCtx(ctx).Response.WriteExit("success")
	return
}

func notify(params map[string]string) string {

	// 获取所有键并排序
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 生成查询字符串
	var strBuilder strings.Builder
	for _, k := range keys {
		strBuilder.WriteString(url.QueryEscape(k) + "=" + url.QueryEscape(params[k]) + "&")
	}

	// 去掉最后的 "&" 并拼接密钥
	queryStr := strBuilder.String()
	if len(queryStr) > 0 {
		queryStr = queryStr[:len(queryStr)-1]
	}

	return queryStr
}
