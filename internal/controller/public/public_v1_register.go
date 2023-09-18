package public

import (
	"context"
	"errors"
	v1 "gov2panel/api/public/v1"
	"gov2panel/internal/logic/user"
	"gov2panel/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {

	switch g.RequestFromCtx(ctx).Method {
	case "GET":
		setTplPublc(ctx, "register", g.Map{"code": req.Code})
	case "POST":
		res = &v1.RegisterRes{}
		if !VerifyCaptcha(req.Id, req.VerifyValue) {
			return res, errors.New("验证码错误")
		}

		err = service.User().RegisterUser(req.UserName, req.Passwd, req.CommissionCode)
		if err != nil {
			return nil, err
		}

		res.Token, res.Expire = user.Auth().LoginHandler(ctx)

		return

	default:
		return
	}

	return
}
