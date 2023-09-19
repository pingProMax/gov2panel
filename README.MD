# GoFrame Template For SingleRepo

Quick Start: 
- https://goframe.org/pages/viewpage.action?pageId=1114399


git commit -m "up"
git push -u origin master

http://127.0.0.1:8080/pay/e_pay_notify?pid=1040&trade_no=2023091415451087211&out_trade_no=1694677507984&type=wxpay&name=product&money=12.11&trade_status=TRADE_SUCCESS&param=10.00|6|3787|1694680294-12.11-6-95751&sign=b710567e1e26bc1d52547ae0cee51b1d&sign_type=MD5


/api/v1/server/UniProxy/config?node_id=1&node_type=v2ray&token=123456
{
	"server_port": 46619,
	"network": "tcp",
	"networkSettings": null,
	"tls": 1,
	"base_config": {
		"push_interval": 60,
		"pull_interval": 60
	},
	"routes": [{
		"id": 1,
		"match": ["*.360.cn"],
		"action": "dns",
		"action_value": "4.4.4.4"
	}, {
		"id": 2,
		"match": ["*.360.cn"],
		"action": "block",
		"action_value": null
	}]
}

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go