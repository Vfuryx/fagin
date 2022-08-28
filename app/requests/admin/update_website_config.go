package request

import (
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type UpdateWebsiteConfig struct {
	WebName              string `form:"web_name" json:"web_name" binding:"required,max=32"`
	WebEnName            string `form:"web_en_name" json:"web_en_name" binding:"required,max=32"`
	WebTitle             string `form:"web_title" json:"web_title" binding:"required,max=32"`
	Keywords             string `form:"keywords" json:"keywords" binding:"required,max=127"`
	Description          string `form:"description" json:"description" binding:"required,max=255"`
	CompanyName          string `form:"company_name" json:"company_name" binding:"required,max=32"`
	ContactNumber        string `form:"contact_number" json:"contact_number" binding:"required,max=16"`
	CompanyAddress       string `form:"company_address" json:"company_address" binding:"required,max=32"`
	Email                string `form:"email" json:"email" binding:"required,email,max=32"`
	ICP                  string `form:"icp" json:"icp" binding:"required,max=32"`
	PublicSecurityRecord string `form:"public_security_record" json:"public_security_record" binding:"required,max=32"`
	WebLogo              string `form:"web_logo" json:"web_logo" binding:"required,max=255"`
	QRCode               string `form:"qr_code" json:"qr_code" binding:"required,max=255"`
}

var _ request.Request = UpdateWebsiteConfig{}

func (UpdateWebsiteConfig) Message() map[string]string {
	return map[string]string{
		// "WebName.required":              "网站名称不能为空",
		// "WebName.max":                   "网站名称不能大于32位",
		// "WebEnName.required":            "网站英文名称不能为空",
		// "WebEnName.max":                 "网站英文名称不能大于32位",
		// "WebTitle.required":             "网站标题不能为空",
		// "WebTitle.max":                  "网站标题不能大于32位",
		// "Keywords.required":             "关键词不能为空",
		// "Keywords.max":                  "关键词不能大于127位",
		// "Description.required":          "描述不能为空",
		// "Description.max":               "描述不能大于255位",
		// "CompanyName.required":          "公司名称不能为空",
		// "CompanyName.max":               "公司名称不能大于32位",
		// "ContactNumber.required":        "联系电话不能为空",
		// "ContactNumber.max":             "联系电话不能大于16位",
		// "CompanyAddress.required":       "公司地址不能为空",
		// "CompanyAddress.max":            "公司地址不能大于32位",
		// "Email.required":                "邮箱地址不能为空",
		// "Email.max":                     "邮箱地址不能大于32位",
		// "Email.email":                   "邮箱地址格式不正确",
		// "ICP.required":                  "ICP备案不能为空",
		// "ICP.max":                       "ICP备案不能大于32位",
		// "PublicSecurityRecord.required": "公安部备案不能为空",
		// "PublicSecurityRecord.max":      "公安部备案不能大于32位",
		// "WebLogo.required":              "网站logo不能为空",
		// "WebLogo.max":                   "网站logo不能大于255位",
		// "QRCode.required":               "二维码不能为空",
		// "QRCode.max":                    "二维码大于255位",
	}
}

func (UpdateWebsiteConfig) Attributes() map[string]string {
	return map[string]string{
		"WebName":              "网站名称",
		"WebEnName":            "网站英文名称",
		"WebTitle":             "网站标题",
		"Keywords":             "关键词",
		"Description":          "描述",
		"CompanyName":          "公司名称",
		"ContactNumber":        "联系电话",
		"CompanyAddress":       "公司地址",
		"Email":                "邮箱地址",
		"ICP":                  "ICP备案",
		"PublicSecurityRecord": "公安部备案",
		"WebLogo":              "网站logo",
		"QRCode":               "二维码",
	}
}
func (r UpdateWebsiteConfig) Validate(ctx *gin.Context) (any, map[string]string) {
	return request.Validate[UpdateWebsiteConfig](r, ctx)
}
