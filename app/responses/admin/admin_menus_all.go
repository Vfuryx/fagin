package admin_responses

import (
	"fagin/app/models/admin_menu"
	"fagin/pkg/response"
)

type adminMenusAll struct {
	ms []admin_menu.AdminMenu
	response.Collect
}

var _ response.IResponse = &adminMenusAll{}

func AdminMenusAll(models ...admin_menu.AdminMenu) *adminMenusAll {
	res := adminMenusAll{ms: models}
	res.SetCollect(&res)
	return &res
}

func (res *adminMenusAll) Serialize() []map[string]interface{} {
	return res.getMenuTree(res.ms, 0)
}

func (res *adminMenusAll) getMenuTree(data []admin_menu.AdminMenu, pid uint) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, 10)
	for index := range data {
		if data[index].ParentID == pid {
			m := map[string]interface{}{
				"id":        data[index].ID,
				"parent_id": data[index].ParentID,
				"icon":      data[index].Icon,
				"title":     data[index].Title,
			}
			if children := res.getMenuTree(data, data[index].ID); len(children) > 0 {
				m["children"] = children
			}
			result = append(result, m)
		}
	}
	return result
}
