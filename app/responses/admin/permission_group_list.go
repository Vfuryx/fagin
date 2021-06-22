package admin_responses

import (
	"fagin/app/constants/time_format"
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
	res.Collect.IResponse = &res
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
			"created_at": model.CreatedAt.Format(time_format.Def),
		}
		sm = append(sm, m)
	}
	return sm
}
