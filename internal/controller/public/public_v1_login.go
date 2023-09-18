package public

import (
	"context"

	v1 "gov2panel/api/public/v1"
	"gov2panel/internal/logic/user"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {

	switch g.RequestFromCtx(ctx).Method {
	case "GET":
		setTplPublc(ctx, "login", nil)
	case "POST":
		res = &v1.LoginRes{}
		res.Token, res.Expire = user.Auth().LoginHandler(ctx)

		return

	default:
		return
	}

	return
}
