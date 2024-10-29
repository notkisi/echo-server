package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) echoGet(ctx *gin.Context) {
	message := ctx.Query("message")
	if message == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "message is required"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": message})
}

func (s *Server) echoPost(ctx *gin.Context) {
	var request struct {
		Message string `json:"message" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": request.Message})
}
