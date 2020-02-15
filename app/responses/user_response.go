package responses

import (
	"fagin/app/models/user"
	"fagin/pkg/response"
)

type userResponse struct {
	Users []user.User
	response.Collect
}

var _ response.Response = &userResponse{}

func UserResponse(users ...user.User) *userResponse {
	ur := userResponse{Users: users}
	ur.Collect.Response = &ur
	return &ur
}

func (ur *userResponse) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, 20)
	for _, u := range ur.Users {
		m := map[string]interface{}{
			"id":       u.ID,
			"username": u.Username,
			"status":   u.Status,
		}
		sm = append(sm, m)
	}
	return sm
}
