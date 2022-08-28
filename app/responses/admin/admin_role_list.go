package responses

import (
	"fagin/app"
	"fagin/app/models/admin_menu"
	"fagin/app/models/admin_role"
	"fagin/pkg/response"
)

type adminRoleList []admin_role.AdminRole

func NewAdminRoleList(models ...admin_role.AdminRole) *response.Collect[adminRoleList] {
	return new(response.Collect[adminRoleList]).SetCollect(models)
}

func (r adminRoleList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for i := range r {
		ids := r.getMenuIDs(r[i].Menus)
		m := map[string]interface{}{
			"id":         r[i].ID,
			"type":       r[i].Type,
			"name":       r[i].Name,
			"key":        r[i].Key,
			"remark":     r[i].Remark,
			"sort":       r[i].Sort,
			"status":     r[i].Status,
			"created_at": app.TimeToStr(r[i].CreatedAt),
			"menu_ids":   ids,
		}
		sm = append(sm, m)
	}

	return sm
}

func (r adminRoleList) getMenuIDs(menus []admin_menu.AdminMenu) []uint {
	ids := make([]uint, 0)

	for i := range menus {
		ids = append(ids, menus[i].ID)
	}

	return ids
}
