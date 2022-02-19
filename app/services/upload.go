package services

import (
	"fagin/pkg/errorw"
	"fagin/utils"
	"mime/multipart"
	"path"
	"time"
)

type UploadService struct {
	BasePath string
}

// NewUploadService 上传服务
func NewUploadService(basePath string) *UploadService {
	return &UploadService{
		BasePath: basePath,
	}
}

func (u *UploadService) UploadFile(dist string, file *multipart.FileHeader) (string, error) {
	const num = 20
	filePath := dist + time.Now().Format("2006123") + "/" + utils.RandString(num) + path.Ext(file.Filename)
	absolutePath := u.BasePath + filePath
	utils.MkdirAll(absolutePath)
	err := utils.SaveUploadedFile(file, absolutePath)

	return filePath, errorw.UP(err)
}
