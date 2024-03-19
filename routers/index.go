package routers

import (
	"gin-rdmg-service/controllers"
	"gin-rdmg-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(r *gin.Engine) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "gin-rdmg-service",
		})
	})

	authGroup := r.Group("/auth")
	authGroup.Use().Use(utils.JWTAuthMiddleware())
	{
		authUser := authGroup.Group("/user")
		authUser.GET("/info", controllers.GetUserInfo)
		authUser.POST("/logout", controllers.Logout)
	}

	noAuthGroup := r.Group("/noauth")
	{
		noAuthUser := noAuthGroup.Group("/user")
		noAuthUser.POST("/login", controllers.Login)
	}

}
