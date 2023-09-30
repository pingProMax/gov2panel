package user

import (
	"context"
	"net/http"

	v1 "gov2panel/api/user/v1"
	"gov2panel/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) User(ctx context.Context, req *v1.UserReq) (res *v1.UserRes, err error) {
	res = &v1.UserRes{}
	setTplUser(ctx, "user", g.Map{"data": res.Data})
	return
}

func (c *ControllerV1) UserUpPasswd(ctx context.Context, req *v1.UserUpPasswdReq) (res *v1.UserUpPasswdRes, err error) {
	return service.User().UpUserPasswdById(req)
}

func (c *ControllerV1) ResetTokenAndUuid(ctx context.Context, req *v1.ResetTokenAndUuidReq) (res *v1.ResetTokenAndUuidRes, err error) {
	res = &v1.ResetTokenAndUuidRes{}
	err = service.User().ResetTokenAndUuidById(req.TUserID)
	return
}

func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	service.User().Logout(ctx)
	ghttp.RequestFromCtx(ctx).Cookie.Remove("jwt")
	ghttp.RequestFromCtx(ctx).Response.RedirectTo("/", http.StatusFound)
	ghttp.RequestFromCtx(ctx).ExitAll()
	return
}

func (c *ControllerV1) Refresh(ctx context.Context, req *v1.RefreshReq) (res *v1.RefreshRes, err error) {
	res = &v1.RefreshRes{}
	res.Token, res.Expire = service.User().Refresh(ctx)
	return
}
