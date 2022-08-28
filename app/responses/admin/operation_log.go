package responses

import (
	"fagin/app"
	aol "fagin/app/models/admin_operation_log"
	"fagin/pkg/response"
)

type operationLog []aol.AdminOperationLog

func NewOperationLog(models ...aol.AdminOperationLog) *response.Collect[operationLog] {
	return new(response.Collect[operationLog]).SetCollect(models)
}

func (res operationLog) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for i := range res {
		m := map[string]interface{}{
			"id":         res[i].ID,
			"user":       res[i].User,
			"method":     res[i].Method,
			"path":       res[i].Path,
			"ip":         res[i].IP,
			"operation":  res[i].Operation,
			"module":     res[i].Module,
			"created_at": app.TimeToStr(res[i].CreatedAt),
		}
		sm = append(sm, m)
	}

	return sm
}
