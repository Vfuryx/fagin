package admin_request

import (
	"fagin/pkg/request"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type UploadWebsiteConfigPic struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

var _ request.Request = &UploadWebsiteConfigPic{}

func (UploadWebsiteConfigPic) FieldMap() map[string]string {
	return map[string]string{
		"File": "required",
	}
}

func (UploadWebsiteConfigPic) Message() map[string]string {
	return map[string]string{
		"File.required": "文件不能为空",
	}
}

func (UploadWebsiteConfigPic) Attributes() map[string]string {
	return map[string]string{
		"File": "文件",
	}
}

func (r *UploadWebsiteConfigPic) Validate(ctx *gin.Context) (map[string]string, bool) {
	const maxFileSize int64 = 20 << 20 // 限定大小 20M
	var v request.Validate
	return v.FileValidate(r, ctx, maxFileSize, func() (map[string]string, bool) {
		// 判断文件类型
		fmt.Println(r.File.Header)
		if r.File.Header.Get("Content-Type") != "image/jpeg" {
			data := map[string]string{"file": "请上传 jpg 文件"}
			return data, false
		}
		return nil, true
	})
}
