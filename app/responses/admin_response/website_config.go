package admin_response

import (
	"fagin/app/models/website_config"
	"fagin/pkg/response"
)

type websiteConfig struct {
	Ms []website_config.WebsiteConfig
	response.Collect
}

var _ response.Response = &websiteConfig{}

func WebsiteConfig(models ...website_config.WebsiteConfig) *websiteConfig {
	res := websiteConfig{Ms: models}
	res.Collect.Response = &res
	return &res
}

func (res *websiteConfig) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, 20)
	for _, model := range res.Ms {
		m := map[string]interface{}{
			"web_name":               model.WebName,
			"web_en_name":            model.WebEnName,
			"web_title":              model.WebTitle,
			"keywords":               model.Keywords,
			"description":            model.Description,
			"company_name":           model.CompanyName,
			"contact_number":         model.ContactNumber,
			"company_address":        model.CompanyAddress,
			"email":                  model.Email,
			"icp":                    model.ICP,
			"public_security_record": model.PublicSecurityRecord,
			"web_logo":               model.WebLogo,
			"qr_code":                model.QRCode,
		}
		sm = append(sm, m)
	}
	return sm
}
