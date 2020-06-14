package admin_response

import (
	"fagin/app/models/admin_menu"
	"fagin/pkg/response"
	"time"
)

type adminMenusList struct {
	Ms []admin_menu.AdminMenu
	response.Collect
}

var _ response.Response = &adminMenusList{}

func AdminMenusList(models ...admin_menu.AdminMenu) *adminMenusList {
	res := adminMenusList{Ms: models}
	res.Collect.Response = &res
	return &res
}

func (res *adminMenusList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, 20)
	sm = getMenuTree(res.Ms)
	return sm
}

func getMenuTree(data []admin_menu.AdminMenu) []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, 20)
	for _, menu := range data {
		mc := getMenuTree(menu.Children)
		m := make(map[string]interface{})
		m["children"] = mc
		m["id"] = menu.ID
		m["parent_id"] = menu.ParentId
		m["paths"] = menu.Paths
		m["name"] = menu.Name
		m["title"] = menu.Title
		m["icon"] = menu.Icon
		m["type"] = menu.Type
		m["path"] = menu.Path
		m["component"] = menu.Component
		m["action"] = menu.Action
		m["permission"] = menu.Permission
		m["sort"] = menu.Sort
		m["visible"] = menu.Visible
		m["is_link"] = menu.IsLink
		m["created_at"] = menu.CreatedAt.Format(time.RFC3339)
		sm = append(sm, m)
	}
	return sm
}
