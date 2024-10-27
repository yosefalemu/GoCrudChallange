package server

import (
	"github.com/gin-gonic/gin"
	"github.com/yosefalemu/GoCrudChallange/middlewares"
	"github.com/yosefalemu/GoCrudChallange/routes"
)

// Server represents the server using gin framework
type Server struct {
	Router *gin.Engine
}

// NewServer creates a new server
func NewServer() *Server {
	return &Server{
		Router: gin.Default(),
	}
}

// Initialize initializes the server
func (s *Server) Initialize() {
	s.Router.Use(middlewares.CORSMiddleware(), middlewares.ErrorHandlerMiddleware())
	routes.RegisterRoutes(s.Router)
}

// Run runs the server
func (s *Server) Run(addr string) {
	s.Router.Run(addr)
}
