package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CaptchaReq struct {
	g.Meta `path:"/captcha" tags:"Captcha" method:"get" summary:"验证码"`
}
type CaptchaRes struct {
	g.Meta       `mime:"application/json" example:"string"`
	CaptchaId    string `json:"captcha_id"`
	CaptchaImage string `json:"captcha_image"`
}

type VerifyCaptchaReq struct {
	g.Meta `path:"/verify_captcha" tags:"VerifyCaptcha" method:"get" summary:"效验验证码"`
	VerifyCaptcha
}
type VerifyCaptcha struct {
	Id          string
	VerifyValue string
}

type VerifyCaptchaRes struct {
	g.Meta `mime:"application/json" example:"string"`
}
