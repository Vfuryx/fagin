package auths

import (
	"fagin/app/models/admin_user"
	"github.com/gin-gonic/gin"
)

func GetAdminID(ctx *gin.Context) uint64 {
	return ctx.GetUint64(admin_user.AdminUserIdKey)
}

func GetAdmin(ctx *gin.Context) (*admin_user.AdminUser, error) {
	uid := ctx.GetUint64(admin_user.AdminUserIdKey)
	admin := admin_user.New()
	err := admin_user.Dao().FindById(uint(uid), []string{"*"})
	if err != nil {
		return nil, err
	}
	return admin, nil
}