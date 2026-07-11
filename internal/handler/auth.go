package handler

import (
	"errors"
	"net/http"

	"github.com/aidostt/task-manager/internal/model"
	"github.com/aidostt/task-manager/internal/service"
	"github.com/gin-gonic/gin"
)

func (h *Handler) register(c *gin.Context) {
	var input model.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input body"})
		return
	}
	accessToken, refreshToken, err := h.services.User.RegisterUser(c, input.Email, input.Password)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrUserAlreadyExists):
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
			return
		case errors.Is(err, service.ErrInvalidCredentials):
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input body"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
	}
	c.JSON(http.StatusCreated, gin.H{"token": accessToken, "refresh": refreshToken})
}
func (h *Handler) login(c *gin.Context) {
	var input model.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input body"})
		return
	}
	accessToken, refreshToken, err := h.services.User.LoginUser(c, input.Email, input.Password)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidCredentials):
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input body"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
	}
	c.JSON(http.StatusCreated, gin.H{"token": accessToken, "refresh": refreshToken})
}
