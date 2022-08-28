package admin

import (
	"encoding/json"
	"fagin/app/caches"
	"fagin/app/errno"
	adminRequest "fagin/app/requests/admin"
	response "fagin/app/responses/admin"
	"fagin/app/services"
	"fagin/config"
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type websiteConfigController struct {
	BaseController
}

var WebsiteConfigController websiteConfigController

func (ctr *websiteConfigController) Info(ctx *gin.Context) {
	webConfig := caches.NewWebsiteConfig(func() (b []byte, err error) {
		column := []string{
			"web_name", "web_en_name", "web_title", "keywords", "description", "company_name",
			"contact_number", "company_address", "email", "icp", "public_security_record",
			"web_logo", "qr_code",
		}
		info, err := services.WebsiteConfigService.ShowInfo(1, column)
		if err != nil {
			return nil, err
		}
		data := response.NewWebsiteConfig(*info).Item()
		return json.Marshal(data)
	})

	str, err := webConfig.Get()
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxShowErr, err)
		return
	}

	var data gin.H

	err = json.Unmarshal(str, &data)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxShowErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, data)
}

func (ctr *websiteConfigController) UpdateInfo(ctx *gin.Context) {
	r, msg := request.Validation[adminRequest.UpdateWebsiteConfig](ctx)
	if len(msg) > 0 {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, msg)
		return
	}

	data := map[string]interface{}{
		"web_name":               r.WebName,
		"web_en_name":            r.WebEnName,
		"web_title":              r.WebTitle,
		"keywords":               r.Keywords,
		"description":            r.Description,
		"company_name":           r.CompanyName,
		"contact_number":         r.ContactNumber,
		"company_address":        r.CompanyAddress,
		"email":                  r.Email,
		"icp":                    r.ICP,
		"public_security_record": r.PublicSecurityRecord,
		"web_logo":               r.WebLogo,
		"qr_code":                r.QRCode,
	}

	err := services.WebsiteConfigService.UpdateInfo(1, data)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

// Upload 上传
func (ctr *websiteConfigController) Upload(ctx *gin.Context) {
	r, msg := request.Validation[adminRequest.UploadWebsiteConfigPic](ctx)
	if len(msg) > 0 {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, msg)
		return
	}

	upload := services.NewUploadService(config.App().PublicPath())

	path, err := upload.UploadFile("/web/website/", r.File)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.ReqUploadFileErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, gin.H{"path": "/public/" + path})
}
