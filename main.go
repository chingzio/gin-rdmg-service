package main

import (
	"fmt"
	"gin-rdmg-service/models"
	"gin-rdmg-service/routers"
	"gin-rdmg-service/utils"
	"net/http"

	_ "gin-rdmg-service/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

// @title gin-k8s-api swagger
// @version 1.0
// @description gin-k8s-api swagger
// @contact.name chingzio
// @contact.email chingzio@163.com
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-Token
func main() {
	r := gin.Default()
	//r.Use(favicon.New("favicon.ico"))
	r.LoadHTMLGlob("templates/**/*")
	r.StaticFS("/static", http.Dir("static/layui"))
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	routers.Index(r)
	fmt.Println("Server run at:")
	fmt.Printf("-  Local:   http://localhost%v/ \r\n", utils.Port)
	fmt.Printf("-  Swagger: http://localhost%v/swagger/index.html \r\n", utils.Port)
	r.Run(utils.Port)
	defer models.CloseDb()
}
