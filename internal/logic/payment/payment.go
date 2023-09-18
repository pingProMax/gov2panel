package payment

import (
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
	"strconv"
	"time"
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
func (s *sPayment) GetPayUrl(res *v1.PayRedirectionReq) (url string, err error) {
	payment := new(entity.V2Payment)
	err = s.Cornerstone.GetOneById(res.PaymentId, payment)
	if err != nil {
		return
	}
	if res.Amount < 0 {
		err = errors.New("金额小于0，你充值你奶奶个腿")
		return
	}
	res.Amount = utils.Decimal(res.Amount)

	switch payment.Payment {
	case "epay":

		epayConfig := model.EpayConfig{}
		err = json.Unmarshal([]byte(payment.Config), &epayConfig)
		if err != nil {
			return
		}
		addr := fmt.Sprintf("%s/submit.php?", epayConfig.Url.String()) //地址

		HandlingFeeAmount := 0.00
		//计算手续费
		if payment.HandlingFeePercent > 0 { //百分比手续费
			HandlingFeeAmount = res.Amount * float64(payment.HandlingFeePercent) / 100
		}

		if payment.HandlingFeeFixed > 0 { //固定手续费
			HandlingFeeAmount = HandlingFeeAmount + payment.HandlingFeeFixed
		}

		priceStr := strconv.FormatFloat(res.Amount, 'f', 2, 64)                          //金额
		transactionId := utils.RechargeOrderNo(res.Amount+HandlingFeeAmount, payment.Id) //订单号

		url = fmt.Sprintf("money=%s&name=%s&notify_url=%s&out_trade_no=%s&param=%s&pid=%s&return_url=%s",
			strconv.FormatFloat(res.Amount+HandlingFeeAmount, 'f', 2, 64), //金额
			transactionId, //name
			payment.NotifyDomain+"/pay/e_pay_notify",                                              //服务器异步通知地址
			strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10),                  //订单号
			priceStr+"|"+strconv.Itoa(payment.Id)+"|"+strconv.Itoa(res.TUserID)+"|"+transactionId, //自定义 用户实际得到的金额|支付方式的id|用户id|订单号
			epayConfig.Pid.String(),             //pid
			payment.NotifyDomain+"/user/wallet", //页面跳转通知地址

		)

		url = addr + url + fmt.Sprintf("&sign=%s&sign_type=MD5", utils.MD5V(url, epayConfig.Key.String()))
	default:
		err = errors.New("该支付类型没有实现")
	}

	return
}
