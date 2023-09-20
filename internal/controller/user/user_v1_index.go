package user

import (
	"context"
	"fmt"

	v1 "gov2panel/api/user/v1"
	"gov2panel/internal/service"
	"gov2panel/internal/utils"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Index(ctx context.Context, req *v1.IndexReq) (res *v1.IndexRes, err error) {
	user, err := service.User().GetUserById(g.RequestFromCtx(ctx).Get("TUserID").Int())
	if err != nil {
		g.RequestFromCtx(ctx).Response.Write(err.Error())
		return
	}
	plan, _ := service.Plan().GetPlanById(user.GroupId)
	setTplUser(ctx,
		"index",
		g.Map{
			"data":               plan,
			"transfer_enable":    fmt.Sprintf("%.2f", utils.BytesToGB(user.TransferEnable)),
			"total_used_traffic": fmt.Sprintf("%.2f", utils.BytesToGB(user.U+user.D)),
		})
	return
}
