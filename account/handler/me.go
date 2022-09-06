package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Me handler calls services for getting
// a user's details
func (h *Handler) Me(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's me",
	})
}
