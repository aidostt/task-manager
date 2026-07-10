package handler

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) authMiddleware(c *gin.Context) {
	token := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	if token == "" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}
	claims, err := h.tokenManager.Parse(token)
	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}
	c.Set("userID", claims.UserID)
	c.Next()
}
