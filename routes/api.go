package routes

import (
	c "fagin/app/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine{

	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())


	api := r.Group("/api/v1")
	{
		api.GET("/login", c.AuthController{}.Login)
		api.GET("/logout", c.AuthController{}.Logout)
		api.GET("/add", c.AuthController{}.CreateUser)
	}


	return r
}




