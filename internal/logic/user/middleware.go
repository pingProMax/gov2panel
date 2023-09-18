package user

import (
	"gov2panel/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gview"
)

type middlewareService struct{}

var middleware = middlewareService{}

func Middleware() *middlewareService {
	return &middleware
}

// 允许跨域
func (s *middlewareService) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func (s *middlewareService) Auth(r *ghttp.Request) {
	Auth().MiddlewareFunc()(r)
	user, err := service.User().GetUserById(r.Get("TUserID").Int())
	if err != nil {
		r.Response.Write(err.Error())
		return
	}
	r.SetCtxVar("database_user", user)
	r.Assigns(gview.Params{
		"user": user,
	})
	r.Middleware.Next()
}
