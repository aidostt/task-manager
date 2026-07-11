package handler

import (
	"github.com/aidostt/task-manager/internal/service"
	"github.com/aidostt/task-manager/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services     *service.Models
	tokenManager jwt.TokenManager
}

func NewHandler(services *service.Models, tokenManager jwt.TokenManager) *Handler {
	return &Handler{services: services, tokenManager: tokenManager}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	auth := router.Group("/auth")
	{
		auth.POST("/login", h.login)
		auth.POST("/register", h.register)
		auth.POST("/refresh", h.refresh)
	}
	api := router.Group("/api", h.authMiddleware)
	{
		task := api.Group("/tasks")
		{
			task.POST("/", h.createTask)
			task.GET("/", h.getTasksByUserID)
			task.GET("/:id", h.getTaskByID)
			task.PUT("/:id", h.updateTask)
			task.DELETE("/:id", h.deleteTask)
		}
	}
	return router
}
