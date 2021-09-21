package admin

import (
	"fagin/app/errno"
	adminRequest "fagin/app/requests/admin"
	"fagin/app/service"
	"fagin/config"
	"github.com/gin-gonic/gin"
)

type companyIntroductionController struct {
	BaseController
}

var CompanyIntroductionController companyIntroductionController

func (cc *companyIntroductionController) ShowCompanyIntroduction(ctx *gin.Context) {
	column := []string{
		"id", "name", "image", "content", "status",
	}
	ci, err := service.CompanyIntroduction.ShowCompanyIntroduction(1, column)
	if err != nil {
		cc.ResponseJsonErrLog(ctx, errno.CtxShowErr, err, nil)
		return
	}

	cc.ResponseJsonOK(ctx, gin.H{
		"id":      ci.ID,
		"name":    ci.Name,
		"image":   ci.Image,
		"content": ci.Content,
		"status":  ci.Status,
	})
	return
}

func (cc *companyIntroductionController) UpdateCompanyIntroduction(ctx *gin.Context) {
	var r = adminRequest.NewUpdateCompanyIntroduction()
	if data, ok := r.Validate(ctx); !ok {
		cc.ResponseJsonErr(ctx, errno.ReqErr, data)
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
		cc.ResponseJsonErrLog(ctx, errno.CtxUpdateErr, err, nil)
		return
	}

	cc.ResponseJsonOK(ctx, nil)
	return
}

// Upload 上传
func (cc *companyIntroductionController) Upload(ctx *gin.Context) {
	var r = adminRequest.NewUploadCompanyImage()
	if data, ok := r.Validate(ctx); !ok {
		cc.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	upload := service.NewUploadService(config.App.PublicPath)
	path, err := upload.UploadFile("/web/company/", r.File)
	if err != nil {
		cc.ResponseJsonErrLog(ctx, errno.ReqUploadFileErr, err, nil)
		return
	}

	cc.ResponseJsonOK(ctx, gin.H{"path": "/public/" + path})
}
