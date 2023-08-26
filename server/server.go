package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mohidex/identity-service/auth"
	"github.com/mohidex/identity-service/db"
)

type Server struct {
	router *gin.Engine
	db     db.Database
	auth   *auth.JWTAuthenticator
}

func NewServer(db db.Database, auth *auth.JWTAuthenticator) *Server {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	s := &Server{
		router: r,
		db:     db,
		auth:   auth,
	}
	s.setupRoutes()
	return s
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}

func (s *Server) setupRoutes() {
	// Initialize routes with dependencies
	routes := NewRoutes(s.db, s.auth)
	routes.Setup(s.router)
}
