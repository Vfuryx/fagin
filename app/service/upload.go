package service

import (
	"fagin/utils"
	"mime/multipart"
	"path"
	"time"
)

type uploadService struct {
	BasePath string
}

func NewUploadService(basePath string) *uploadService {
	return &uploadService{
		BasePath: basePath,
	}
}

func (u *uploadService) UploadFile(dist string, file *multipart.FileHeader) (string, error) {
	filePath := dist + time.Now().Format("2006123") + "/" + utils.RandString(20) + path.Ext(file.Filename)
	absolutePath := u.BasePath + filePath
	utils.MkdirAll(absolutePath)
	err := utils.SaveUploadedFile(file, absolutePath)
	return filePath, err
}
