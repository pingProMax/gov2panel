package proxy_service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/service"
	"gov2panel/internal/utils"
	"net/url"
	"strings"

	"github.com/gogf/gf/v2/util/gconv"
)

// 专门处理订阅的
// GetV2rayUrl
func (s *sProxyService) GetV2rayUrl(
	ctx context.Context,
	v2Service *entity.V2ProxyService,
	user *entity.V2User,
	v2ServiceRelayList []*entity.V2ServiceRelay,
	asn string, //用户asn
) (v2rayUrl string, err error) {

	if strings.Contains(v2Service.Host, ",") {
		v2Service.Host = utils.GetRandomString(v2Service.Host)
	}

	// 定义中继服务器 $relay[]
	prefix := "$relay["
	suffix := "]"
	// 检查并截取
	if strings.HasPrefix(v2Service.Host, prefix) && strings.HasSuffix(v2Service.Host, suffix) {
		// 移除前缀和后缀，剩下就是中间的值
		val := v2Service.Host[len(prefix) : len(v2Service.Host)-len(suffix)]
		v2Service.Host = service.ServerRelay().GetRandomRelayByFilter(v2ServiceRelayList, val, asn)
	}

	v2Service.Host = strings.ReplaceAll(v2Service.Host, "$uuid", user.Uuid)

	serviceJson := make(map[string]interface{})
	err = json.Unmarshal([]byte(v2Service.ServiceJson), &serviceJson)
	if err != nil {
		return "", err
	}

	switch strings.Split(v2Service.Agreement, "/")[1] {
	case "vmess":
		s := map[string]string{
			"v":    "2",
			"add":  v2Service.Host, //链接地址
			"ps":   v2Service.Name, //名字
			"port": v2Service.Port, //端口
			"id":   user.Uuid,      //uuid
			"aid":  "0",
			"net":  gconv.String(serviceJson["net"]),
			"type": gconv.String(serviceJson["type"]),
			"tls":  gconv.String(serviceJson["tls"]),
			"sni":  gconv.String(serviceJson["sni"]),
			"alpn": gconv.String(serviceJson["alpn"]),
			"host": gconv.String(serviceJson["host"]),
			"path": gconv.String(serviceJson["path"]),
			"scy":  gconv.String(serviceJson["scy"]),
			"fp":   utils.GetRandomString(gconv.String(serviceJson["fp"])),
		}
		ds, err := json.Marshal(s)
		if err != nil {
			return "", err
		}

		v2rayUrl = fmt.Sprintf("%s://%s", "vmess", base64.StdEncoding.EncodeToString(ds))
		return v2rayUrl, nil

	case "vless":
		// vless://uuid@127.0.0.1:8888?encryption=none&security=reality&sni=sni.com&fp=qq&pbk=PublicKey&sid=ShortId&spx=SpiderX&type=tcp&headerType=http&host=host.com#vless
		// vless://uuid@127.0.0.1:8888?encryption=none&security=tls&sni=sni.com&alpn=http%2F1.1&fp=qq&pbk=PublicKey&sid=ShortId&spx=SpiderX&type=tcp&headerType=http&host=host.com#vless
		// vless://78f10ea1-81a4-4bf5-876f-90e3001f37dc@127.0.0.1:8888?encryption=none&flow=xtls-rprx-vision&security=tls&sni=sni.com&alpn=http%2F1.1&fp=qq&pbk=PublicKey&sid=ShortId&spx=SpiderX&type=tcp&headerType=http&host=host.com#vless

		// 概述 https://github.com/XTLS/Xray-core/discussions/716
		// (tcp http 无法导入path？？？)
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
		resultThis := fmt.Sprintf(
			"%s://%s@%s:%s?type=%s&encryption=%s&security=%s&path=%s&host=%s&headerType=%s&seed=%s&serviceName=%s&mode=%s&authority=%s&extra=%s&fp=%s&sni=%s&alpn=%s&flow=%s&pbk=%s&sid=%s&pqv=%s&spx=%s#%s",
			strings.Split(v2Service.Agreement, "/")[1],
			url.QueryEscape(user.Uuid),
			v2Service.Host,
			v2Service.Port,
			gconv.String(serviceJson["type"]),
			gconv.String(serviceJson["encryption"]),
			gconv.String(serviceJson["security"]),
			url.QueryEscape(gconv.String(serviceJson["path"])),
			url.QueryEscape(gconv.String(serviceJson["host"])),
			gconv.String(serviceJson["headerType"]),
			url.QueryEscape(gconv.String(serviceJson["seed"])),
			url.QueryEscape(gconv.String(serviceJson["serviceName"])),
			utils.GetRandomString(gconv.String(serviceJson["mode"])),
			url.QueryEscape(gconv.String(serviceJson["authority"])),
			url.QueryEscape(gconv.String(serviceJson["extra"])),
			utils.GetRandomString(gconv.String(serviceJson["fp"])),
			gconv.String(serviceJson["sni"]),
			url.QueryEscape(gconv.String(serviceJson["alpn"])),
			gconv.String(serviceJson["flow"]),
			gconv.String(serviceJson["pbk"]),
			utils.GetRandomString(gconv.String(serviceJson["sid"])),
			gconv.String(serviceJson["pqv"]),
			url.QueryEscape(utils.GetRandomString(gconv.String(serviceJson["spx"]))),
			v2Service.Name,
		)

		u, _ := url.Parse(resultThis)
		// 获取查询参数
		query := u.Query()

		// 删除值为空字符串的参数
		for key, values := range query {
			filtered := values[:0]
			for _, v := range values {
				if v != "" {
					filtered = append(filtered, v)
				}
			}
			if len(filtered) == 0 {
				query.Del(key)
			} else {
				query[key] = filtered
			}
		}

		// 重新设置 URL 的查询参数
		u.RawQuery = query.Encode()

		v2rayUrl = u.String()

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
		v2rayUrl = fmt.Sprintf(
			"%s://%s@%s:%s#%s",
			"ss",
			str,
			v2Service.Host,
			v2Service.Port,
			v2Service.Name,
		)
		return

	case "trojan":
		//trojan://密码@地址:端口?security=tls&sni=sni.com&alpn=http%2F1.1&fp=chrome&type=tcp&headerType=none&host=host.com#名字

		v2rayUrl = fmt.Sprintf(
			"%s://%s@%s:%s?security=%s&sni=%s&alpn=%s&fp=%s&type=%s&headerType=%s&host=%s#%s",
			"trojan",
			user.Uuid,
			v2Service.Host,
			v2Service.Port,
			gconv.String(serviceJson["security"]),
			gconv.String(serviceJson["sni"]),
			gconv.String(serviceJson["alpn"]),
			utils.GetRandomString(gconv.String(serviceJson["fp"])),
			gconv.String(serviceJson["type"]),
			gconv.String(serviceJson["headerType"]),
			gconv.String(serviceJson["host"]),
			v2Service.Name,
		)

		return

	}

	return
}
