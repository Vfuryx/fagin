package admin

import (
	"fagin/app/errno"
	adminRequest "fagin/app/requests/admin"
	"fagin/app/services"
	"fagin/config"

	"github.com/gin-gonic/gin"
)

type companyIntroductionController struct {
	BaseController
}

var CompanyIntroductionController companyIntroductionController

func (ctr *companyIntroductionController) ShowCompanyIntroduction(ctx *gin.Context) {
	column := []string{
		"id", "name", "image", "content", "status",
	}

	ci, err := services.CompanyIntroduction.ShowCompanyIntroduction(1, column)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxShowErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, gin.H{
		"id":      ci.ID,
		"name":    ci.Name,
		"image":   ci.Image,
		"content": ci.Content,
		"status":  ci.Status,
	})
}

func (ctr *companyIntroductionController) UpdateCompanyIntroduction(ctx *gin.Context) {
	var r = adminRequest.NewUpdateCompanyIntroduction()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, data)
		return
	}

	data := gin.H{
		"name":    r.Name,
		"image":   r.Image,
		"content": r.Content,
		"status":  *r.Status,
	}

	err := services.CompanyIntroduction.UpdateCompanyIntroduction(1, data)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

// Upload 上传
func (ctr *companyIntroductionController) Upload(ctx *gin.Context) {
	var r = adminRequest.NewUploadCompanyImage()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, data)
		return
	}

	upload := services.NewUploadService(config.App().PublicPath())

	path, err := upload.UploadFile("/web/company/", r.File)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.ReqUploadFileErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, gin.H{"path": "/public/" + path})
}
