package auths

import (
	"fagin/app/models/admin_user"

	"github.com/gin-gonic/gin"
)

// GetAdminID 获取后台用户ID
func GetAdminID(ctx *gin.Context) uint {
	return ctx.GetUint(admin_user.AdminUserIDKey)
}

// GetAdmin 获取后台用户
func GetAdmin(ctx *gin.Context) (*admin_user.AdminUser, error) {
	uid := GetAdminID(ctx)
	admin := admin_user.New()

	err := admin_user.NewDao().FindByID(uid, []string{"*"})
	if err != nil {
		return nil, err
	}

	return admin, nil
}
