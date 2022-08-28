package responses

import (
	"fagin/app"
	all "fagin/app/models/admin_login_log"
	"fagin/pkg/response"
)

type loginLog []all.AdminLoginLog

func NewLoginLog(models ...all.AdminLoginLog) *response.Collect[loginLog] {
	return new(response.Collect[loginLog]).SetCollect(models)
}

func (res loginLog) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for i := range res {
		m := map[string]interface{}{
			"id":         res[i].ID,
			"login_name": res[i].LoginName,
			"ip":         res[i].IP,
			"location":   res[i].Location,
			"browser":    res[i].Browser,
			"os":         res[i].OS,
			"created_at": app.TimeToStr(res[i].CreatedAt),
		}
		sm = append(sm, m)
	}

	return sm
}
