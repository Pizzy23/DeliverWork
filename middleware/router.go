package middleware

import (
	"deliver/docs"
	login "deliver/internal/login/handlers"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Sua API de Comida
// @version 1.0
// @description API para gerenciar informações de alimentos
// @termsOfService http://swagger.io/terms/
// @host localhost:8080
// @BasePath /api
func SetupRouter() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"
	v1 := r.Group("/api")
	{
		v1.PUT("/login", login.LoginUser)
		v1.PUT("/loggout", login.LoggedUser)
		v1.POST("/register", login.CreateUser)
		v1.POST("/auth", login.AuthUser)
		v1.POST("/token", login.TokenUser)
		v1.GET("/user", login.SeachUser)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
