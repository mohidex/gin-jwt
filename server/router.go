package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mohidex/identity-service/auth"
	"github.com/mohidex/identity-service/controllers"
	"github.com/mohidex/identity-service/db"
	"github.com/mohidex/identity-service/middleware"
)

type Routes struct {
	db   db.Database
	auth *auth.JWTAuthenticator
}

func NewRoutes(db db.Database, auth *auth.JWTAuthenticator) *Routes {
	return &Routes{
		db:   db,
		auth: auth,
	}
}

func (r *Routes) Setup(router *gin.Engine) {

	health := &controllers.HealthController{}

	router.GET("/health", health.Status)

	// Create a new route group
	v1 := router.Group("v1")
	{
		userHandler := &controllers.UserController{DB: r.db, Auth: r.auth}
		v1.POST("/signup", userHandler.Register)
		v1.POST("/login", userHandler.Login)

		userRoutes := v1.Group("/user")
		userRoutes.Use(middleware.AuthMiddleware(r.auth))
		userRoutes.GET("/me", userHandler.AutorizeToken)

		adminRoutes := v1.Group("/admin")
		adminRoutes.Use(middleware.AuthMiddleware(r.auth))
		adminRoutes.Use(middleware.AdminMiddleware())
		adminRoutes.GET("/users", userHandler.GetUsers)
	}
}
