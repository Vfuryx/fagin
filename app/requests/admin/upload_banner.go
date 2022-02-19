package request

import (
	"fagin/pkg/request"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type UploadBanner struct {
	File *multipart.FileHeader `form:"file" binding:"required"`

	request.Validation `binding:"-"`
}

func NewUploadBanner() *UploadBanner {
	r := new(UploadBanner)
	r.SetRequest(r)

	return r
}

func (*UploadBanner) Message() map[string]string {
	return map[string]string{
		// "File.required": "文件不能为空",
	}
}

func (*UploadBanner) Attributes() map[string]string {
	return map[string]string{
		"File": "文件",
	}
}

func (r *UploadBanner) Validate(ctx *gin.Context) (map[string]string, bool) {
	const maxFileSize int64 = 20 << 20 // 限定大小 20M

	return request.FileValidate(r, ctx, maxFileSize, func() (map[string]string, bool) {
		// 判断文件类型
		if r.File.Header.Get("Content-Type") != "image/jpeg" {
			data := map[string]string{"file": "请上传 jpg 文件"}
			return data, false
		}
		return nil, true
	})
}
