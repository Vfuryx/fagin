package responses

import (
	"fagin/app/models/website_config"
	"fagin/pkg/response"
)

type websiteConfig struct {
	ms []*website_config.WebsiteConfig

	response.Collect
}

func NewWebsiteConfig(models ...*website_config.WebsiteConfig) response.Response {
	res := websiteConfig{ms: models}
	res.SetCollect(&res)

	return &res
}

func (res *websiteConfig) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for i := range res.ms {
		m := map[string]interface{}{
			"web_name":               res.ms[i].WebName,
			"web_en_name":            res.ms[i].WebEnName,
			"web_title":              res.ms[i].WebTitle,
			"keywords":               res.ms[i].Keywords,
			"description":            res.ms[i].Description,
			"company_name":           res.ms[i].CompanyName,
			"contact_number":         res.ms[i].ContactNumber,
			"company_address":        res.ms[i].CompanyAddress,
			"email":                  res.ms[i].Email,
			"icp":                    res.ms[i].ICP,
			"public_security_record": res.ms[i].PublicSecurityRecord,
			"web_logo":               res.ms[i].WebLogo,
			"qr_code":                res.ms[i].QRCode,
		}
		sm = append(sm, m)
	}

	return sm
}
