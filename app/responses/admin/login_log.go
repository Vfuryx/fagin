package responses

import (
	"fagin/app"
	all "fagin/app/models/admin_login_log"
	"fagin/pkg/response"
)

type loginLog struct {
	ms []*all.AdminLoginLog

	response.Collect
}

func NewLoginLog(models ...*all.AdminLoginLog) response.Response {
	res := loginLog{ms: models}
	res.SetCollect(&res)

	return &res
}

func (res *loginLog) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for i := range res.ms {
		m := map[string]interface{}{
			"id":         res.ms[i].ID,
			"login_name": res.ms[i].LoginName,
			"ip":         res.ms[i].IP,
			"location":   res.ms[i].Location,
			"browser":    res.ms[i].Browser,
			"os":         res.ms[i].OS,
			"created_at": app.TimeToStr(res.ms[i].CreatedAt),
		}
		sm = append(sm, m)
	}

	return sm
}
