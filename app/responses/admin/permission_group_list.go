package admin_responses

import (
	"fagin/app"
	"fagin/app/models/admin_permission_group"
	"fagin/pkg/response"
)

type permissionGroupList struct {
	ms []admin_permission_group.AdminPermissionGroup
	response.Collect
}

var _ response.IResponse = &permissionGroupList{}

func PermissionGroupList(models ...admin_permission_group.AdminPermissionGroup) *permissionGroupList {
	res := permissionGroupList{ms: models}
	res.SetCollect(&res)
	return &res
}

func (res *permissionGroupList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, 20)
	for _, model := range res.ms {
		m := map[string]interface{}{
			"id":         model.ID,
			"name":       model.Name,
			"type":       model.Type,
			"sort":       model.Sort,
			"created_at": app.TimeToStr(model.CreatedAt),
		}
		sm = append(sm, m)
	}
	return sm
}
