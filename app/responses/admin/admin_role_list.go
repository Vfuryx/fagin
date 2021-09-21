package admin_responses

import (
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

func (res *adminRoleList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, 20)
	for _, model := range res.Ms {
		ids := getMenuIDs(model.Menus)
		m := map[string]interface{}{
			"id":         model.ID,
			"type":       model.Type,
			"name":       model.Name,
			"key":        model.Key,
			"sort":       model.Sort,
			"status":     model.Status,
			"created_at": model.CreatedAt,
			"menu_ids":   ids,
		}
		sm = append(sm, m)
	}
	return sm
}

func getMenuIDs(menus []admin_menu.AdminMenu) []uint {
	ids := make([]uint, 0, 20)
	for _, menu := range menus {
		ids = append(ids, menu.ID)
	}
	return ids
}
