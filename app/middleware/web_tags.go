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

type webTags struct{}

var WebTags webTags

func (*webTags) WebTags() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		homeTags := caches.NewHomeTags(func() ([]byte, error) {
			params := gin.H{
				"status":  1,
				"orderBy": "id asc",
			}
			columns := []string{"id", "name", "status"}
			with := gin.H{}
			data, err := service.Tag.All(params, columns, with)
			if err != nil {
				return nil, err
			}
			return json.Marshal(data)
		})
		str, err := homeTags.Get()
		if err != nil {
			go app.Log().Println(err, string(debug.Stack()))
			response.JsonErr(ctx, errno.CtxShowErr, nil)
			return
		}
		ctx.Set("web_tags", str)
		ctx.Next()
	}
}
