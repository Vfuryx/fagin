package responses

import (
	"fagin/app/models/admin_menu"
	"fagin/pkg/response"
)

type adminRoleRoute struct {
	ms []*admin_menu.AdminMenu

	response.Collect
}

func NewAdminRoleRoute(models ...*admin_menu.AdminMenu) response.Response {
	res := adminRoleRoute{ms: models}
	res.SetCollect(&res)

	return &res
}

func (res *adminRoleRoute) Serialize() []map[string]interface{} {
	return res.RouteTree(res.ms, 0)
}

func (res adminRoleRoute) RouteTree(data []*admin_menu.AdminMenu, pid uint) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, response.DefCap)

	for index := range data {
		if data[index].ParentID == pid {
			m := map[string]interface{}{
				"id":       data[index].ID,
				"parentID": data[index].ParentID,
				"name":     data[index].Name,
				"path":     data[index].Path,
				"title":    data[index].Title,
			}
			if children := res.RouteTree(data, data[index].ID); len(children) > 0 {
				m["children"] = children
			}

			result = append(result, m)
		}
	}

	return result
}