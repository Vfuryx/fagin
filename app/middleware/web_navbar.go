package middleware

import (
	"encoding/json"
	"fagin/app"
	"fagin/app/caches"
	"fagin/app/errno"
	"fagin/app/service"
	"fagin/pkg/response"
	"github.com/gin-gonic/gin"
	"runtime/debug"
)

type webNavbar struct{}

var WebNavbar webNavbar

func (*webNavbar) WebNavbar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		navbar := caches.NewHomeNavbar(func() ([]byte, error) {
			c, err := service.Category.All(gin.H{"status": 1, "orderBy": "sort desc"}, []string{"id", "name"}, nil)
			if err != nil {
				return nil, err
			}
			return json.Marshal(c)
		})
		str, err := navbar.Get()
		if err != nil {
			go app.Log().Println(errno.MidErr, err, string(debug.Stack()))
			response.JsonErr(ctx, errno.MidErr, nil)
			ctx.Abort()
			return
		}
		ctx.Set("web_cate", str)
		ctx.Next()
	}
}
