package model

import "github.com/gogf/gf/v2/frame/g"

//易支付
type Epay struct {
	Pid         int    `json:"pid"`
	TradeNo     string `json:"trade_no"`
	OutTradeNo  string `json:"out_trade_no"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Money       string `json:"money"`
	TradeStatus string `json:"trade_status"`
	Param       string `json:"param"`
	Sign        string `json:"sign"`
	SignType    string `json:"sign_type"`
}

//易支付 配置文件
type EpayConfig struct {
	Url *g.Var
	Pid *g.Var
	Key *g.Var
}
