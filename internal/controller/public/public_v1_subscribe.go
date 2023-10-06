package public

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	v1 "gov2panel/api/public/v1"
	"gov2panel/internal/service"
	"gov2panel/internal/utils"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) Subscribe(ctx context.Context, req *v1.SubscribeReq) (res *v1.SubscribeRes, err error) {
	res = &v1.SubscribeRes{}
	user, err := service.User().GetUserByTokenAndUDAndGTExpiredAt(req.Token)
	if err != nil {
		return
	}

	serviceArr, err := service.ProxyService().GetServiceListByPlanIdAndShow1(user.GroupId)
	if err != nil {
		return
	}

	result := ""

	switch req.Flag {
	case "v2rayn", "v2rayng", "shadowrocket":
		if req.Flag == "shadowrocket" {
			result = result + fmt.Sprintf("STATUS=↑:%.2fGB,↓:%.2fGB,TOT:%.2fGBExpires:%s\n", utils.BytesToGB(user.U), utils.BytesToGB(user.D), utils.BytesToGB(user.TransferEnable), user.ExpiredAt)
		}

		// base64编码   单个：协议://base64编码

		for _, service := range serviceArr {
			serviceJson := make(map[string]interface{})
			json.Unmarshal([]byte(service.ServiceJson), &serviceJson)
			switch strings.Split(service.Agreement, "/")[1] {
			case "vmess":
				s := map[string]string{
					"v":    "2",
					"add":  service.Host, //链接地址
					"ps":   service.Name, //名字
					"port": service.Port, //端口
					"id":   user.Uuid,    //uuid
					"aid":  "0",
					"net":  gconv.String(serviceJson["net"]),
					"type": gconv.String(serviceJson["type"]),
					"tls":  gconv.String(serviceJson["tls"]),
					"sni":  gconv.String(serviceJson["sni"]),
					"alpn": gconv.String(serviceJson["alpn"]),
					"host": gconv.String(serviceJson["host"]),
					"path": gconv.String(serviceJson["path"]),
					"scy":  gconv.String(serviceJson["scy"]),
				}
				ds, err := json.Marshal(s)
				if err != nil {
					return res, err

				}

				result = result + fmt.Sprintf("%s://%s\n", "vmess", base64.StdEncoding.EncodeToString(ds))
			case "vless":
				// vless://uuid@127.0.0.1:8888?encryption=none&security=reality&sni=sni.com&fp=qq&pbk=PublicKey&sid=ShortId&spx=SpiderX&type=tcp&headerType=http&host=host.com#vless
				// vless://uuid@127.0.0.1:8888?encryption=none&security=tls&sni=sni.com&alpn=http%2F1.1&fp=qq&pbk=PublicKey&sid=ShortId&spx=SpiderX&type=tcp&headerType=http&host=host.com#vless
				// vless://78f10ea1-81a4-4bf5-876f-90e3001f37dc@127.0.0.1:8888?encryption=none&flow=xtls-rprx-vision&security=tls&sni=sni.com&alpn=http%2F1.1&fp=qq&pbk=PublicKey&sid=ShortId&spx=SpiderX&type=tcp&headerType=http&host=host.com#vless

				result = result + fmt.Sprintf(
					"%s://%s@%s:%s?encryption=%s&flow=%s&security=%s&sni=%s&alpn=%s&fp=%s&pbk=%s&sid=%s&spx=%s&type=%s&serviceName=%s&mode=%s&headerType=%s&quicSecurity=%s&key=%s&host=%s&path=%s&seed=%s#%s\n",
					"vless",
					user.Uuid,
					service.Host,
					service.Port,
					gconv.String(serviceJson["encryption"]),
					gconv.String(serviceJson["security"]),
					gconv.String(serviceJson["flow"]),
					gconv.String(serviceJson["sni"]),
					gconv.String(serviceJson["alpn"]),
					gconv.String(serviceJson["fp"]),
					gconv.String(serviceJson["pbk"]),
					gconv.String(serviceJson["sid"]),
					gconv.String(serviceJson["spx"]),
					gconv.String(serviceJson["type"]),
					gconv.String(serviceJson["serviceName"]),
					gconv.String(serviceJson["mode"]),
					gconv.String(serviceJson["headerType"]),
					gconv.String(serviceJson["quicSecurity"]),
					gconv.String(serviceJson["key"]),
					gconv.String(serviceJson["host"]),
					gconv.String(serviceJson["path"]),
					gconv.String(serviceJson["seed"]),
					service.Name,
				)

			case "ss2022":
				// ss://base64(加密方式:密码)@地址:端口#别名
				// ss://OjY4ZDJjNTFmLTUzMTEtNDc2MS1hYTNhLTllNDg1MmYzMGYyNQ==@127.0.0.1:9996#ss2022
				result = result + fmt.Sprintf(
					"%s://%s@%s:%s#%s\n",
					"ss",
					base64.StdEncoding.EncodeToString(
						gconv.Bytes(gconv.String(serviceJson["encryption"])+":"+user.Uuid),
					),
					service.Host,
					service.Port,
					service.Name,
				)

			case "trojan":
				//trojan://密码@地址:端口?security=tls&sni=sni.com&alpn=http%2F1.1&fp=chrome&type=tcp&headerType=none&host=host.com#名字

				result = result + fmt.Sprintf(
					"%s://%s@%s:%s?security=%s&sni=%s&alpn=%s&fp=%s&type=%s&headerType=%s&host=%s#%s\n",
					"trojan",
					user.Uuid,
					service.Host,
					service.Port,
					gconv.String(serviceJson["security"]),
					gconv.String(serviceJson["sni"]),
					gconv.String(serviceJson["alpn"]),
					gconv.String(serviceJson["fp"]),
					gconv.String(serviceJson["type"]),
					gconv.String(serviceJson["headerType"]),
					gconv.String(serviceJson["host"]),
					service.Name,
				)

			}

		}

	case "shadowsocks":
		for _, service := range serviceArr {
			serviceJson := make(map[string]interface{})
			json.Unmarshal([]byte(service.ServiceJson), &serviceJson)
			switch strings.Split(service.Agreement, "/")[1] {
			case "ss2022":
				// ss://base64(加密方式:密码)@地址:端口#别名
				// ss://OjY4ZDJjNTFmLTUzMTEtNDc2MS1hYTNhLTllNDg1MmYzMGYyNQ==@127.0.0.1:9996#ss2022
				result = result + fmt.Sprintf(
					"%s://%s@%s:%s#%s\n",
					"ss",
					base64.StdEncoding.EncodeToString(
						gconv.Bytes(gconv.String(serviceJson["encryption"])+":"+user.Uuid),
					),
					service.Host,
					service.Port,
					service.Name,
				)
			}

		}

	}

	ghttp.RequestFromCtx(ctx).Response.WriteExit(base64.StdEncoding.EncodeToString([]byte(result)))

	return nil, nil
}
