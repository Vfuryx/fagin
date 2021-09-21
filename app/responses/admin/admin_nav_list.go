package admin_responses

import (
	"fagin/app/models/admin_menu"
	"fagin/pkg/response"
)

type adminNavList struct {
	ms []admin_menu.AdminMenu
	response.Collect
}

var _ response.IResponse = &adminNavList{}

func AdminNavList(models ...admin_menu.AdminMenu) *adminNavList {
	res := adminNavList{ms: models}
	res.SetCollect(&res)
	return &res
}

func (res *adminNavList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, 20)
	for _, model := range res.ms {
		// 'id': 1,
		// 'name': 'dashboard',
		// 'parentId': 0,
		// 'path': '/dashboard/analysis',
		// 'component': 'RouteView',
		// 'redirect': '/dashboard/workplace'
		// 'meta': {
		//   'icon': 'dashboard',
		//   'title': '仪表盘',
		//   'target': '_blank',
		//   'show': true
		// },
		m := map[string]interface{}{
			"id":        model.ID,
			"name":      model.Name,
			"parentId":  model.ParentID,
			"path":      model.Path,
			"component": model.Component,
			"redirect":  model.Redirect,
			"sort":      model.Sort,
			"status":    model.Status,
			"meta": map[string]interface{}{
				"icon":         model.Icon,
				"title":        model.Title,
				"target":       model.Target,
				"show":         model.IsShow,
				"hideChildren": model.IsHideChild,
			},
		}
		sm = append(sm, m)
	}
	return sm
}
