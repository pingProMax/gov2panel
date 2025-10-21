package user

import (
	"context"
	"fmt"
	"strconv"
	"time"

	v1 "gov2panel/api/user/v1"
	"gov2panel/internal/utils"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

func (c *ControllerV1) Flow(ctx context.Context, req *v1.FlowReq) (res *v1.FlowRes, err error) {

	res = &v1.FlowRes{}
	res.Data = make([]map[string]interface{}, 0)

	timeNow := time.Now()

	for i := 0; i < 7; i++ {
		if i > 0 {
			timeNow = timeNow.Add(-time.Duration(1) * 24 * time.Hour)
		}
		dataStr := fmt.Sprintf("%s%s%s", strconv.Itoa(timeNow.Year()), strconv.Itoa(int(timeNow.Month())), strconv.Itoa(timeNow.Day()))
		ketStr := fmt.Sprintf("USER_%s_%s_FLOW_UPLOAD", strconv.Itoa(c.getUser(ctx).Id), dataStr)
		userFlow, err := gcache.Get(ctx, ketStr)
		if err != nil {
			return res, err
		}

		res.Data = append(res.Data, map[string]interface{}{
			"date": dataStr,
			"flow": utils.BytesToGB(userFlow.Int64()),
		})

	}

	setTplUser(ctx, "flow", g.Map{
		"data": res.Data,
	})

	return
}
