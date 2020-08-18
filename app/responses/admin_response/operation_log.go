package admin_response

import (
	aol "fagin/app/models/admin_operation_log"
	"fagin/pkg/response"
)

type operationLog struct {
	Ms []aol.AdminOperationLog
	response.Collect
}

var _ response.Response = &operationLog{}

func OperationLog(models ...aol.AdminOperationLog) *operationLog {
	res := operationLog{Ms: models}
	res.Collect.Response = &res
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
			"created_at": model.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		sm = append(sm, m)
	}
	return sm
}
