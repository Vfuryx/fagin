package api

import (
	"fagin/app"
	"fagin/app/errno"
	"fagin/app/requests/api"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
	"net/http"
)

type indexController struct{}

var IndexController indexController

// @Summary Add new user to the database
// @Description Add a new user
// @securityDefinitions.basic BasicAuth
// @Tags user
// @Accept  json
// @Produce  json
func (indexController) Index(ctx *gin.Context) {
	var v api_request.IndexRequest
	data, ok := v.Validate(ctx)
	if !ok {
		app.JsonResponse(ctx, errno.ErrBind, data)
		ctx.Abort()
		return
	}

	fmt.Println(v)

	// 获取 session
	session := sessions.Default(ctx)
	if id := session.Get("user"); id == nil {
		session.Set("user", 100)
		// 将 sessionID 存入cookie
		if err := session.Save(); err != nil {
			log.Info(err)
		}
	}

	ctx.String(http.StatusOK, "%v", session.Get("user"))

	//u := user.User{
	//	Username:	"xfy",
	//	Info: 		info.Info{
	//		Email:		"12121@qq.com",
	//	},
	//}
	//
	//db.ORM.Create(&u)
	//
	//app.JsonResponse(ctx, nil, nil)
	//
	//return

	//cm := service.CasbinModel{RoleName:"admin",Path:"/api/v1/",Method:"get"}

	//fmt.Println(service.Canbin.)

	//fmt.Println(service.Canbin.AddRoleForUser("1", "aome"))
	//fmt.Println(service.Canbin.T())
	//service.Canbin.AddCasbin(cm)
	//service.Canbin.ReCasbin(cm)

	//time.Sleep(1 * time.Second)
	//fmt.Println(service.Canbin.E == nil)

	//app.JsonResponse(ctx, nil, gin.H{
	//	"id":       service.Login.ID,
	//	"username": service.Login.Username,
	//	"data":     1111,
	//})
}
