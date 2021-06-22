package admin_responses

import (
	p "fagin/app/models/admin_permission"
	group "fagin/app/models/admin_permission_group"
	"fagin/pkg/response"
	"github.com/gin-gonic/gin"
)

type groupPermissionList struct {
	ms []group.AdminPermissionGroup
	response.Collect
}

var _ response.IResponse = &groupPermissionList{}

func GroupPermissionList(models ...group.AdminPermissionGroup) *groupPermissionList {
	res := groupPermissionList{ms: models}
	res.Collect.IResponse = &res
	return &res
}

func (res *groupPermissionList) Serialize() []map[string]interface{} {
	var permissions []p.AdminPermission
	params := gin.H{}
	sm := make([]map[string]interface{}, 0, 20)
	for _, model := range res.ms {
		params["gid"] = model.ID
		_ = p.Dao().Query(params, []string{"id", "name", "method"}, nil).Find(&permissions)
		ps := make([]map[string]interface{}, 0, 20)
		item := gin.H{}
		for _, permission := range permissions {
			item = gin.H{
				"id": permission.ID,
				"name": permission.Name,
				"method": permission.Method,
			}
			ps = append(ps, item)
		}
		m := map[string]interface{}{
			"id":          model.ID,
			"name":        model.Name,
			"sort":        model.Sort,
			"permissions": ps,
		}
		sm = append(sm, m)
	}
	return sm
}
