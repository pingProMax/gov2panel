package public

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"

	v1 "gov2panel/api/public/v1"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/model/model"
	"gov2panel/internal/service"
	"gov2panel/internal/utils"
)

func (c *ControllerV1) EPayNotify(ctx context.Context, req *v1.EPayNotifyReq) (res *v1.EPayNotifyRes, err error) {
	paramList := strings.Split(req.Epay.Param, "|") //用户实际得到的金额|支付方式的id|用户id|订单号
	if len(paramList) != 4 {
		ghttp.RequestFromCtx(ctx).Response.WriteExit("你你妈妈呢呢1？")
	}
	paymentId, _ := strconv.Atoi(paramList[1])
	payment, err := service.Payment().GetPaymentById(paymentId)
	if err != nil {
		ghttp.RequestFromCtx(ctx).Response.WriteExit("你你妈妈呢呢2？")
	}

	usetId, _ := strconv.Atoi(paramList[2])

	urlData := fmt.Sprintf("money=%s&name=%s&out_trade_no=%s&param=%s&pid=%s&trade_no=%s&trade_status=%s&type=%s",
		req.Epay.Money,             //金额
		req.Epay.Name,              //name
		req.Epay.OutTradeNo,        //商户系统内部的订单号
		req.Epay.Param,             //自定义
		strconv.Itoa(req.Epay.Pid), //pid
		req.Epay.TradeNo,           //易支付订单号
		req.Epay.TradeStatus,       //支付状态
		req.Epay.Type,              //支付方式
	)

	//{1040 2023091415451087211 1694677507984 wxpay product 2.1 TRADE_SUCCESS 1.00|6|3787 e5615bcd45e87e4ad8173ad60ec6a620 MD5}

	epayConfig := model.EpayConfig{}
	err = json.Unmarshal([]byte(payment.Config), &epayConfig)
	if err != nil {
		return
	}

	//验证签名
	if req.SignType != "MD5" {
		ghttp.RequestFromCtx(ctx).Response.WriteExit("error")
	}
	fmt.Println(req.Epay.Sign, utils.MD5V(urlData, epayConfig.Key.String()), paramList)
	if req.Epay.Sign != utils.MD5V(urlData, epayConfig.Key.String()) {
		ghttp.RequestFromCtx(ctx).Response.WriteExit("你你妈妈呢呢3？")
	}

	//充值
	num, _ := strconv.ParseFloat(paramList[0], 64)      //用户实际到账金额
	payNum, _ := strconv.ParseFloat(req.Epay.Money, 64) //支付金额
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

func (c *ControllerV1) AlphaNotify(ctx context.Context, req *v1.AlphaNotifyReq) (res *v1.AlphaNotifyRes, err error) {

	ghttp.RequestFromCtx(ctx).Response.WriteExit("success")
	return
}
