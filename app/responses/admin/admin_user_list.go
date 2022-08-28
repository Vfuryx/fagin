package responses

import (
	"fagin/app"
	"fagin/app/models/admin_user"
	"fagin/pkg/response"

	"github.com/gin-gonic/gin"
)

type adminUserList []admin_user.AdminUser

func NewAdminUserList(models ...admin_user.AdminUser) *response.Collect[adminUserList] {
	return new(response.Collect[adminUserList]).SetCollect(models)
}

func (res adminUserList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for i := range res {
		roles := make([]gin.H, 0)
		for _, r := range res[i].Roles {
			roles = append(roles, gin.H{"id": r.ID, "name": r.Name})
		}

		m := map[string]interface{}{
			"id":        res[i].ID,
			"username":  res[i].Username,
			"nick_name": res[i].NickName,
			"email":     res[i].Email,
			"remark":    res[i].Remark,
			"phone":     res[i].Phone,
			"roles":     roles,
			"status":    res[i].Status,
			"create_at": app.TimeToStr(res[i].CreatedAt),
		}
		sm = append(sm, m)
	}

	return sm
}
