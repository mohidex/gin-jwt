package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mohidex/identity-service/auth"
	"github.com/mohidex/identity-service/db"
	"github.com/mohidex/identity-service/metricsutil"
)

type Server struct {
	router      *gin.Engine
	db          db.Database
	auth        auth.Authenticator
	metricsutil metricsutil.Metrics
}

func NewServer(db db.Database, auth auth.Authenticator, metricsutil metricsutil.Metrics) *Server {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	s := &Server{
		router:      r,
		db:          db,
		auth:        auth,
		metricsutil: metricsutil,
	}
	s.setupRoutes()
	return s
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}

func (s *Server) setupRoutes() {
	// Initialize routes with dependencies
	routes := NewRoutes(s.db, s.auth, s.metricsutil)
	routes.Setup(s.router)
}
