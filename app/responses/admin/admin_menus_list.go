package responses

import (
	"fagin/app"
	"fagin/app/models/admin_menu"
	"fagin/pkg/response"
)

type adminMenusList []admin_menu.AdminMenu

func NewAdminMenusList(models ...admin_menu.AdminMenu) *response.Collect[adminMenusList] {
	return new(response.Collect[adminMenusList]).SetCollect(models)
}

func (res adminMenusList) Serialize() []map[string]interface{} {
	return res.getMenuTree(res, 0)
}

func (res adminMenusList) getMenuTree(data []admin_menu.AdminMenu, pid uint) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, response.DefCap)

	for index := range data {
		if data[index].ParentID == pid {
			m := map[string]interface{}{
				"id":         data[index].ID,
				"parent_id":  data[index].ParentID,
				"icon":       data[index].Icon,
				"title":      data[index].Title,
				"permission": data[index].Permission,
				"path":       data[index].Path,
				"component":  data[index].Component,
				"sort":       data[index].Sort,
				"status":     data[index].Status,
				"created_at": app.TimeToStr(data[index].CreatedAt),
			}
			if children := res.getMenuTree(data, data[index].ID); len(children) > 0 {
				m["children"] = children
			}

			result = append(result, m)
		}
	}

	return result
}
