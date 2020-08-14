package admin_response

import (
	"fagin/app/models/admin_role"
	"fagin/pkg/response"
)

type adminSelectRoleList struct {
	Ms []admin_role.AdminRole
	response.Collect
}

var _ response.Response = &adminSelectRoleList{}

func AdminSelectRoleList(models ...admin_role.AdminRole) *adminSelectRoleList {
	res := adminSelectRoleList{Ms: models}
	res.Collect.Response = &res
	return &res
}

func (res *adminSelectRoleList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, 20)
	for _, model := range res.Ms {
		m := map[string]interface{}{
			"id":   model.ID,
			"name": model.Name,
		}
		sm = append(sm, m)
	}
	return sm
}
