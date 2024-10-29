package api

import (
	"github.com/gin-gonic/gin"
	"os"
	"testing"
)

func newTestServer(t *testing.T) *Server {
	router := gin.Default()
	server := &Server{
		router: router,
	}

	server.routes()
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
