package user

import (
	"fmt"
	"gov2panel/internal/service"
	"net/http"
	"time"

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

func (s *middlewareService) AuthAdmin(r *ghttp.Request) {
	Auth().MiddlewareFunc()(r)
	user, err := service.User().GetUserById(r.Get("TUserID").Int())
	if err != nil {
		r.Response.WriteExit(err.Error())
	}
	if user.IsAdmin != 1 {
		r.Response.Write("you are a big sb！！！")
		return
	}

	mapClaims, _, err := Auth().GetClaimsFromJWT(r.GetCtx())
	if err != nil {
		return
	}

	if int64(mapClaims["exp"].(float64))-time.Now().UnixMilli() < 43200000 { //自动刷新token
		token, expire, err := Auth().RefreshToken(r.GetCtx())
		if err != nil {
			r.Response.WriteExit(err.Error())
		}
		r.Cookie.SetCookie("jwt", token, r.Server.GetCookieDomain(), r.Server.GetCookiePath(), time.Duration(expire.UnixMilli()))

		fmt.Println("刷新token", token)
	}

	r.SetCtxVar("database_user", user)
	r.Assigns(gview.Params{
		"user": user,
	})
	r.Middleware.Next()
}

func (s *middlewareService) AuthUser(r *ghttp.Request) {
	Auth().MiddlewareFunc()(r)
	user, err := service.User().GetUserById(r.Get("TUserID").Int())
	if err != nil {
		r.Response.WriteExit(err.Error())
	}

	if user.Banned == 1 {
		service.User().Logout(r.GetCtx())
		ghttp.RequestFromCtx(r.GetCtx()).Cookie.Remove("jwt")
		ghttp.RequestFromCtx(r.GetCtx()).Response.RedirectTo("/", http.StatusFound)
		ghttp.RequestFromCtx(r.GetCtx()).ExitAll()
		return
	}

	mapClaims, _, err := Auth().GetClaimsFromJWT(r.GetCtx())
	if err != nil {
		return
	}

	if int64(mapClaims["exp"].(float64))-time.Now().UnixMilli() < 43200000 { //自动刷新token
		token, expire, err := Auth().RefreshToken(r.GetCtx())
		if err != nil {
			r.Response.WriteExit(err.Error())
		}
		r.Cookie.SetCookie("jwt", token, r.Server.GetCookieDomain(), r.Server.GetCookiePath(), time.Duration(expire.UnixMilli()))

		fmt.Println("刷新token", token)
	}

	r.SetCtxVar("database_user", user)
	r.Assigns(gview.Params{
		"user": user,
	})
	r.Middleware.Next()
}
