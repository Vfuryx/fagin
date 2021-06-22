package api

import (
	"fagin/app/errno"
	"fagin/app/requests/api"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type indexController struct{
	BaseController
}

var IndexController indexController

// Index
// @Summary Add new user to the database
// @Description Add a new user
// @securityDefinitions.basic BasicAuth
// @Tags user
// @Accept  json
// @Produce  json
func (ic *indexController) Index(ctx *gin.Context) {
	var v = api_request.NewIndexRequest()

	if data, ok := v.Validate(ctx); !ok {
		ic.ResponseJsonErr(ctx, errno.Serve.BindErr, data)
		return
	}

	fmt.Println(v)

	// 获取 session
	session := sessions.Default(ctx)
	if id := session.Get("user"); id == nil {
		session.Set("user", 100)
		// 将 sessionID 存入cookie
		if err := session.Save(); err != nil {
			fmt.Println(err)
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
	//db.ORM().Create(&u)
	//
	//app.ResponseJson(ctx, nil, nil)
	//
	//return

	//cm := service.CasbinModel{RoleName:"admin",Path:"/api/v1/",Method:"get"}

	//fmt.Println(service.Casbin.)

	//fmt.Println(service.Casbin.AddRoleForUser("1", "aome"))
	//fmt.Println(service.Casbin.T())
	//service.Casbin.AddCasbin(cm)
	//service.Casbin.ReCasbin(cm)

	//time.Sleep(1 * time.Second)
	//fmt.Println(service.Casbin.E == nil)

	//app.ResponseJson(ctx, nil, gin.H{
	//	"id":       service.Login.ID,
	//	"username": service.Login.Username,
	//	"data":     1111,
	//})
}
