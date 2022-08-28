package responses

import (
	"fagin/app/models/website_config"
	"fagin/pkg/response"
)

type websiteConfig []website_config.WebsiteConfig

func NewWebsiteConfig(models ...website_config.WebsiteConfig) *response.Collect[websiteConfig] {
	return new(response.Collect[websiteConfig]).SetCollect(models)
}

func (res websiteConfig) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for i := range res {
		m := map[string]interface{}{
			"web_name":               res[i].WebName,
			"web_en_name":            res[i].WebEnName,
			"web_title":              res[i].WebTitle,
			"keywords":               res[i].Keywords,
			"description":            res[i].Description,
			"company_name":           res[i].CompanyName,
			"contact_number":         res[i].ContactNumber,
			"company_address":        res[i].CompanyAddress,
			"email":                  res[i].Email,
			"icp":                    res[i].ICP,
			"public_security_record": res[i].PublicSecurityRecord,
			"web_logo":               res[i].WebLogo,
			"qr_code":                res[i].QRCode,
		}
		sm = append(sm, m)
	}

	return sm
}
