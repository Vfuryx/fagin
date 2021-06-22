package admin_request

import (
	"fagin/pkg/request"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type uploadVideo struct {
	File               *multipart.FileHeader `form:"file" binding:"required"`
	request.Validation `binding:"-"`
}

func NewUploadVideo() *uploadVideo {
	r := new(uploadVideo)
	r.Request = r
	return r
}

func (uploadVideo) Message() map[string]string {
	return map[string]string{
		"File.required": "文件不能为空",
	}
}

func (uploadVideo) Attributes() map[string]string {
	return map[string]string{
		"File": "文件",
	}
}

// Validate 重写验证方法
func (r *uploadVideo) Validate(ctx *gin.Context) (map[string]string, bool) {
	const maxFileSize int64 = 100 << 20 // 限定大小
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
