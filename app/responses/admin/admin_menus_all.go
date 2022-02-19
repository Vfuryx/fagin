package responses

import (
	"fagin/app/models/admin_menu"
	"fagin/pkg/response"
)

type adminMenusAll struct {
	ms []*admin_menu.AdminMenu

	response.Collect
}

func NewAdminMenusAll(models ...*admin_menu.AdminMenu) response.Response {
	res := adminMenusAll{ms: models}
	res.SetCollect(&res)

	return &res
}

func (res *adminMenusAll) Serialize() []map[string]interface{} {
	return res.getMenuTree(res.ms, 0)
}

func (res *adminMenusAll) getMenuTree(data []*admin_menu.AdminMenu, pid uint) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, response.DefCap)

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
