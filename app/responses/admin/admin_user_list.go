package admin_responses

import (
	"fagin/app"
	"fagin/app/models/admin_user"
	"fagin/pkg/response"
	"github.com/gin-gonic/gin"
)

type adminUserList struct {
	Ms []admin_user.AdminUser
	response.Collect
}

var _ response.IResponse = &adminUserList{}

func AdminUserList(models ...admin_user.AdminUser) *adminUserList {
	res := adminUserList{Ms: models}
	res.Collect.IResponse = &res
	return &res
}

func (res *adminUserList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, 20)
	for _, model := range res.Ms {
		roles := make([]gin.H, 0, 5)
		for _, r := range model.Roles {
			roles = append(roles, gin.H{"id": r.ID, "name": r.Name})
		}
		m := map[string]interface{}{
			"id":        model.ID,
			"username":  model.Username,
			"nick_name": model.NickName,
			"phone":     model.Phone,
			"roles":     roles,
			"create_at": model.CreatedAt.Format(app.TimeFormat),
		}
		sm = append(sm, m)
	}
	return sm
}
