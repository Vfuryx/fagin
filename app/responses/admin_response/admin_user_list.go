package admin_response

import (
	"fagin/app/models/admin_user"
	"fagin/pkg/response"
)

type adminUserList struct {
	Ms []admin_user.AdminUser
	response.Collect
}

var _ response.Response = &adminUserList{}

func AdminUserList(models ...admin_user.AdminUser) *adminUserList {
	res := adminUserList{Ms: models}
	res.Collect.Response = &res
	return &res
}

func (res *adminUserList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, 20)
	for _, model := range res.Ms {
		m := map[string]interface{}{
			"id":        model.ID,
			"username":  model.Username,
			"nick_name": model.NickName,
			"email":     model.Email,
			"phone":     model.Phone,
			"avatar":    model.Avatar,
			"sex":       model.Sex,
			"remark":    model.Remark,
			"status":    model.Status,
			"role_id":   model.RoleID,
			"create_at": model.CreatedAt,
		}
		sm = append(sm, m)
	}
	return sm
}
