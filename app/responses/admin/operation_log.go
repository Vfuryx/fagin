package admin_responses

import (
	"fagin/app/constants/time_format"
	aol "fagin/app/models/admin_operation_log"
	"fagin/pkg/response"
)

type operationLog struct {
	Ms []aol.AdminOperationLog
	response.Collect
}

var _ response.IResponse = &operationLog{}

func OperationLog(models ...aol.AdminOperationLog) *operationLog {
	res := operationLog{Ms: models}
	res.Collect.IResponse = &res
	return &res
}

func (res *operationLog) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, 20)
	for _, model := range res.Ms {
		m := map[string]interface{}{
			"id":         model.ID,
			"user":       model.User,
			"method":     model.Method,
			"path":       model.Path,
			"ip":         model.IP,
			"operation":  model.Operation,
			"module":     model.Module,
			"created_at": model.CreatedAt.Format(time_format.Def),
		}
		sm = append(sm, m)
	}
	return sm
}
