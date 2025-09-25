package public

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	v1 "gov2panel/api/public/v1"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/service"
	"gov2panel/internal/utils"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) Subscribe(ctx context.Context, req *v1.SubscribeReq) (res *v1.SubscribeRes, err error) {
	res = &v1.SubscribeRes{}
	user, err := service.User().GetUserByToken(req.Token)
	if err != nil {
		return
	}

	serviceArr, err := service.ProxyService().GetServiceListByPlanIdAndShow1(user.GroupId)
	if err != nil {
		return
	}

	result := ""

	//app
	if req.FlagInfoHide {
		result = base64Sub(serviceArr, user)
		ghttp.RequestFromCtx(ctx).Response.WriteExit(base64.StdEncoding.EncodeToString([]byte(result)))

		return nil, nil
	}

	switch req.Flag {

	case "1", "2", "6": //v2rayn //v2rayng //nekobox
		result = V2rayNGSub(serviceArr, user)
	case "3": //shadowrocket
		result = ShadowrocketSub(serviceArr, user)
	case "4": //clash
		result = ClashSub(serviceArr, user)

		//https://www.clashverge.dev/guide/url_schemes.html#content-disposition
		//如果响应头中存在 profile-update-interval 字段，则配置文件的 更新间隔 将被设置为对应的值（单位: 小时）。
		ghttp.RequestFromCtx(ctx).Response.Header().Add("profile-update-interval", "24")
		//如果响应头中存在 subscription-userinfo 字段，则其对应的流量信息(单位: 字节)、到期信息(时间戳)会显示在订阅卡片上。
		ghttp.RequestFromCtx(ctx).Response.Header().Add("subscription-userinfo", fmt.Sprintf("upload=%d; download=%d; total=%d; expire=%d", user.U, user.D, user.TransferEnable, user.ExpiredAt.Unix()))
		//如果响应头中存在 profile-web-page-url 字段，则右键订阅卡片将会显示 首页 按钮。
		// ghttp.RequestFromCtx(ctx).Response.Header().Add("profile-web-page-url", "")

		ghttp.RequestFromCtx(ctx).Response.WriteExit([]byte(result))
	case "5": //shadowsocks
		result = ShadowsocksSub(serviceArr, user)
	default:
		result = V2rayNGSub(serviceArr, user)
	}

	ghttp.RequestFromCtx(ctx).Response.WriteExit(base64.StdEncoding.EncodeToString([]byte(result)))

	return nil, nil
}

// 订阅处理
func base64Sub(serviceArr []*entity.V2ProxyService, user *entity.V2User) (result string) {
	for _, service := range serviceArr {
		service.Host = GetRandIp(service.Host)
		service.Host = strings.ReplaceAll(service.Host, "$uuid$", user.Uuid)

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
				"fp":   gconv.String(serviceJson["fp"]),
			}
			ds, err := json.Marshal(s)
			if err != nil {
				return err.Error()
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
				gconv.String(serviceJson["flow"]),
				gconv.String(serviceJson["security"]),
				gconv.String(serviceJson["sni"]),
				gconv.String(serviceJson["alpn"]),
				gconv.String(serviceJson["fp"]),
				gconv.String(serviceJson["pbk"]),
				gconv.String(serviceJson["sid"]),
				gconv.String(strings.ReplaceAll(gconv.String(serviceJson["spx"]), "$uuid$", user.Uuid)),
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

			// 概述 https://github.com/XTLS/Xray-core/discussions/716
			// 为什么不用新方法，因为(tcp http 无法导入path？？？)
			// 同样 github.com/xtls/libxray 无法将vmess xhttp url解析成json。
			// 这个bug似乎很久i了, 不理解 故意的？
			/*
				protocol://
					$(uuid)
					@
					remote-host
					:
					remote-port
				?
					<protocol-specific fields>
					<transport-specific fields>
					<tls-specific fields>
				#$(descriptive-text)
			*/
			// resultThis := fmt.Sprintf(
			// 	"%s://%s@%s:%s?type=%s&encryption=%s&security=%s&path=%s&host=%s&headerType=%s&seed=%s&serviceName=%s&mode=%s&authority=%s&extra=%s&fp=%s&sni=%s&alpn=%s&flow=%s&pbk=%s&sid=%s&pqv=%s&spx=%s#%s",
			// 	strings.Split(service.Agreement, "/")[1],
			// 	url.QueryEscape(user.Uuid),
			// 	service.Host,
			// 	service.Port,
			// 	gconv.String(serviceJson["type"]),
			// 	gconv.String(serviceJson["encryption"]),
			// 	gconv.String(serviceJson["security"]),
			// 	url.QueryEscape(gconv.String(serviceJson["path"])),
			// 	url.QueryEscape(gconv.String(serviceJson["host"])),
			// 	gconv.String(serviceJson["headerType"]),
			// 	url.QueryEscape(gconv.String(serviceJson["seed"])),
			// 	url.QueryEscape(gconv.String(serviceJson["serviceName"])),
			// 	url.QueryEscape(gconv.String(serviceJson["mode"])),
			// 	url.QueryEscape(gconv.String(serviceJson["authority"])),
			// 	url.QueryEscape(gconv.String(serviceJson["extra"])),
			// 	gconv.String(serviceJson["fp"]),
			// 	gconv.String(serviceJson["sni"]),
			// 	url.QueryEscape(gconv.String(serviceJson["alpn"])),
			// 	gconv.String(serviceJson["flow"]),
			// 	gconv.String(serviceJson["pbk"]),
			// 	gconv.String(serviceJson["sid"]),
			// 	gconv.String(serviceJson["pqv"]),
			// 	url.QueryEscape(gconv.String(serviceJson["spx"])),
			// 	service.Name,
			// )

			// u, _ := url.Parse(resultThis)

			// // 获取查询参数
			// query := u.Query()

			// // 删除值为空字符串的参数
			// for key, values := range query {
			// 	filtered := values[:0]
			// 	for _, v := range values {
			// 		if v != "" {
			// 			filtered = append(filtered, v)
			// 		}
			// 	}
			// 	if len(filtered) == 0 {
			// 		query.Del(key)
			// 	} else {
			// 		query[key] = filtered
			// 	}
			// }

			// // 重新设置 URL 的查询参数
			// u.RawQuery = query.Encode()

			// result = result + u.String() + "\n"

			//对应的配置文档
			/*
				{
				  //https://github.com/XTLS/Xray-core/discussions/716
				  //订阅用数据
				  "type": "", //传输方式 tcp、kcp、ws、http、grpc、httpupgrade、xhttp 其中之一
				  "encryption": "", //加密方式， VMess时 默认为 auto；VLESS时 默认为 none；
				  "security": "",  //底层传输安全，当前可选值有 none、tls、reality；默认为 none，
				  "path": "", //路径
				  "host": "", //Host
				  "headerType": "", //(mKCP) 的伪装头部类型 当前可选值有 none / srtp / utp / wechat-video / dtls / wireguard 默认值为 none ，
				  "seed": "", //(mKCP) 种子，
				  "serviceName": "", //(gRPC) 的 ServiceName 不可为空字符串
				  "mode": "",  // (gRPC) mode gun、multi、guna ；xhtp mode
				  "authority": "", // (gRPC) authority
				  "extra": "", // (XHTTP) extra
				  "fp":"", // LS Client Hello 指纹 若使用 REALITY，此项不可省略
				  "sni": "", // TLS SNI
				  "alpn": "",
				  "flow": "", //XTLS 的流控方式。可选值为 xtls-rprx-vision 等 若使用 若使用 XTLS，此项不可省略，否则无此项。此项不可为空字符串
				  "pbk": "", //REALITY 的密码，对应配置文件中的 password 项目
				  "sid": "", //REALITY 的 ID，对应配置文件中的 shortId 项目
				  "pqv": "", //REALITY 的 ML-DSA-65 公钥，对应配置文件中的 mldsa65Verify 项目
				  "spx": "", //REALITY 的爬虫，对应配置文件中的 spiderX 项目

				  //后端对接用数据
				  "xrayConifg":{

				  }
				}

				{
				  "type": "tcp",
				  "encryption": "auto",
				  "security": "none",
				  "path": "",
				  "host": "",
				  "headerType": "none",
				  "seed": "",
				  "serviceName": "",
				  "mode": "",
				  "authority": "",
				  "extra": "",
				  "fp":"",
				  "sni": "",
				  "alpn": "",
				  "flow": "",
				  "pbk": "",
				  "sid": "",
				  "pqv": "",
				  "spx": "",

				  "xrayConifg":{

				  }
				}
			*/

		case "ss2022":
			ssPasswd := user.Uuid
			if gconv.String(serviceJson["cypher_method"]) == "2022-blake3-aes-128-gcm" {
				ssPasswd = gconv.String(serviceJson["server_key"]) + ":" + base64.StdEncoding.EncodeToString(
					gconv.Bytes(user.Uuid[0:16]),
				)
			}

			if gconv.String(serviceJson["cypher_method"]) == "2022-blake3-aes-256-gcm" {
				ssPasswd = gconv.String(serviceJson["server_key"]) + ":" + base64.StdEncoding.EncodeToString(
					gconv.Bytes(user.Uuid[0:32]),
				)
			}

			str := base64.StdEncoding.EncodeToString(
				gconv.Bytes(gconv.String(serviceJson["cypher_method"]) + ":" + ssPasswd),
			)
			str = strings.ReplaceAll(str, "+", "-")
			str = strings.ReplaceAll(str, "/", "_")
			str = strings.ReplaceAll(str, "=", "")

			// ss://base64(加密方式:密码)@地址:端口#别名
			// ss://OjY4ZDJjNTFmLTUzMTEtNDc2MS1hYTNhLTllNDg1MmYzMGYyNQ==@127.0.0.1:9996#ss2022
			result = result + fmt.Sprintf(
				"%s://%s@%s:%s#%s\n",
				"ss",
				str,
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

	return
}

// ShadowrocketSub订阅
func ShadowrocketSub(serviceArr []*entity.V2ProxyService, user *entity.V2User) (result string) {
	result = result + fmt.Sprintf("STATUS=↑:%.2fGB,↓:%.2fGB,TOT:%.2fGBExpires:%s\n", utils.BytesToGB(user.U), utils.BytesToGB(user.D), utils.BytesToGB(user.TransferEnable), user.ExpiredAt)
	isExpired := user.ExpiredAt.Before(gtime.New(time.Now()))
	if isExpired {
		return
	}
	if user.TransferEnable-user.U-user.D <= 0 {
		return
	}

	result = result + base64Sub(serviceArr, user)

	return
}

// v2rayNG订阅
func V2rayNGSub(serviceArr []*entity.V2ProxyService, user *entity.V2User) (result string) {

	//-----剩余流量
	ps2 := fmt.Sprintf("剩余流量：%.2f GB", utils.BytesToGB(user.TransferEnable-user.U-user.D))
	if user.TransferEnable-user.U-user.D <= 0 {
		ps2 = "流量已用完！！！"
	}
	s2 := map[string]string{
		"v":    "2",
		"add":  "127.0.0.1", //链接地址
		"ps":   ps2,         //名字
		"port": "443",       //端口
		"id":   user.Uuid,   //uuid
		"net":  "tcp",
	}
	ds2, err := json.Marshal(s2)
	if err != nil {
		return err.Error()
	}
	result = result + fmt.Sprintf("%s://%s\n", "vmess", base64.StdEncoding.EncodeToString(ds2))
	if user.TransferEnable-user.U-user.D <= 0 {
		return
	}

	//-----到期时间
	expiredAtStr := user.ExpiredAt.Format("Y-m-d H:i")
	isExpired := user.ExpiredAt.Before(gtime.New(time.Now()))
	ps1 := fmt.Sprintf("套餐到期：%s", expiredAtStr)
	if isExpired {
		ps1 = fmt.Sprintf("套餐已到期！到期时间：%s", expiredAtStr)
	}
	s1 := map[string]string{
		"v":    "2",
		"add":  "127.0.0.1", //链接地址
		"ps":   ps1,         //名字
		"port": "443",       //端口
		"id":   user.Uuid,   //uuid
		"net":  "tcp",
	}
	ds1, err := json.Marshal(s1)
	if err != nil {
		return err.Error()
	}
	result += fmt.Sprintf("%s://%s\n", "vmess", base64.StdEncoding.EncodeToString(ds1))
	if isExpired {
		return
	}

	result = result + base64Sub(serviceArr, user)

	return
}

// Shadowsocks订阅
func ShadowsocksSub(serviceArr []*entity.V2ProxyService, user *entity.V2User) (result string) {
	result = result + fmt.Sprintf(
		"%s://%s@%s:%s#%s\n",
		"ss",
		base64.StdEncoding.EncodeToString(
			gconv.Bytes("aes-128-gcm:"+user.Uuid),
		),
		"127.0.0.1",
		"80",
		"剩余流量："+fmt.Sprintf("%.2f GB", utils.BytesToGB(user.TransferEnable-user.U-user.D)),
	)
	result = result + fmt.Sprintf(
		"%s://%s@%s:%s#%s\n",
		"ss",
		base64.StdEncoding.EncodeToString(
			gconv.Bytes("aes-128-gcm:"+user.Uuid),
		),
		"127.0.0.1",
		"80",
		"套餐到期："+user.ExpiredAt.Format("Y-m-d H:i"),
	)

	isExpired := user.ExpiredAt.Before(gtime.New(time.Now()))
	if isExpired {
		return
	}
	if user.TransferEnable-user.U-user.D <= 0 {
		return
	}

	for _, service := range serviceArr {
		service.Host = GetRandIp(service.Host)
		service.Host = strings.ReplaceAll(service.Host, "$uuid$", user.Uuid)

		serviceJson := make(map[string]interface{})
		json.Unmarshal([]byte(service.ServiceJson), &serviceJson)
		switch strings.Split(service.Agreement, "/")[1] {
		case "ss2022":
			ssPasswd := user.Uuid
			if gconv.String(serviceJson["cypher_method"]) == "2022-blake3-aes-128-gcm" {
				ssPasswd = gconv.String(serviceJson["server_key"]) + ":" + base64.StdEncoding.EncodeToString(
					gconv.Bytes(user.Uuid[0:16]),
				)
			}

			if gconv.String(serviceJson["cypher_method"]) == "2022-blake3-aes-256-gcm" {
				ssPasswd = gconv.String(serviceJson["server_key"]) + ":" + base64.StdEncoding.EncodeToString(
					gconv.Bytes(user.Uuid[0:32]),
				)
			}

			str := base64.StdEncoding.EncodeToString(
				gconv.Bytes(gconv.String(serviceJson["cypher_method"]) + ":" + ssPasswd),
			)
			str = strings.ReplaceAll(str, "+", "-")
			str = strings.ReplaceAll(str, "/", "_")
			str = strings.ReplaceAll(str, "=", "")

			// ss://base64(加密方式:密码)@地址:端口#别名
			// ss://OjY4ZDJjNTFmLTUzMTEtNDc2MS1hYTNhLTllNDg1MmYzMGYyNQ==@127.0.0.1:9996#ss2022
			result = result + fmt.Sprintf(
				"%s://%s@%s:%s#%s\n",
				"ss",
				str,
				service.Host,
				service.Port,
				service.Name,
			)
		}
	}

	return
}

// Clash订阅
func ClashSub(serviceArr []*entity.V2ProxyService, user *entity.V2User) (result string) {

	isExpired := user.ExpiredAt.Before(gtime.New(time.Now()))
	if isExpired {
		return
	}
	if user.TransferEnable-user.U-user.D <= 0 {
		return
	}

	content, err := os.ReadFile("./manifest/config/clash.yaml")
	if err != nil {
		result = "读取文件错误./manifest/config/clash.yaml，错误：" + err.Error()
		return
	}

	result = string(content)

	nodeNameArr := make([]string, 0)

	nodeInfoArr := make([]string, 0)

	for _, service := range serviceArr {
		service.Host = GetRandIp(service.Host)
		service.Host = strings.ReplaceAll(service.Host, "$uuid$", user.Uuid)
		serviceJson := make(map[string]interface{})
		json.Unmarshal([]byte(service.ServiceJson), &serviceJson)

		nodeNameArr = append(nodeNameArr, "      - "+service.Name)

		d := map[string]interface{}{
			"name":   service.Name,
			"port":   service.Port,
			"server": service.Host,
		}

		switch strings.Split(service.Agreement, "/")[1] {
		case "vmess":
			d["type"] = "vmess"
			d["uuid"] = user.Uuid
			d["alterId"] = 0
			d["cipher"] = gconv.String(serviceJson["scy"]) //加密方式
			if gconv.String(serviceJson["tls"]) == "tls" { //tls
				d["tls"] = true
				d["skip-cert-verify"] = false
			}
			d["servername"] = gconv.String(serviceJson["sni"])
			d["network"] = gconv.String(serviceJson["net"]) //传输协议
			switch gconv.String(serviceJson["net"]) {       //传输协议
			case "ws":
				d["ws-path"] = ""
				d["ws-headers"] = map[string]interface{}{
					"Host": gconv.String(serviceJson["host"]),
				}
			case "h2":
				d["h2-opts"] = map[string]interface{}{
					"host": []string{gconv.String(serviceJson["host"])},
					"path": gconv.String(serviceJson["path"]),
				}
			case "grpc":
				d["grpc-opts"] = map[string]interface{}{
					"grpc-service-name": gconv.String(serviceJson["path"]),
				}

			}

			if gconv.String(serviceJson["type"]) == "http" { //伪装类型
				d["network"] = "http"
				d["http-opts"] = map[string]interface{}{
					"method":  "GET",
					"path":    []string{gconv.String(serviceJson["path"])},
					"headers": map[string]interface{}{"Host": []string{gconv.String(serviceJson["host"])}},
				}
			}

		case "vless":
			d["type"] = "vless"
			d["uuid"] = user.Uuid
			d["alterId"] = 0
			d["cipher"] = gconv.String(serviceJson["scy"]) //加密方式
			if gconv.String(serviceJson["tls"]) == "tls" { //tls
				d["tls"] = true
				d["skip-cert-verify"] = false
			}
			d["servername"] = gconv.String(serviceJson["sni"])
			d["network"] = gconv.String(serviceJson["net"]) //传输协议
			switch gconv.String(serviceJson["net"]) {       //传输协议
			case "ws":
				d["ws-path"] = ""
				d["ws-headers"] = map[string]interface{}{
					"Host": gconv.String(serviceJson["host"]),
				}
			case "h2":
				d["h2-opts"] = map[string]interface{}{
					"host": []string{gconv.String(serviceJson["host"])},
					"path": gconv.String(serviceJson["path"]),
				}
			case "grpc":
				d["grpc-opts"] = map[string]interface{}{
					"grpc-service-name": gconv.String(serviceJson["path"]),
				}

			}

			if gconv.String(serviceJson["type"]) == "http" { //伪装类型
				d["network"] = "http"
				d["http-opts"] = map[string]interface{}{
					"method":  "GET",
					"path":    []string{gconv.String(serviceJson["path"])},
					"headers": map[string]interface{}{"Host": []string{gconv.String(serviceJson["host"])}},
				}
			}
		case "ss2022":
			ssPasswd := user.Uuid
			if gconv.String(serviceJson["cypher_method"]) == "2022-blake3-aes-128-gcm" {
				ssPasswd = gconv.String(serviceJson["server_key"]) + ":" + base64.StdEncoding.EncodeToString(
					gconv.Bytes(user.Uuid[0:16]),
				)
			}

			if gconv.String(serviceJson["cypher_method"]) == "2022-blake3-aes-256-gcm" {
				ssPasswd = gconv.String(serviceJson["server_key"]) + ":" + base64.StdEncoding.EncodeToString(
					gconv.Bytes(user.Uuid[0:32]),
				)
			}
			d["type"] = "ss"
			d["cipher"] = gconv.String(serviceJson["cypher_method"]) //加密方式
			d["password"] = ssPasswd                                 //密码

		case "trojan":
			d["type"] = "trojan"
			d["password"] = user.Uuid
			d["sni"] = gconv.String(serviceJson["sni"])
			d["alpn"] = []string{gconv.String(serviceJson["alpn"])}
			d["skip-cert-verify"] = false
		}

		j, _ := gjson.EncodeString(d)
		fmt.Println(string(j))
		nodeInfoArr = append(nodeInfoArr, "  - "+j)
	}

	result = strings.ReplaceAll(result, "{{node_info}}", strings.Join(nodeInfoArr, "\n"))
	result = strings.ReplaceAll(result, "{{node_name}}", strings.Join(nodeNameArr, "\n"))

	return
}

// 获取随机IP 127.0.0.1,127.0.0.2,127.0.0.3 这样的格式随机获取一个
func GetRandIp(ipStr string) string {
	if strings.Contains(ipStr, ",") {
		ips := strings.Split(ipStr, ",")
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		return ips[r.Intn(len(ips))]
	}
	return ipStr

}
