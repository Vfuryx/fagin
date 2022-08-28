package responses

import (
	"fagin/app/models/admin_role"
	"fagin/pkg/response"
)

type adminSelectRoleList []admin_role.AdminRole

func NewAdminSelectRoleList(models ...admin_role.AdminRole) *response.Collect[adminSelectRoleList] {
	return new(response.Collect[adminSelectRoleList]).SetCollect(models)
}

func (res adminSelectRoleList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for i := range res {
		m := map[string]interface{}{
			"id":   res[i].ID,
			"name": res[i].Name,
		}
		sm = append(sm, m)
	}

	return sm
}
