package request

import (
	"fagin/app/errno"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

type Request interface {
	Validate(*gin.Context) (map[string]string, bool)
	Message() map[string]string
	Attributes() map[string]string
}

type Validation struct {
	Request
}

func (v Validation) Validate(ctx *gin.Context) (map[string]string, bool) {
	return Validated(v.Request, ctx)
}

func (v Validation) ValidateUri(ctx *gin.Context) (map[string]string, bool) {
	return ValidateUri(v.Request, ctx)
}

// FileValidate 文件验证
// maxSize 限定大小
// typeFunc 判断文件类型方法
func (v Validation) FileValidate(ctx *gin.Context, maxSize int64, typeFunc func() (map[string]string, bool)) (map[string]string, bool) {
	return FileValidate(v.Request, ctx, maxSize, typeFunc)
}

func Validated(request Request, ctx *gin.Context) (map[string]string, bool) {
	err := ctx.ShouldBind(request)
	if err == nil {
		return nil, true
	}
	return validate(err, request), false
}

func ValidateUri(request Request, ctx *gin.Context) (map[string]string, bool) {
	err := ctx.ShouldBindUri(request)
	if err == nil {
		return nil, true
	}
	return validate(err, request), false
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

// FileValidate 文件验证
// maxSize 限定大小
// typeFunc 判断文件类型方法
func FileValidate(request Request, ctx *gin.Context, maxSize int64, typeFunc func() (map[string]string, bool)) (map[string]string, bool) {

	// 检查上传文件的大小
	if ok := ParseMultipartForm(ctx, maxSize); !ok {
		// 一般要前端严格限定上传大小
		return map[string]string{"file": errno.Serve.UploadSizeExceededErr.Message}, false
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

// ParseMultipartForm 检查提交的表单是否超过限定大小，
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
