package utils

import (
	"fagin/app"
	"github.com/mojocn/base64Captcha"
)

// 验证码类
//configJsonBody json request body.
type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

var store = base64Captcha.DefaultMemStore

// 生成验证码
func NewCaptcha() (id, b64s string, err error) {
	e := configJsonBody{}
	e.Id = app.RandString(16)
	e.DriverDigit = base64Captcha.DefaultDriverDigit
	driver := e.DriverDigit
	captcha := base64Captcha.NewCaptcha(driver, store)
	return captcha.Generate()
}

// 校验验证码
func VerifyCaptcha(id, code string, clear bool) bool {
	return store.Verify(id, code, clear)
}
