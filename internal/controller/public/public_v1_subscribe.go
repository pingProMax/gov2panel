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
				"add":  service.Host,
				"ps":   service.Name,
				"port": service.Port,
				"id":   user.Uuid,
				"aid":  "0",
				"net":  serviceJson["network"].(string),
				"type": "none",
				"path": "",
				"tls":  "",
				"host": "",
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
