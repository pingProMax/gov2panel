package public

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"

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

	paramList := strings.Split(paramMap["out_trade_no"], "-") //时间戳-充值金额(实际支付的)-用户得到金额-payID-用户ID
	if len(paramList) != 5 {
		ghttp.RequestFromCtx(ctx).Response.WriteExit("你你妈妈呢呢1？")
	}
	paymentId := gconv.Int(paramList[3])
	payment, err := service.Payment().GetPaymentById(paymentId)
	if err != nil {
		ghttp.RequestFromCtx(ctx).Response.WriteExit("你你妈妈呢呢2？")
	}

	usetId := gconv.Int(paramList[4])

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
		ghttp.RequestFromCtx(ctx).Response.WriteExit("你你妈妈呢呢3？" + urlData)
	}

	//充值
	num := gconv.Float64(paramList[2])                 //用户实际到账金额
	payNum := gconv.Float64(paramMap["actual_amount"]) //支付金额
	err = service.RechargeRecords().SaveRechargeRecords(
		&entity.V2RechargeRecords{
			Amount:        num,
			UserId:        usetId,
			OperateType:   1,
			RechargeName:  payment.Name,
			Remarks:       "",
			TransactionId: gconv.String(paramMap["out_trade_no"]),
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
		if params[k] != "" {
			strBuilder.WriteString(k + "=" + params[k] + "&")
		}
	}

	// 去掉最后的 "&" 并拼接密钥
	queryStr := strBuilder.String()
	if len(queryStr) > 0 {
		queryStr = queryStr[:len(queryStr)-1]
	}

	queryStr, err := url.QueryUnescape(queryStr)
	if err != nil {
		return err.Error()
	}
	return queryStr
}

func (c *ControllerV1) BEpusdtNotify(ctx context.Context, req *v1.BEpusdtNotifyReq) (res *v1.BEpusdtNotifyRes, err error) {

	r := g.RequestFromCtx(ctx)
	// rw := r.Response.RawWriter()
	paramMap := r.GetMap()

	if gconv.Int(paramMap["status"]) != 2 {
		ghttp.RequestFromCtx(ctx).Response.WriteExit("支付状态不对！")
	}

	signature := paramMap["signature"]

	// 删除 signature
	delete(paramMap, "signature")

	paramList := strings.Split(gconv.String(paramMap["order_id"]), "-") //时间戳-充值金额(实际支付的)-用户得到金额-payID-用户ID
	if len(paramList) != 5 {
		ghttp.RequestFromCtx(ctx).Response.WriteExit("你你妈妈呢呢1？")
	}

	paymentId := gconv.Int(paramList[3])
	payment, err := service.Payment().GetPaymentById(paymentId)
	if err != nil {
		ghttp.RequestFromCtx(ctx).Response.WriteExit("你你妈妈呢呢2？")
	}
	usetId := gconv.Int(paramList[4])

	bepUsdtConfig := model.BepUsdtConfig{}
	err = json.Unmarshal([]byte(payment.Config), &bepUsdtConfig)
	if err != nil {
		return
	}

	//签名
	bepusdtSigStr := service.Payment().BepusdtGenerateSignature(paramMap, bepUsdtConfig.Key.String())

	//验证签名
	fmt.Println(signature, bepusdtSigStr, gconv.String(paramMap))
	if signature != bepusdtSigStr {
		ghttp.RequestFromCtx(ctx).Response.WriteExit("你你妈妈呢呢3？")
	}

	//充值
	num := gconv.Float64(paramList[2])                 //用户实际到账金额
	payNum := gconv.Float64(paramMap["actual_amount"]) //支付金额
	err = service.RechargeRecords().SaveRechargeRecords(
		&entity.V2RechargeRecords{
			Amount:        num,
			UserId:        usetId,
			OperateType:   1,
			RechargeName:  payment.Name,
			Remarks:       "",
			TransactionId: gconv.String(paramMap["order_id"]),
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
