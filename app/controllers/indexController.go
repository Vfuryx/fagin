package controllers

import (
	"fagin/app"
	"fagin/app/models/info"
	"fagin/app/models/user"
	"fagin/pkg/db"
	"github.com/gin-gonic/gin"
)

type indexController struct{}

var IndexController = indexController{}


// @Summary Add new user to the database
// @Description Add a new user
// @Tags user
// @Accept  json
// @Produce  json
func (indexController) Index(ctx *gin.Context) {

	u := user.User{
		Username:	"xfy",
		Info: 		info.Info{
			Email:		"12121@qq.com",
		},
	}

	db.ORM.Create(&u)

	app.JosnResponse(ctx, nil, nil)

	return

	//cm := service.CasbinModel{RoleName:"admin",Path:"/api/v1/",Method:"get"}

	//fmt.Println(service.Canbin.)

	//fmt.Println(service.Canbin.AddRoleForUser("1", "aome"))
	//fmt.Println(service.Canbin.T())
	//service.Canbin.AddCasbin(cm)
	//service.Canbin.ReCasbin(cm)

	//time.Sleep(1 * time.Second)
	//fmt.Println(service.Canbin.E == nil)

	//app.JosnResponse(ctx, nil, gin.H{
	//	"id":       service.Login.ID,
	//	"username": service.Login.Username,
	//	"data":     1111,
	//})
}
