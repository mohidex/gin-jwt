package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mohidex/identity-service/auth"
	"github.com/mohidex/identity-service/controllers"
	"github.com/mohidex/identity-service/db"
	"github.com/mohidex/identity-service/metricsutil"
	"github.com/mohidex/identity-service/middlewares"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Routes struct {
	db          db.Database
	auth        auth.Authenticator
	metricsutil metricsutil.Metrics
}

func NewRoutes(db db.Database, auth auth.Authenticator, metricsutil metricsutil.Metrics) *Routes {
	return &Routes{
		db:          db,
		auth:        auth,
		metricsutil: metricsutil,
	}
}

func (r *Routes) Setup(router *gin.Engine) {

	prometheus.DefaultRegisterer.Unregister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))
	prometheus.DefaultRegisterer.Unregister(prometheus.NewGoCollector())

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.Use(middlewares.PrometheusMiddleware(r.metricsutil))

	health := &controllers.HealthController{}

	router.GET("/health", health.Status)

	// Create a new route group
	v1 := router.Group("v1")
	{
		userHandler := &controllers.UserController{DB: r.db, Auth: r.auth}
		v1.POST("/signup", userHandler.Register)
		v1.POST("/login", userHandler.Login)

		userRoutes := v1.Group("/user")
		userRoutes.Use(middlewares.AuthMiddleware(r.auth))
		userRoutes.GET("/me", userHandler.AutorizeToken)

		adminRoutes := v1.Group("/admin")
		adminRoutes.Use(middlewares.AuthMiddleware(r.auth))
		adminRoutes.Use(middlewares.AdminMiddleware())
		adminRoutes.GET("/users", userHandler.GetUsers)
	}
}
