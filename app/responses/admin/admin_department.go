package responses

import (
	"fagin/app"
	"fagin/app/models/admin_department"
	"fagin/pkg/response"
)

type adminDepartment []admin_department.AdminDepartment

func NewAdminDepartment(models ...admin_department.AdminDepartment) *response.Collect[adminDepartment] {
	return new(response.Collect[adminDepartment]).SetCollect(models)

}

func (res adminDepartment) Serialize() []map[string]interface{} {
	return res.getTree(res, 0)
}

func (res adminDepartment) getTree(data []admin_department.AdminDepartment, pid uint) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, response.DefCap)

	for index := range data {
		if data[index].ParentID == pid {
			m := map[string]interface{}{
				"id":         data[index].ID,
				"parent_id":  data[index].ParentID,
				"name":       data[index].Name,
				"remark":     data[index].Remark,
				"sort":       data[index].Sort,
				"status":     data[index].Status,
				"created_at": app.TimeToStr(data[index].CreatedAt),
			}
			if children := res.getTree(data, data[index].ID); len(children) > 0 {
				m["children"] = children
			}

			result = append(result, m)
		}
	}

	return result
}
