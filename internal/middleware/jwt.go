package middleware

import (
	"gov2panel/internal/model/model"
	"gov2panel/internal/service"
	"net/http"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/golang-jwt/jwt/v5"
)

// JWTAuth is a middleware that validates JWT tokens in the request header
// It checks for the presence and validity of the token, and sets the username
// in the context for downstream handlers
func AdminJWTAuth(r *ghttp.Request) {
	req := ghttp.RequestFromCtx(r.GetCtx())

	tokenString := r.Cookie.Get("jwt").String()
	if tokenString == "" {
		req.Response.RedirectTo("/", http.StatusFound)
		// r.SetError(gerror.NewCode(gcode.CodeNotAuthorized, "No token provided"))
		return
	}

	// Remove 'Bearer ' prefix if present
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	// Parse and validate the token
	claims := &model.JWTClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(g.Cfg().MustGet(gctx.New(), "jwtkey").Bytes()), nil
	})

	if err != nil || !token.Valid {
		req.Response.RedirectTo("/", http.StatusFound)
		// r.SetError(gerror.NewCode(gcode.CodeNotAuthorized, "Invalid token"))
		return
	}

	user, err := service.User().GetUserById(claims.TUserID)
	if err != nil {
		r.Response.WriteExit(err.Error())
	}
	if user.IsAdmin != 1 {
		r.Response.Write("you are a big sb！！！")
		return
	}

	remainingTime := claims.ExpiresAt.Unix() - time.Now().Unix()
	if remainingTime < 302400 {
		signedToken, _, err := service.User().CreateToken(r.GetCtx(), user)
		if err != nil {
			return
		}
		r.Cookie.Set("jwt", signedToken)
	}

	r.SetCtxVar("database_user", user)
	r.Assigns(gview.Params{
		"user": user,
	})

	r.Middleware.Next()
}

// JWTAuth is a middleware that validates JWT tokens in the request header
// It checks for the presence and validity of the token, and sets the username
// in the context for downstream handlers
func UserJWTAuth(r *ghttp.Request) {
	req := ghttp.RequestFromCtx(r.GetCtx())

	tokenString := r.Cookie.Get("jwt").String()
	if tokenString == "" {
		req.Response.RedirectTo("/", http.StatusFound)
		// r.SetError(gerror.NewCode(gcode.CodeNotAuthorized, "No token provided"))
		return
	}

	// Remove 'Bearer ' prefix if present
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	// Parse and validate the token
	claims := &model.JWTClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(g.Cfg().MustGet(gctx.New(), "jwtkey").Bytes()), nil
	})

	if err != nil || !token.Valid {
		req.Response.RedirectTo("/", http.StatusFound)
		// r.SetError(gerror.NewCode(gcode.CodeNotAuthorized, "Invalid token"))
		return
	}

	user, err := service.User().GetUserById(claims.TUserID)
	if err != nil {
		r.Response.WriteExit(err.Error())
	}
	if user.Banned == 1 {
		req.Cookie.Remove("jwt")
		req.Response.RedirectTo("/", http.StatusFound)
		req.ExitAll()
		return
	}

	remainingTime := claims.ExpiresAt.Unix() - time.Now().Unix()
	if remainingTime < 302400 {
		signedToken, _, err := service.User().CreateToken(r.GetCtx(), user)
		if err != nil {
			return
		}
		r.Cookie.Set("jwt", signedToken)
	}

	r.SetCtxVar("database_user", user)
	r.Assigns(gview.Params{
		"user": user,
	})

	r.Middleware.Next()
}
