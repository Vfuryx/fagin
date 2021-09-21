package admin_responses

import (
	"fagin/app"
	"fagin/app/models/admin_permission"
	"fagin/pkg/response"
)

type permissionList struct {
	ms []admin_permission.AdminPermission
	response.Collect
}

var _ response.IResponse = &permissionList{}

func PermissionList(models ...admin_permission.AdminPermission) *permissionList {
	res := permissionList{ms: models}
	res.SetCollect(&res)
	return &res
}

func (res *permissionList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, 20)
	for _, model := range res.ms {
		m := map[string]interface{}{
			"id":         model.ID,
			"name":       model.Name,
			"gid":        model.GID,
			"paths":      model.Path,
			"method":     model.Method,
			"sort":       model.Sort,
			"status":     model.Status,
			"group":      model.Group.Name,
			"created_at": app.TimeToStr(model.CreatedAt),
		}
		sm = append(sm, m)
	}
	return sm
}
