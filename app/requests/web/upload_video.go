package request

import (
	"fagin/pkg/request"
	"fmt"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type UploadVideo struct {
	File *multipart.FileHeader `form:"file" binding:"required"`

	request.Validation `binding:"-"`
}

func NewUploadVideo() *UploadVideo {
	r := new(UploadVideo)
	r.SetRequest(r)

	return r
}

func (*UploadVideo) Message() map[string]string {
	return map[string]string{
		// "File.required": "文件不能为空",
	}
}

func (*UploadVideo) Attributes() map[string]string {
	return map[string]string{
		"File": "文件",
	}
}

func (r *UploadVideo) Validate(ctx *gin.Context) (map[string]string, bool) {
	const maxFileSize int64 = 200 << 20 // 限定大小 200M

	return request.FileValidate(r, ctx, maxFileSize, func() (map[string]string, bool) {
		// 判断文件类型
		fmt.Println(r.File.Header)
		if r.File.Header.Get("Content-Type") != "video/mp4" {
			data := map[string]string{"file": "请上传 mp4 文件"}
			return data, false
		}
		return nil, true
	})
}

// func (r *UploadVideo) Validate(ctx *gin.Context) (map[string]string, bool) {
//
// 	// 绑定
// 	err := ctx.ShouldBind(r)
// 	if err != nil {
// 		_ = ctx.Request.Body.Close()
// 		// 错误输出处理
// 		data := validate(err, r)
// 		return data, false
// 	}
//
// 	// 判断文件类型
// 	if r.File.Header.Get("Content-Type") != "video/mp4" {
// 		data := map[string]string{"file": "请上传 mp4 文件"}
// 		return data, false
// 	}
//
// 	return nil, true
// }
//
// // 自定义处理方法
// func validate(err error, request request.Request) map[string]string {
// 	var data = map[string]string{}
// 	// 字段名
// 	var name = "file"
// 	if err != nil {
// 		switch err.(type) {
// 		case validator.ValidationErrors:
// 			for _, value := range err.(validator.ValidationErrors) {
// 				name := request.FieldMap()[value.Field()]
// 				data[name] = request.Message()[value.Field()+`.`+value.Tag()]
// 			}
// 		default:
// 			// default
// 			data[name] = "上传文件错误"
//
// 			if err == http.ErrMissingFile {
// 				data[name] = "不是文件"
// 			}
// 			if err.Error() == "multipart: NextPart: EOF" {
// 				data[name] = "文件不能为空"
// 			}
// 			if err.Error() == "http: request body too large" {
// 				data[name] = "上传的文件过大"
// 			}
// 		}
// 	}
// 	return data
// }
