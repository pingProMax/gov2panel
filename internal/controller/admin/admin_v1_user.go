package admin

import (
	"context"
	"fmt"
	"net/http"

	v1 "gov2panel/api/admin/v1"
	"gov2panel/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) User(ctx context.Context, req *v1.UserReq) (res *v1.UserRes, err error) {

	switch g.RequestFromCtx(ctx).Method {
	case "GET":
		setTplAdmin(ctx, "user", nil)
	case "POST":
		res = &v1.UserRes{}
		fmt.Println(req.DS)
		res.UserList, res.Totle, err = service.User().GetUserList(req, req.Sort, req.Order, req.Offset, req.Limit)
		return

	default:
		return
	}
	return
}

func (c *ControllerV1) UserAE(ctx context.Context, req *v1.UserAEReq) (res *v1.UserAERes, err error) {
	err = service.User().AEUser(&req.V2User)
	if err != nil {
		return res, err
	}
	return res, err
}

func (c *ControllerV1) UserDel(ctx context.Context, req *v1.UserDelReq) (res *v1.UserDelRes, err error) {
	err = service.User().DelUser(req.Ids)
	if err != nil {
		return res, err
	}
	return res, err
}

func (c *ControllerV1) UserUpBanned1(ctx context.Context, req *v1.UserUpBanned1Req) (res *v1.UserUpBanned1Res, err error) {
	err = service.User().UpUserBanned1(req.Ids)
	if err != nil {
		return res, err
	}
	return res, err
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

func (c *ControllerV1) ResetTokenAndUuidById(ctx context.Context, req *v1.ResetTokenAndUuidReq) (res *v1.ResetTokenAndUuidRes, err error) {
	res = &v1.ResetTokenAndUuidRes{}
	err = service.User().ResetTokenAndUuidById(req.UserId)
	return
}
