package public

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	v1 "gov2panel/api/public/v1"
	"gov2panel/internal/service"

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
	case "v2rayn", "v2rayng":
		// base64编码   单个：协议://base64编码

		for _, service := range serviceArr {
			serviceJson := make(map[string]interface{})
			json.Unmarshal([]byte(service.ServiceJson), &serviceJson)

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

			result = result + fmt.Sprintf("%s://%s\n", strings.Split(service.Agreement, "/")[1], base64.StdEncoding.EncodeToString(ds))
		}

	}

	ghttp.RequestFromCtx(ctx).Response.WriteExit(base64.StdEncoding.EncodeToString([]byte(result)))

	return nil, nil
}
