package responses

import (
	"fagin/app"
	"fagin/app/models/admin_menu"
	"fagin/app/models/admin_role"
	"fagin/pkg/response"
)

type adminRoleList struct {
	ms []*admin_role.AdminRole

	response.Collect
}

func NewAdminRoleList(models ...*admin_role.AdminRole) response.Response {
	res := adminRoleList{ms: models}
	res.SetCollect(&res)

	return &res
}

func (r *adminRoleList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for i := range r.ms {
		ids := r.getMenuIDs(r.ms[i].Menus)
		m := map[string]interface{}{
			"id":         r.ms[i].ID,
			"type":       r.ms[i].Type,
			"name":       r.ms[i].Name,
			"key":        r.ms[i].Key,
			"remark":     r.ms[i].Remark,
			"sort":       r.ms[i].Sort,
			"status":     r.ms[i].Status,
			"created_at": app.TimeToStr(r.ms[i].CreatedAt),
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
