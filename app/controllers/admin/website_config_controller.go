package admin

import (
	"encoding/json"
	"fagin/app/caches"
	"fagin/app/errno"
	adminRequest "fagin/app/requests/admin"
	response "fagin/app/responses/admin"
	"fagin/app/service"
	"fagin/config"
	"github.com/gin-gonic/gin"
)

type websiteConfigController struct {
	BaseController
}

var WebsiteConfigController websiteConfigController

func (wc *websiteConfigController) Info(ctx *gin.Context) {
	webConfig := caches.NewWebsiteConfig(func(key string) (b []byte, err error) {
		column := []string{
			"web_name", "web_en_name", "web_title", "keywords", "description", "company_name",
			"contact_number", "company_address", "email", "icp", "public_security_record",
			"web_logo", "qr_code",
		}
		info, err := service.WebsiteConfigService.ShowInfo(1, column)
		if err != nil {
			return nil, err
		}
		data := response.WebsiteConfig(*info).Item()
		return json.Marshal(data)
	})
	str, err := webConfig.Get("info")
	if err != nil {
		wc.ResponseJsonErrLog(ctx, errno.Serve.ShowErr, err, nil)
		return
	}

	var data gin.H
	err = json.Unmarshal([]byte(str), &data)
	if err != nil {
		wc.ResponseJsonErrLog(ctx, errno.Serve.ShowErr, err, nil)
		return
	}
	wc.ResponseJsonOK(ctx, data)
	return
}

func (wc *websiteConfigController) UpdateInfo(ctx *gin.Context) {
	var r = adminRequest.NewUpdateWebsiteConfig()
	if data, ok := r.Validate(ctx); !ok {
		wc.ResponseJsonErr(ctx, errno.Serve.BindErr, data)
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

	err := service.WebsiteConfigService.UpdateInfo(1, data)
	if err != nil {
		wc.ResponseJsonErrLog(ctx, errno.Serve.UpdateErr, err, nil)
		return
	}

	wc.ResponseJsonOK(ctx, nil)
	return
}

// Upload 上传
func (wc *websiteConfigController) Upload(ctx *gin.Context) {
	var r = adminRequest.NewUploadWebsiteConfigPic()
	if data, ok := r.Validate(ctx); !ok {
		wc.ResponseJsonErr(ctx, errno.Serve.BindErr, data)
		return
	}

	upload := service.NewUploadService(config.App.PublicPath)
	path, err := upload.UploadFile("/web/website/", r.File)
	if err != nil {
		wc.ResponseJsonErrLog(ctx, errno.Serve.UploadFileErr, err, nil)
		return
	}

	wc.ResponseJsonOK(ctx, gin.H{"path": "/public/" + path})
}
