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

//Alpha支付
type Alpha struct {
	AppId       int    `json:"app_id"`
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
type AlphaConfig struct {
	ApiUrl    *g.Var `json:"api_url"`
	AppId     *g.Var `json:"app_id"`
	AppSecret *g.Var `json:"app_secret"`
}
