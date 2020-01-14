package admin

import (
	"fagin/app"
	"fagin/app/errno"
	admin_request "fagin/app/requests/admin"
	"fagin/app/responses/admin_response"
	"fagin/app/service"
	"fagin/config"
	"fagin/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"time"
)

type websiteConfigController struct{}

var WebsiteConfigController websiteConfigController

func (websiteConfigController) Info(ctx *gin.Context) {
	column := []string{"web_name", "web_en_name", "web_title", "keywords", "description", "company_name",
		"contact_number", "company_address", "email", "icp", "public_security_record", "web_logo", "qr_code",
	}
	wc, err := service.WebsiteConfigService.ShowInfo(1, column)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.ErrWebsiteConfig, nil)
		return
	}

	data := admin_response.WebsiteConfig(*wc).Item()
	app.JsonResponse(ctx, errno.OK, data)
	return
}

func (websiteConfigController) UpdateInfo(ctx *gin.Context) {
	var r admin_request.UpdateWebsiteConfig
	if data, ok := r.Validate(ctx); !ok {
		app.JsonResponse(ctx, errno.ErrBind, data)
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
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.ErrUpdateWebsiteConfig, nil)
		return
	}

	app.JsonResponse(ctx, errno.OK, nil)
	return
}

// 上传
func (websiteConfigController) Upload(ctx *gin.Context) {
	var r admin_request.UploadWebsiteConfigPic
	if data, ok := r.Validate(ctx); !ok {
		app.JsonResponseWithStatus(ctx, http.StatusBadRequest, errno.ErrBind, data)
		return
	}

	filePath := "/web/website/" + time.Now().Format("2006123") + app.RandString(20) + path.Ext(r.File.Filename)

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
		app.JsonResponse(ctx, errno.ErrUploadFile, nil)
		return
	}

	app.JsonResponse(ctx, errno.OK, gin.H{"path": filePath})

	return
}