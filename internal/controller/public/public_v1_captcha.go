package public

import (
	"context"
	"image/color"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/mojocn/base64Captcha"

	v1 "gov2panel/api/public/v1"
)

var store = base64Captcha.DefaultMemStore

// 获取验证码
func MakeCaptcha() (string, string, error) {
	//定义一个driver
	var driver base64Captcha.Driver
	//创建一个字符串类型的验证码驱动DriverString, DriverChinese :中文驱动
	driverString := base64Captcha.DriverString{
		Height:          40,                                     //高度
		Width:           100,                                    //宽度
		NoiseCount:      0,                                      //干扰数
		ShowLineOptions: 2 | 4,                                  //展示个数
		Length:          4,                                      //长度
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm", //验证码随机字符串来源
		BgColor: &color.RGBA{ // 背景颜色
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"}, // 字体
	}
	driver = driverString.ConvertFonts()
	//生成验证码
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	return id, b64s, err
}

// 校验验证码
func VerifyCaptcha(id string, VerifyValue string) bool {
	// 参数说明: id 验证码id, verifyValue 验证码的值, true: 验证成功后是否删除原来的验证码
	if store.Verify(id, VerifyValue, true) {
		return true
	} else {
		return false
	}
}

func (c *ControllerV1) Captcha(ctx context.Context, req *v1.CaptchaReq) (res *v1.CaptchaRes, err error) {
	res = &v1.CaptchaRes{}
	res.CaptchaId, res.CaptchaImage, err = MakeCaptcha()
	return
}

func (c *ControllerV1) VerifyCaptcha(ctx context.Context, req *v1.VerifyCaptchaReq) (res *v1.VerifyCaptchaRes, err error) {
	res = &v1.VerifyCaptchaRes{}
	if VerifyCaptcha(req.Id, req.VerifyValue) {
		return res, err
	}
	return res, gerror.NewCode(gcode.CodeValidationFailed)
}
