package admin_responses

import (
	"fagin/app"
	"fagin/app/models/admin_menu"
	"fagin/pkg/response"
)

type adminMenusList struct {
	ms []admin_menu.AdminMenu
	response.Collect
}

var _ response.IResponse = &adminMenusList{}

func AdminMenusList(models ...admin_menu.AdminMenu) *adminMenusList {
	res := adminMenusList{ms: models}
	res.SetCollect(&res)
	return &res
}

func (res *adminMenusList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, 20)
	sm = getMenuTree(res.ms)
	return sm
}

func getMenuTree(data []admin_menu.AdminMenu) []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, 20)
	for _, menu := range data {
		mc := getMenuTree(menu.Children)
		m := make(map[string]interface{})
		m["children"] = mc
		m["id"] = menu.ID
		m["parent_id"] = menu.ParentID
		m["paths"] = menu.Paths
		m["name"] = menu.Name
		m["title"] = menu.Title
		m["icon"] = menu.Icon
		m["type"] = menu.Type
		m["path"] = menu.Path
		m["sort"] = menu.Sort
		m["is_show"] = menu.IsShow
		m["component"] = menu.Component
		m["redirect"] = menu.Redirect
		m["target"] = menu.Target
		m["status"] = menu.Status
		m["is_hide_child"] = menu.IsHideChild
		m["created_at"] = app.TimeToStr(menu.CreatedAt)
		sm = append(sm, m)
	}
	return sm
}
