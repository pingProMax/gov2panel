package admin

import (
	"context"

	v1 "gov2panel/api/admin/v1"
	"gov2panel/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Setting(ctx context.Context, req *v1.SettingReq) (res *v1.SettingRes, err error) {
	switch g.RequestFromCtx(ctx).Method {
	case "GET":
		setTplAdmin(ctx, "setting", nil)
	case "POST":
		res = &v1.SettingRes{}

		res.Data, err = service.Setting().GetSettingAllList(req.V2Setting)

		return

	default:
		return
	}
	return
}

func (c *ControllerV1) SettingAE(ctx context.Context, req *v1.SettingAEReq) (res *v1.SettingAERes, err error) {
	err = service.Setting().AESetting(&req.V2Setting)
	if err != nil {
		return res, err
	}
	return res, err
}

func (c *ControllerV1) SettingDel(ctx context.Context, req *v1.SettingDelReq) (res *v1.SettingDelRes, err error) {
	err = service.Setting().DelSetting(req.Codes)
	if err != nil {
		return res, err
	}
	return res, err
}
