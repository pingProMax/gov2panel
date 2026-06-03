// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package public

import (
	"context"

	"gov2panel/api/public/v1"
)

type IPublicV1 interface {
	Captcha(ctx context.Context, req *v1.CaptchaReq) (res *v1.CaptchaRes, err error)
	VerifyCaptcha(ctx context.Context, req *v1.VerifyCaptchaReq) (res *v1.VerifyCaptchaRes, err error)
	Index(ctx context.Context, req *v1.IndexReq) (res *v1.IndexRes, err error)
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
	EPayNotify(ctx context.Context, req *v1.EPayNotifyReq) (res *v1.EPayNotifyRes, err error)
	Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error)
	Subscribe(ctx context.Context, req *v1.SubscribeReq) (res *v1.SubscribeRes, err error)
}
