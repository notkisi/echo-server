package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	router := gin.Default()
	server := &Server{
		router: router,
	}

	server.routes()
	return server
}

func (s *Server) routes() {
	s.router.GET("/echo", s.echoGet)
	s.router.POST("/echo", s.echoPost)
}

func (s *Server) Run() {
	s.router.Run(":8080")
}
