package payment

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	v1 "gov2panel/api/user/v1"
	"gov2panel/internal/dao"
	"gov2panel/internal/logic/cornerstone"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/model/model"
	"gov2panel/internal/service"
	"gov2panel/internal/utils"
	"net/url"
	"sort"
	"strconv"
	"strings"

	gfJson "github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tidwall/gjson"
)

type sPayment struct {
	Cornerstone cornerstone.Cornerstone
}

func init() {
	service.RegisterPayment(New())
}

func New() *sPayment {
	return &sPayment{
		Cornerstone: *cornerstone.NewCornerstoneWithTable(dao.V2Payment.Table()),
	}
}

// AE设置
func (s *sPayment) AEPayment(data *entity.V2Payment) (err error) {
	if data.Id != 0 {
		err = s.Cornerstone.UpdateById(data.Id, data)
		return
	}

	err = s.Cornerstone.Save(data)
	return
}

// 删除
func (s *sPayment) DelPayment(ids []int) error {
	return s.Cornerstone.DelByIds(ids)
}

// 获取id
func (s *sPayment) GetPaymentById(id int) (data *entity.V2Payment, err error) {
	data = new(entity.V2Payment)
	err = s.Cornerstone.GetOneById(id, data)
	return
}

// 获取所有
func (s *sPayment) AdminGetPaymentAllList(req entity.V2Payment) (m []*entity.V2Payment, err error) {
	m = make([]*entity.V2Payment, 0)
	err = s.Cornerstone.GetDB().
		OmitEmpty().
		Where(dao.V2Payment.Columns().Id, req.Id).
		WhereLike(dao.V2Payment.Columns().Name, "%"+req.Name+"%").
		OrderDesc("order_id").Scan(&m)
	return m, err

}

// 获取显示的支付
func (s *sPayment) GetPaymentShowList() (m []*entity.V2Payment, err error) {
	m = make([]*entity.V2Payment, 0)
	err = s.Cornerstone.GetDB().
		FieldsEx(
			dao.V2Payment.Columns().Config,
			dao.V2Payment.Columns().Remarks,
		).
		Where(dao.V2Payment.Columns().Enable, 1).
		OrderDesc("order_id").Scan(&m)
	return m, err
}

// 支付业务，获取支付url
func (s *sPayment) GetPayUrl(ctx context.Context, res *v1.PayRedirectionReq) (urlStr string, err error) {
	payment := new(entity.V2Payment)
	err = s.Cornerstone.GetOneById(res.PaymentId, payment)
	if err != nil {
		return
	}
	if res.Amount < 0 {
		err = errors.New("金额小于0，你充值你奶奶个腿")
		return
	}
	res.Amount = utils.Decimal(res.Amount) //只保留两位小数

	HandlingFeeAmount := 0.00 //手续费金额
	//计算手续费
	if payment.HandlingFeePercent > 0 { //百分比手续费
		HandlingFeeAmount = res.Amount * float64(payment.HandlingFeePercent) / 100
		HandlingFeeAmount = utils.Decimal(HandlingFeeAmount)
	}

	if payment.HandlingFeeFixed > 0 { //固定手续费
		HandlingFeeAmount = HandlingFeeAmount + payment.HandlingFeeFixed
		HandlingFeeAmount = utils.Decimal(HandlingFeeAmount)
	}

	transactionId := utils.RechargeOrderNo(res.Amount+HandlingFeeAmount, res.Amount, payment.Id, service.User().GetCtxUser(ctx).Id) //订单号 系统用

	payAmountStr := strconv.FormatFloat(res.Amount+HandlingFeeAmount, 'f', 2, 64) //支付金额
	// 2. 再将字符串解析回 float64
	payAmount, err := strconv.ParseFloat(payAmountStr, 64)
	if err != nil {
		// 处理错误
		return
	}

	switch payment.Payment {
	case "epay":

		epayConfig := model.EpayConfig{}
		err = json.Unmarshal([]byte(payment.Config), &epayConfig)
		if err != nil {
			return
		}
		addr := fmt.Sprintf("%s/submit.php?", epayConfig.Url.String()) //地址

		urlStr = fmt.Sprintf("money=%s&name=%s&notify_url=%s&out_trade_no=%s&pid=%s&return_url=%s",
			payAmountStr,  //金额
			transactionId, //name
			payment.NotifyDomain+"/pay/e_pay_notify", //服务器异步通知地址
			transactionId,               //订单号
			epayConfig.Pid.String(),     //pid
			res.Redirect+"/user/wallet", //页面跳转通知地址
		)

		urlStr = addr + urlStr + fmt.Sprintf("&sign=%s&sign_type=MD5", utils.MD5V(urlStr, epayConfig.Key.String()))
	case "alpha":

		alphaConfig := model.AlphaConfig{}
		err = json.Unmarshal([]byte(payment.Config), &alphaConfig)
		if err != nil {
			return
		}
		addr := fmt.Sprintf("%s/api/v1/tron", alphaConfig.ApiUrl.String()) //地址

		urlStr = fmt.Sprintf(
			"app_id=%s&notify_url=%s&out_trade_no=%s&return_url=%s&total_amount=%s",
			alphaConfig.AppId,
			url.QueryEscape(payment.NotifyDomain+"/pay/e_pay_notify"),
			transactionId,
			url.QueryEscape(payment.NotifyDomain+"/user/wallet"),
			strconv.FormatFloat((res.Amount+HandlingFeeAmount)*100, 'f', 2, 64),
		)

		urlStr = urlStr + fmt.Sprintf("&sign=%s", utils.MD5V(urlStr, alphaConfig.AppSecret.String()))
		c := g.Client()
		c.SetHeader("User-Agent", "Alpha")
		if r, err := c.Post(gctx.New(), addr, urlStr); err != nil {
			err = errors.New(err.Error())

		} else {
			defer r.Close()
			jsonStr := r.ReadAllString()
			urlStr = gjson.Get(jsonStr, "url").String()
		}
	case "bepusdt":

		bepUsdtConfig := model.BepUsdtConfig{}
		err = json.Unmarshal([]byte(payment.Config), &bepUsdtConfig)
		if err != nil {
			return
		}

		addr := fmt.Sprintf("%s/api/v1/order/create-order", bepUsdtConfig.Url.String()) //地址

		requestData := g.Map{
			"order_id":     transactionId,                                    //商户订单编号（唯一标识）
			"notify_url":   payment.NotifyDomain + "/pay/bepusdt_pay_notify", //支付结果异步回调地址
			"redirect_url": res.Redirect + "/user/wallet",                    //支付成功后商户跳转地址
			"amount":       payAmount,                                        //支付金额（法币金额）；留空或传 0 则进入地址独占模式，收到任意金额均触发回调
			"name":         transactionId,                                    //商品名称
			"timeout":      bepUsdtConfig.Timeout,                            //订单超时时间（秒），最低 120 秒
		}

		// 1. 计算签名
		signature := s.BepusdtGenerateSignature(requestData, bepUsdtConfig.Key.String())
		// 2. 将计算好的签名赋值回 map
		requestData["signature"] = signature

		c := g.Client()
		c.SetHeader("User-Agent", "gov2panel")
		if r, err := c.Post(gctx.New(), addr, gfJson.New(requestData).MustToJsonString()); err != nil {
			err = errors.New(err.Error())
		} else {
			defer r.Close()
			jsonStr := r.ReadAllString()
			fmt.Println(jsonStr, 77777)
			urlStr = gjson.Get(jsonStr, "data.payment_url").String()
		}
	default:
		err = errors.New("该支付类型没有实现")
	}

	return
}

// BepusdtGenerateSignature 生成 MD5 签名
func (s *sPayment) BepusdtGenerateSignature(data g.Map, token string) string {
	var keys []string

	// 第一步：筛选出非空且非 signature 的参数名
	for k, v := range data {
		valStr := gconv.String(v) // 使用 GoFrame 的 gconv 安全转换为字符串
		if k == "signature" || valStr == "" {
			continue
		}
		keys = append(keys, k)
	}

	// 按参数名 ASCII 码从小到大排序（字典序）
	sort.Strings(keys)

	// 按 key=value 格式拼接，使用 & 连接
	var parts []string
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s=%s", k, gconv.String(data[k])))
	}
	signStr := strings.Join(parts, "&")

	// 第二步：末尾追加 API Token（注意没有 &）
	signStr += token

	// 对完整字符串进行 MD5 加密并转为小写
	hasher := md5.New()
	hasher.Write([]byte(signStr))
	return hex.EncodeToString(hasher.Sum(nil))
}
