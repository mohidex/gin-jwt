package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mohidex/identity-service/controllers"
	"github.com/mohidex/identity-service/middleware"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	v1 := router.Group("v1")
	{
		user := new(controllers.UserController)
		v1.POST("/signup", user.Register)
		v1.POST("/login", user.Login)

		userRoutes := v1.Group("/user")
		userRoutes.Use(middleware.JWTAuthMiddleware())
		userRoutes.GET("/me", user.AutorizeToken)
	}
	return router

}
