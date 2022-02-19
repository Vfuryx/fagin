package request

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// 定义一个全局翻译器T
var trans ut.Translator
var zhT = zh.New() // 中文翻译器
var enT = en.New() // 英文翻译器

// InitTrans 初始化翻译器
func InitTrans(locale string) (err error) {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return
	}

	v.RegisterTagNameFunc(func(fld reflect.StructField) string { // 注册一个获取json tag的自定义方法
		const num = 2
		name := strings.SplitN(fld.Tag.Get("json"), ",", num)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// 第一个参数是备用（fallback）的语言环境 后面的参数是应该支持的语言环境（支持多个）
	uni := ut.New(enT, zhT, enT)

	if trans, ok = uni.GetTranslator(locale); !ok { // locale 通常取决于 http 请求头的 'Accept-Language' 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
	}

	switch locale { // 注册翻译器
	case "zh":
		err = zhTranslations.RegisterDefaultTranslations(v, trans)
	default:
		err = enTranslations.RegisterDefaultTranslations(v, trans)
	}

	return
}
