package responses

import (
	"fagin/app/models/admin_role"
	"fagin/pkg/response"
)

type adminSelectRoleList struct {
	ms []*admin_role.AdminRole

	response.Collect
}

func NewAdminSelectRoleList(models ...*admin_role.AdminRole) response.Response {
	res := adminSelectRoleList{ms: models}
	res.SetCollect(&res)

	return &res
}

func (res *adminSelectRoleList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for i := range res.ms {
		m := map[string]interface{}{
			"id":   res.ms[i].ID,
			"name": res.ms[i].Name,
		}
		sm = append(sm, m)
	}

	return sm
}
