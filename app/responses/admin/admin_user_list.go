package responses

import (
	"fagin/app"
	"fagin/app/models/admin_user"
	"fagin/pkg/response"

	"github.com/gin-gonic/gin"
)

type adminUserList struct {
	ms []*admin_user.AdminUser

	response.Collect
}

func NewAdminUserList(models ...*admin_user.AdminUser) response.Response {
	res := adminUserList{ms: models}
	res.SetCollect(&res)

	return &res
}

func (res *adminUserList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for i := range res.ms {
		roles := make([]gin.H, 0)
		for _, r := range res.ms[i].Roles {
			roles = append(roles, gin.H{"id": r.ID, "name": r.Name})
		}

		m := map[string]interface{}{
			"id":        res.ms[i].ID,
			"username":  res.ms[i].Username,
			"nick_name": res.ms[i].NickName,
			"email":     res.ms[i].Email,
			"remark":    res.ms[i].Remark,
			"phone":     res.ms[i].Phone,
			"roles":     roles,
			"status":    res.ms[i].Status,
			"create_at": app.TimeToStr(res.ms[i].CreatedAt),
		}
		sm = append(sm, m)
	}

	return sm
}
