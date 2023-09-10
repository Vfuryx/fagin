package role_transform

import (
	"fadmin/app/models/admin/domain/entities/role"
	"fadmin/pkg/response"
)

type adminRoleList []role.AdminRole

func NewAdminRoleList(models ...role.AdminRole) *response.Collect[adminRoleList] {
	return new(response.Collect[adminRoleList]).SetCollect(models)
}

func (r adminRoleList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for i := range r {
		m := map[string]interface{}{
			"id":         r[i].ID,
			"type":       r[i].Type,
			"name":       r[i].Name,
			"key":        r[i].Key,
			"remark":     r[i].Remark,
			"sort":       r[i].Sort,
			"status":     r[i].Status,
			"created_at": r[i].CreatedAt,
		}
		sm = append(sm, m)
	}

	return sm
}
