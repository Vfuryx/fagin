package middleware

import (
	"encoding/json"
	"fagin/app/caches"
	"fagin/app/errno"
	"fagin/app/models/admin_user"
	"fagin/app/service/admin_auth"
	"fagin/pkg/auths"
	"fagin/pkg/casbins"
	"fagin/pkg/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type adminAuth struct{}

var AdminAuth adminAuth

// IsLogin 验证后台管理员是否登录
func (*adminAuth) IsLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, err := admin_auth.AdminAuth.ParseRequest(ctx)
		if err != nil {
			response.JsonErr(ctx, err, nil)
			ctx.Abort()
			return
		}

		ctx.Set(admin_user.AdminUserNameKey, c.Name)
		ctx.Set(admin_user.AdminUserIdKey, c.UserID)

		ctx.Next()
	}
}

// AuthCheckRole 权限检查中间件
func (*adminAuth) AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := auths.GetAdminID(c)

		roles, err := casbins.Casbin.GetRolesForUser(strconv.FormatUint(userID, 10))
		if err != nil {
			response.JsonErr(c, errno.MidAuthCheckRoleErr, nil)
			c.Abort()
			return
		}

		isAdmin := false
		for _, r := range roles {
			if r == "admin" {
				isAdmin = true
			}
		}
		// 是超级管理员
		if isAdmin {
			c.Next()
		} else {
			ok, err := casbins.Casbin.CheckRoles(roles, c.FullPath(), c.Request.Method)
			if ok && err == nil {
				c.Next()
			} else {
				response.JsonErr(c, errno.MidAuthCheckRoleErr, nil)
				c.Abort()
				return
			}
		}

	}
}

// AuthCheckRoleCache 权限检查中间件 有缓存
func (*adminAuth) AuthCheckRoleCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := auths.GetAdminID(c)
		strUID := strconv.FormatUint(uid, 10)

		roleCache := caches.NewAdminCasbin(func() ([]byte, error) {
			// 缓存用户的角色
			roles, err := casbins.Casbin.GetRolesForUser(strUID)
			if err != nil {
				return nil, err
			}
			return json.Marshal(roles)
		})
		str, err := roleCache.Get(strUID)
		if err != nil {
			response.JsonErr(c, errno.MidAuthCheckRoleErr, nil)
			c.Abort()
			return
		}

		var roles []string
		err = json.Unmarshal(str, &roles)
		if err != nil {
			response.JsonErr(c, errno.MidAuthCheckRoleErr, nil)
			c.Abort()
			return
		}

		isAdmin := false
		for _, r := range roles {
			if r == "admin" {
				isAdmin = true
			}
		}

		//是超级管理员
		if isAdmin {
			c.Next()
		} else {
			fullPath := c.FullPath()
			method := c.Request.Method
			rbacCache := caches.NewAdminRBAC(func() ([]byte, error) {
				ok, err := casbins.Casbin.CheckRoles(roles, fullPath, method)
				if err != nil {
					return nil, err
				}
				if ok {
					return []byte{'1'}, nil
				}
				return []byte{'0'}, nil
			})
			// 缓存用户的授权
			str, err = rbacCache.Get(strUID, method, fullPath)
			if string(str) == "1" && err == nil {
				c.Next()
			} else {
				response.JsonErr(c, errno.MidAuthCheckRoleErr, nil)
				c.Abort()
				return
			}
		}
	}
}
