package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mohidex/identity-service/controllers"
	"github.com/mohidex/identity-service/db"
	"github.com/mohidex/identity-service/middleware"
	"github.com/mohidex/identity-service/settings"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	gormDB := settings.GetDB()
	dbInstance := db.NewPgDB(gormDB)

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	v1 := router.Group("v1")
	{
		userHandler := &controllers.UserController{DB: dbInstance}
		v1.POST("/signup", userHandler.Register)
		v1.POST("/login", userHandler.Login)

		userRoutes := v1.Group("/user")
		userRoutes.Use(middleware.JWTAuthMiddleware())
		userRoutes.GET("/me", userHandler.AutorizeToken)
	}
	return router

}
