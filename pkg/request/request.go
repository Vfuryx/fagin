package request

import (
	"errors"
	"fagin/app/errno"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

// Request 请求接口
type Request interface {
	Validate(*gin.Context) (map[string]string, bool)
	Message() map[string]string
	Attributes() map[string]string
}

// Validation 验证
type Validation struct {
	Request
}

// SetRequest 设置请求实例
func (v *Validation) SetRequest(r Request) {
	v.Request = r
}

// Validate 验证
func (v *Validation) Validate(ctx *gin.Context) (map[string]string, bool) {
	return Validated(v.Request, ctx)
}

// ValidateUri 验证 Uri
func (v *Validation) ValidateUri(ctx *gin.Context) (map[string]string, bool) {
	return ValidateUri(v.Request, ctx)
}

// FileValidate 文件验证
// maxSize 限定大小
// typeFunc 判断文件类型方法
func (v *Validation) FileValidate(ctx *gin.Context, maxSize int64, typeFunc func() (map[string]string, bool)) (map[string]string, bool) {
	return FileValidate(v.Request, ctx, maxSize, typeFunc)
}

// Validated 验证
func Validated(request Request, ctx *gin.Context) (map[string]string, bool) {
	err := ctx.ShouldBind(request)
	if err == nil {
		return nil, true
	}
	return validate(err, request), false
}

// ValidateUri 验证 Uri
func ValidateUri(request Request, ctx *gin.Context) (map[string]string, bool) {
	err := ctx.ShouldBindUri(request)
	if err == nil {
		return nil, true
	}
	return validate(err, request), false
}

func validationErrorMessageHandle(request Request, errs validator.ValidationErrors) map[string]string {
	var data = map[string]string{}
	// 获取 Attributes
	attributes := request.Attributes()
	for _, value := range errs {
		// 查询自定义的message
		msg, ok := request.Message()[value.StructField()+`.`+value.Tag()]
		if !ok {
			// 获取翻译数据
			ts := errs.Translate(trans)
			// 替换字段名称
			msg = strings.Replace(ts[value.Namespace()], value.Field(), attributes[value.StructField()], 1)
		} else {
			// 占位符 :attribute 替换
			msg = replaceAttribute(msg, attributes[value.StructField()])
			// 占位符 :input 替换
			if v, ok := value.Value().(string); ok {
				msg = replaceInput(msg, v)
			}
			// 根据规则替换占位符
			msg = callReplacer(msg, value.StructField(), value.Tag(), value.Param())
		}
		data[value.Field()] = msg
	}
	return data
}

func validate(err error, request Request) map[string]string {
	var data = map[string]string{}
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			data = validationErrorMessageHandle(request, err.(validator.ValidationErrors))
		default:
			data["error"] = err.Error()
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
		return map[string]string{"file": errno.ReqUploadSizeExceededErr.Error()}, false
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

// 自定义处理方法
func fileValidate(err error, request Request) map[string]string {
	var data = map[string]string{}
	// 字段名
	var name = "file"
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			data = validationErrorMessageHandle(request, err.(validator.ValidationErrors))
		default:
			// default
			data[name] = "上传文件错误"

			if errors.Is(err, http.ErrMissingFile) {
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
