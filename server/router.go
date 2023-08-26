package server

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mohidex/identity-service/auth"
	"github.com/mohidex/identity-service/controllers"
	"github.com/mohidex/identity-service/db"
	"github.com/mohidex/identity-service/middleware"
	"github.com/mohidex/identity-service/settings"
)

var (
	privateKey = os.Getenv("JWT_PRIVATE_KEY")
	tokenTTL   = os.Getenv("JWT_TTL")
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	gormDB := settings.GetDB()
	dbInstance := db.NewPgDB(gormDB)
	tokenTtl, _ := strconv.Atoi(tokenTTL)
	jwtAuthenticator := auth.NewJWTAuthenticator(privateKey, tokenTtl)

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	v1 := router.Group("v1")
	{
		userHandler := &controllers.UserController{DB: dbInstance, Auth: jwtAuthenticator}
		v1.POST("/signup", userHandler.Register)
		v1.POST("/login", userHandler.Login)

		userRoutes := v1.Group("/user")
		userRoutes.Use(middleware.AuthMiddleware(jwtAuthenticator))
		userRoutes.GET("/me", userHandler.AutorizeToken)
	}
	return router

}
