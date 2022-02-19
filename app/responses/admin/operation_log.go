package responses

import (
	"fagin/app"
	aol "fagin/app/models/admin_operation_log"
	"fagin/pkg/response"
)

type operationLog struct {
	ms []*aol.AdminOperationLog

	response.Collect
}

func NewOperationLog(models ...*aol.AdminOperationLog) response.Response {
	res := operationLog{ms: models}
	res.SetCollect(&res)

	return &res
}

func (res *operationLog) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for i := range res.ms {
		m := map[string]interface{}{
			"id":         res.ms[i].ID,
			"user":       res.ms[i].User,
			"method":     res.ms[i].Method,
			"path":       res.ms[i].Path,
			"ip":         res.ms[i].IP,
			"operation":  res.ms[i].Operation,
			"module":     res.ms[i].Module,
			"created_at": app.TimeToStr(res.ms[i].CreatedAt),
		}
		sm = append(sm, m)
	}

	return sm
}
