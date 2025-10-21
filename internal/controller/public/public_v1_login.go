package public

import (
	"context"

	v1 "gov2panel/api/public/v1"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	r := g.RequestFromCtx(ctx)
	switch r.Method {
	case "GET":
		setTplPublc(ctx, "login", nil)
	case "POST":
		res = &v1.LoginRes{}
		var user *entity.V2User
		user, err = service.User().Login(req.UserName, req.Passwd) //不检查错误，统一返回错误
		if err != nil {
			return
		}

		if user == nil {
			return
		}

		signedToken, claims, errr := service.User().CreateToken(ctx, user)
		if errr != nil {
			return nil, gerror.NewCode(gcode.CodeInternalError, "Failed to generate token")
		}
		res.Token, res.Expire = signedToken, claims.ExpiresAt.Time
		r.Cookie.Set("jwt", signedToken)
		return

	default:
		return
	}

	return
}
