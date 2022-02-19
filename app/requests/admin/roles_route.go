package request

import (
	"fagin/pkg/request"
)

type RolesRoute struct {
	Roles []uint `form:"roles[]" json:"roles" binding:"required,dive,required"`

	request.Validation `binding:"-"`
}

func NewRolesRoute() *RolesRoute {
	r := new(RolesRoute)
	r.SetRequest(r)

	return r
}

func (*RolesRoute) Message() map[string]string {
	return map[string]string{}
}

func (*RolesRoute) Attributes() map[string]string {
	return map[string]string{
		"Roles": "角色组",
	}
}

// func (r *RolesRoute) Validate(ctx *gin.Context) (map[string]string, bool) {
//	return request.Validated(r, ctx)
//}
