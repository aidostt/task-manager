package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getUserID(c *gin.Context) (uuid.UUID, error) {
	userIdStr := c.MustGet("userID").(string)
	return uuid.Parse(userIdStr)
}
