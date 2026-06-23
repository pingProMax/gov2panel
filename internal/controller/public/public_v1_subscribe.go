package public

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	v1 "gov2panel/api/public/v1"
	"gov2panel/internal/logic/service_relay"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/service"
	"gov2panel/internal/utils"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
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

	clientIp := ghttp.RequestFromCtx(ctx).GetClientIp()
	asn := service_relay.GetASN(clientIp)
	fmt.Println(gconv.String(user.Id)+"@gov2panel.subscribe", clientIp, asn)

	userAgent := ghttp.RequestFromCtx(ctx).UserAgent()
	userAgentList := strings.Split(g.RequestFromCtx(ctx).GetCtxVar("setting").MapStrStr()["subscribe_user_agent"], "|")
	if !utils.ContainsAny(userAgent, userAgentList) {
		err = errors.New("订阅方式已经停用，请提交工单联系管理员！")
		fmt.Println(gconv.String(user.Id)+"@gov2panel.subscribe", clientIp, asn, userAgent, "订阅方式已经停用，请提交工单联系管理员！")
		return
	}

	result := ""

	//app
	if req.FlagInfoHide {
		result = base64Sub(ctx, serviceArr, user, asn)
		ghttp.RequestFromCtx(ctx).Response.WriteExit(base64.StdEncoding.EncodeToString([]byte(result)))

		return nil, nil
	}

	switch req.Flag {

	case "1", "2", "6": //v2rayn //v2rayng //nekobox
		result = V2rayNGSub(ctx, serviceArr, user, asn)
	case "3": //shadowrocket
		result = ShadowrocketSub(ctx, serviceArr, user, asn)
	case "4": //clash
		result = ClashSub(ctx, serviceArr, user, asn)

		//https://www.clashverge.dev/guide/url_schemes.html#content-disposition
		//如果响应头中存在 profile-update-interval 字段，则配置文件的 更新间隔 将被设置为对应的值（单位: 小时）。
		ghttp.RequestFromCtx(ctx).Response.Header().Add("profile-update-interval", "24")
		//如果响应头中存在 subscription-userinfo 字段，则其对应的流量信息(单位: 字节)、到期信息(时间戳)会显示在订阅卡片上。
		ghttp.RequestFromCtx(ctx).Response.Header().Add("subscription-userinfo", fmt.Sprintf("upload=%d; download=%d; total=%d; expire=%d", user.U, user.D, user.TransferEnable, user.ExpiredAt.Unix()))
		//如果响应头中存在 profile-web-page-url 字段，则右键订阅卡片将会显示 首页 按钮。
		// ghttp.RequestFromCtx(ctx).Response.Header().Add("profile-web-page-url", "")

		ghttp.RequestFromCtx(ctx).Response.WriteExit([]byte(result))
	default:
		result = V2rayNGSub(ctx, serviceArr, user, asn)
	}

	ghttp.RequestFromCtx(ctx).Response.WriteExit(base64.StdEncoding.EncodeToString([]byte(result)))

	return nil, nil
}

// 订阅处理
func base64Sub(ctx context.Context, v2ServiceList []*entity.V2ProxyService, user *entity.V2User, asn string) (result string) {
	serviceRelayList, err := service.ServerRelay().GetServiceRelayListByShow(1)
	if err != nil {
		return
	}

	for _, v2Service := range v2ServiceList {

		v2rayUrl, err := service.ProxyService().GetV2rayUrl(ctx, v2Service, user, serviceRelayList, asn)
		if err != nil {
			return
		}
		result = result + v2rayUrl + "\n"

	}

	return
}

// ShadowrocketSub订阅
func ShadowrocketSub(ctx context.Context, serviceArr []*entity.V2ProxyService, user *entity.V2User, asn string) (result string) {
	result = result + fmt.Sprintf("STATUS=↑:%.2fGB,↓:%.2fGB,TOT:%.2fGBExpires:%s\n", utils.BytesToGB(user.U), utils.BytesToGB(user.D), utils.BytesToGB(user.TransferEnable), user.ExpiredAt)
	isExpired := user.ExpiredAt.Before(gtime.New(time.Now()))
	if isExpired {
		return
	}
	if user.TransferEnable-user.U-user.D <= 0 {
		return
	}

	result = result + base64Sub(ctx, serviceArr, user, asn)

	return
}

// v2rayNG订阅
func V2rayNGSub(ctx context.Context, serviceArr []*entity.V2ProxyService, user *entity.V2User, asn string) (result string) {

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

	result = result + base64Sub(ctx, serviceArr, user, asn)

	return
}

// https://wiki.metacubex.one/config/proxies/transport/?h=xhttp
// Clash订阅
func ClashSub(ctx context.Context, serviceArr []*entity.V2ProxyService, user *entity.V2User, asn string) (result string) {

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

	serviceRelayArr, err := service.ServerRelay().GetServiceRelayListByShow(1)
	if err != nil {
		return
	}

	for _, service := range serviceArr {
		if strings.Contains(service.Host, ",") {
			service.Host = utils.GetRandomString(service.Host)
		}

		// 定义中继服务器 $relay[]
		prefix := "$relay["
		suffix := "]"
		// 检查并截取
		if strings.HasPrefix(service.Host, prefix) && strings.HasSuffix(service.Host, suffix) {
			// 移除前缀和后缀，剩下就是中间的值
			val := service.Host[len(prefix) : len(service.Host)-len(suffix)]
			service.Host = getRandomRelayByFilter(serviceRelayArr, val, asn)
		}

		service.Host = strings.ReplaceAll(service.Host, "$uuid", user.Uuid)
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
			d["client-fingerprint"] = utils.GetRandomString(gconv.String(serviceJson["fp"]))
			d["encryption"] = gconv.String(serviceJson["encryption"])
			d["network"] = gconv.String(serviceJson["network"]) //传输协议
			switch gconv.String(serviceJson["network"]) {       //传输协议
			case "xhttp":
				d["xhttp-opts"] = map[string]interface{}{
					"path": utils.GetRandomString(gconv.String(serviceJson["path"])),
					"host": utils.GetRandomString(gconv.String(serviceJson["host"])),
					"mode": utils.GetRandomString(gconv.String(serviceJson["mode"])),
				}

			case "ws":
				d["ws-opts"] = map[string]interface{}{
					"path": utils.GetRandomString(gconv.String(serviceJson["path"])),
					"headers": map[string]interface{}{
						"Host": utils.GetRandomString(gconv.String(serviceJson["host"])),
					},
				}
			case "h2":
				d["h2-opts"] = map[string]interface{}{
					"host": []string{utils.GetRandomString(gconv.String(serviceJson["host"]))},
					"path": utils.GetRandomString(gconv.String(serviceJson["path"])),
				}
			case "grpc":
				d["grpc-opts"] = map[string]interface{}{
					"grpc-service-name": utils.GetRandomString(gconv.String(serviceJson["path"])),
				}

			}

			if gconv.String(serviceJson["type"]) == "http" { //伪装类型
				d["network"] = "http"
				d["http-opts"] = map[string]interface{}{
					"method":  "GET",
					"path":    []string{utils.GetRandomString(gconv.String(serviceJson["path"]))},
					"headers": map[string]interface{}{"Host": []string{utils.GetRandomString(gconv.String(serviceJson["host"]))}},
				}
			}

		case "vless":
			d["type"] = "vless"

			if gconv.String(serviceJson["security"]) == "reality" {
				d["tls"] = true
				d["servername"] = gconv.String(serviceJson["sni"])
				d["client-fingerprint"] = utils.GetRandomString(gconv.String(serviceJson["fp"]))
				d["skip-cert-verify"] = false
				d["sni"] = gconv.String(serviceJson["sni"])
				d["reality-opts"] = map[string]interface{}{
					"public-key":             gconv.String(serviceJson["pbk"]),
					"short-id":               utils.GetRandomString(gconv.String(serviceJson["sid"])),
					"support-x25519mlkem768": true,
				}
			}

			d["uuid"] = user.Uuid
			d["alterId"] = 0

			if gconv.String(serviceJson["tls"]) == "tls" { //tls
				d["tls"] = true
				d["alpn"] = strings.Split(gconv.String(serviceJson["alpn"]), ",")
				d["skip-cert-verify"] = false
				d["servername"] = gconv.String(serviceJson["sni"])
			}
			d["encryption"] = gconv.String(serviceJson["encryption"])
			d["network"] = gconv.String(serviceJson["network"]) //传输协议
			switch gconv.String(d["network"]) {                 //传输协议
			case "xhttp":
				d["xhttp-opts"] = map[string]interface{}{
					"path": utils.GetRandomString(gconv.String(serviceJson["path"])),
					"host": utils.GetRandomString(gconv.String(serviceJson["host"])),
					"mode": utils.GetRandomString(gconv.String(serviceJson["mode"])),
				}

			case "ws":
				d["ws-opts"] = map[string]interface{}{
					"path": utils.GetRandomString(gconv.String(serviceJson["path"])),
					"headers": map[string]interface{}{
						"Host": utils.GetRandomString(gconv.String(serviceJson["host"])),
					},
				}
			case "h2":
				d["h2-opts"] = map[string]interface{}{
					"host": []string{utils.GetRandomString(gconv.String(serviceJson["host"]))},
					"path": utils.GetRandomString(gconv.String(serviceJson["path"])),
				}
			case "grpc":
				d["grpc-opts"] = map[string]interface{}{
					"grpc-service-name": utils.GetRandomString(gconv.String(serviceJson["serviceName"])),
				}

			}

			if gconv.String(serviceJson["headerType"]) == "http" { //伪装类型
				d["network"] = "http"
				d["http-opts"] = map[string]interface{}{
					"method":  "GET",
					"path":    []string{utils.GetRandomString(gconv.String(serviceJson["path"]))},
					"headers": map[string]interface{}{"Host": []string{utils.GetRandomString(gconv.String(serviceJson["host"]))}},
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

		nodeInfoArr = append(nodeInfoArr, "  - "+j)
	}

	result = strings.ReplaceAll(result, "{{node_info}}", strings.Join(nodeInfoArr, "\n"))
	result = strings.ReplaceAll(result, "{{node_name}}", strings.Join(nodeNameArr, "\n"))

	return
}

// getRandomRelayByFilter 因为service.变量重名了，这样处理
func getRandomRelayByFilter(m []*entity.V2ServiceRelay, targetNameGroup string, targetAsn string) string {
	return service.ServerRelay().GetRandomRelayByFilter(m, targetNameGroup, targetAsn)
}
