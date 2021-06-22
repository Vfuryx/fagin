package admin_responses

import (
	"fagin/app/models/admin_role"
	"fagin/pkg/response"
)

type adminSelectRoleList struct {
	Ms []admin_role.AdminRole
	response.Collect
}

var _ response.IResponse = &adminSelectRoleList{}

func AdminSelectRoleList(models ...admin_role.AdminRole) *adminSelectRoleList {
	res := adminSelectRoleList{Ms: models}
	res.Collect.IResponse = &res
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
