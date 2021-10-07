package admin_responses

import (
	"fagin/app"
	"fagin/app/models/admin_menu"
	"fagin/app/models/admin_role"
	"fagin/pkg/response"
)

type adminRoleList struct {
	Ms []admin_role.AdminRole
	response.Collect
}

var _ response.IResponse = &adminRoleList{}

func AdminRoleList(models ...admin_role.AdminRole) *adminRoleList {
	res := adminRoleList{Ms: models}
	res.SetCollect(&res)
	return &res
}

func (r *adminRoleList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, 20)
	for _, model := range r.Ms {
		ids := r.getMenuIDs(model.Menus)
		m := map[string]interface{}{
			"id":         model.ID,
			"type":       model.Type,
			"name":       model.Name,
			"key":        model.Key,
			"remark":     model.Remark,
			"sort":       model.Sort,
			"status":     model.Status,
			"created_at": app.TimeToStr(model.CreatedAt),
			"menu_ids":   ids,
		}
		sm = append(sm, m)
	}
	return sm
}

func (r adminRoleList) getMenuIDs(menus []admin_menu.AdminMenu) []uint {
	ids := make([]uint, 0, 20)
	for _, menu := range menus {
		ids = append(ids, menu.ID)
	}
	return ids
}
