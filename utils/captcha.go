package utils

import (
	"github.com/mojocn/base64Captcha"
)

// 验证码类
// configJsonBody json request body.
type configJSONBody struct {
	ID            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

var store = base64Captcha.DefaultMemStore

// NewCaptcha 生成验证码
func NewCaptcha() (id, b64s string, err error) {
	const num = 16

	e := configJSONBody{}
	e.ID = RandString(num)
	e.DriverDigit = base64Captcha.DefaultDriverDigit
	driver := e.DriverDigit
	captcha := base64Captcha.NewCaptcha(driver, store)

	return captcha.Generate()
}

// VerifyCaptcha 校验验证码
func VerifyCaptcha(id, code string, clear bool) bool {
	return store.Verify(id, code, clear)
}
