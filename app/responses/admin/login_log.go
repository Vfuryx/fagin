package admin_responses

import (
	"fagin/app"
	all "fagin/app/models/admin_login_log"
	"fagin/pkg/response"
)

type loginLog struct {
	Ms []all.AdminLoginLog
	response.Collect
}

var _ response.Response = &loginLog{}

func LoginLog(models ...all.AdminLoginLog) *loginLog {
	res := loginLog{Ms: models}
	res.SetCollect(&res)
	return &res
}

func (res *loginLog) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, 20)
	for _, model := range res.Ms {
		m := map[string]interface{}{
			"id":         model.ID,
			"login_name": model.LoginName,
			"ip":         model.IP,
			"location":   model.Location,
			"browser":    model.Browser,
			"os":         model.OS,
			"created_at": app.TimeToStr(model.CreatedAt),
		}
		sm = append(sm, m)
	}
	return sm
}
