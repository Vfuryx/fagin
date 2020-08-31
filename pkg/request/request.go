package request

import (
	"fagin/app"
	"fagin/app/errno"
	"fagin/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
)

type Request interface {
	Validate(*gin.Context) (map[string]string, bool)
	Message() map[string]string
	Attributes() map[string]string
}

type Validate struct{}

func (Validate) Validate(request Request, ctx *gin.Context) (map[string]string, bool) {
	err := ctx.ShouldBind(request)
	if err != nil {
		data := validate(err, request)
		return data, false
	}
	return nil, true
}

func validate(err error, request Request) map[string]string {
	var data = map[string]string{}
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errs := err.(validator.ValidationErrors)
			// 获取翻译数据
			ts := errs.Translate(trans)
			for _, value := range errs {
				// 查询自定义的message
				msg, ok := request.Message()[value.StructField()+`.`+value.Tag()]
				if !ok { // 没有
					// 使用翻译的
					msg = strings.Replace(ts[value.Namespace()], value.Field(), request.Attributes()[value.StructField()], 1)
				}
				data[value.Field()] = msg
			}
		default:
			data["error"] = err.Error()
		}
	}
	return data
}

// 文件验证
// maxSize 限定大小
// typeFunc 判断文件类型方法
func (Validate) FileValidate(request Request, ctx *gin.Context, maxSize int64, typeFunc func() (map[string]string, bool)) (map[string]string, bool) {

	// 检查上传文件的大小
	if ok := ParseMultipartForm(ctx, maxSize); !ok {
		// 一般要前端严格限定上传大小
		return map[string]string{"file": errno.Api.ErrUploadSizeExceeded.Message}, false
	}

	// 绑定
	err := ctx.ShouldBind(request)
	if err != nil {
		_ = ctx.Request.Body.Close()
		// 错误输出处理
		data := fileValidate(err, request)
		return data, false
	}
	return typeFunc()
}

// 自定义处理方法
func fileValidate(err error, request Request) map[string]string {
	var data = map[string]string{}
	// 字段名
	var name = "file"
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errs := err.(validator.ValidationErrors)
			// 获取翻译数据
			ts := errs.Translate(trans)
			for _, value := range errs {
				// 查询自定义的message
				msg, ok := request.Message()[value.StructField()+`.`+value.Tag()]
				if !ok { // 没有
					// 使用翻译的
					msg = strings.Replace(ts[value.Namespace()], value.Field(), request.Attributes()[value.StructField()], 1)
				}
				data[value.Field()] = msg
			}
		default:
			// default
			data[name] = "上传文件错误"

			if err == http.ErrMissingFile {
				data[name] = "不是文件"
			}
			if err.Error() == "multipart: NextPart: EOF" {
				data[name] = "文件不能为空"
			}
			if err.Error() == "http: request body too large" {
				data[name] = "上传的文件过大"
			}
		}
	}
	return data
}

// 检查提交的表单是否超过限定大小，
// false 响应的状态码为 400 或 500 ，否则会出现断开链接的结果
func ParseMultipartForm(ctx *gin.Context, maxMemory int64) bool {
	if ctx.Request.ContentLength > maxMemory {
		return false
	}
	ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, maxMemory)
	err := ctx.Request.ParseMultipartForm(maxMemory)
	if err != nil {
		return false
	}
	return true
}

func CreateRequestTemplate(path, name string) {
	filePath := config.App.AppPath + "/requests/" + path + ".go"
	sl := strings.Split(filePath, "/")
	packageName := sl[len(sl)-2]
	dirPath := strings.Join(sl[:len(sl)-1], "/")

	//os.Stat获取文件信息
	if _, err := os.Stat(filePath); err == nil {
		panic("文件已存在")
	}

	// 创建文件夹 可以多层
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	const temp = `package %s_request

import (
	"github.com/gin-gonic/gin"
	"%[3]s/pkg/request"
)

type %[2]s struct {}

var _ request.Request = &%[2]s{}

func (%[2]s) Message() map[string]string {
	return map[string]string{
	}
}

func (%[2]s) Attributes() map[string]string {
	return map[string]string{
	}
}

func (r *%[2]s) Validate(ctx *gin.Context) (map[string]string, bool) {
	var v request.Validate
	return v.Validate(r, ctx)
}
`
	content := fmt.Sprintf(temp, packageName, app.Camel(name), config.App.Name)

	if _, err = file.WriteString(content); err != nil {
		panic(err)
	}

	if err = file.Close(); err != nil {
		panic(err)
	}

	log.Printf("request create run successfully")
}

// 定义一个全局翻译器T
var trans ut.Translator

// 翻译
func InitTrans(locale string) (err error) {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return
	}
	// 注册一个获取json tag的自定义方法
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	zhT := zh.New() // 中文翻译器
	enT := en.New() // 英文翻译器

	// 第一个参数是备用（fallback）的语言环境
	// 后面的参数是应该支持的语言环境（支持多个）
	// uni := ut.New(zhT, zhT) 也是可以的
	uni := ut.New(enT, zhT, enT)

	// locale 通常取决于 http 请求头的 'Accept-Language'

	// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
	trans, ok = uni.GetTranslator(locale)
	if !ok {
		return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
	}

	// 注册翻译器
	switch locale {
	case "en":
		err = enTranslations.RegisterDefaultTranslations(v, trans)
	case "zh":
		err = zhTranslations.RegisterDefaultTranslations(v, trans)
	default:
		err = enTranslations.RegisterDefaultTranslations(v, trans)
	}
	return
}
