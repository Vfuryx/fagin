package api_responses

import (
	"fagin/app/models/user"
	"fagin/pkg/response"
)

type userResponse struct {
	ms []*user.User

	response.Collect
}

func NewUserResponse(users ...*user.User) response.Response {
	ur := userResponse{ms: users}
	ur.Collect.Response = &ur

	return &ur
}

func (ur *userResponse) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for _, u := range ur.ms {
		m := map[string]interface{}{
			"id":       u.ID,
			"username": u.Username,
			"status":   u.Status,
		}
		sm = append(sm, m)
	}

	return sm
}
