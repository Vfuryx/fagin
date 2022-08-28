package api_responses

import (
	"fagin/app/models/user"
	"fagin/pkg/response"
)

type userResponse []user.User

func NewUserResponse(models ...user.User) *response.Collect[userResponse] {
	return new(response.Collect[userResponse]).SetCollect(models)
}

func (res userResponse) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for _, u := range res {
		m := map[string]interface{}{
			"id":       u.ID,
			"username": u.Username,
			"status":   u.Status,
		}
		sm = append(sm, m)
	}

	return sm
}
