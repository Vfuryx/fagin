package admin

import (
	"fagin/app"
	"fagin/app/errno"
	admin_request "fagin/app/requests/admin"
	"fagin/app/service"
	"fagin/config"
	"fagin/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"time"
)

type companyIntroductionController struct{}

var CompanyIntroductionController companyIntroductionController

func (companyIntroductionController) ShowCompanyIntroduction(ctx *gin.Context) {
	column := []string{
		"id", "name", "image", "content", "status",
	}

	ci, err := service.CompanyIntroduction.ShowCompanyIntroduction(1, column)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.Api.ErrCompanyIntroduction, nil)
		return
	}

	app.JsonResponse(ctx, errno.OK, gin.H{
		"id":      ci.ID,
		"name":    ci.Name,
		"image":   ci.Image,
		"content": ci.Content,
		"status":  ci.Status,
	})
	return
}

func (companyIntroductionController) UpdateCompanyIntroduction(ctx *gin.Context) {
	var r admin_request.UpdateCompanyIntroduction
	if data, ok := r.Validate(ctx); !ok {
		app.JsonResponse(ctx, errno.InternalServerError, data)
		return
	}

	data := gin.H{
		"name":    r.Name,
		"image":   r.Image,
		"content": r.Content,
		"status":  *r.Status,
	}

	err := service.CompanyIntroduction.UpdateCompanyIntroduction(1, data)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.Api.ErrUpdateCompanyIntroduction, err)
		return
	}

	app.JsonResponse(ctx, errno.OK, nil)
	return
}

// 上传
func (companyIntroductionController) Upload(ctx *gin.Context) {
	var r admin_request.UploadCompanyImage
	if data, ok := r.Validate(ctx); !ok {
		app.JsonResponseWithStatus(ctx, http.StatusBadRequest, errno.Api.ErrBind, data)
		return
	}

	filePath := "/web/company/" + time.Now().Format("2006123") + app.RandString(20) + path.Ext(r.File.Filename)

	// 是否创建目录
	dir := path.Dir(config.App.PublicPath + filePath)
	_, err := os.Stat(dir)
	if err != nil && os.IsNotExist(err) {
		log.Log.Errorln(err)
		_ = os.MkdirAll(dir, os.ModePerm)
	}

	err = ctx.SaveUploadedFile(r.File, config.App.PublicPath+filePath)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.Api.ErrUploadFile, nil)
		return
	}

	app.JsonResponse(ctx, errno.OK, gin.H{"path": filePath})

	return
}
