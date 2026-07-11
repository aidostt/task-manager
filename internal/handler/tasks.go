package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/aidostt/task-manager/internal/model"
	"github.com/aidostt/task-manager/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) createTask(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	var input model.CreateTaskInput
	if err = c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid input body"})
		return
	}
	task := model.Task{
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
		Priority:    input.Priority,
		UserID:      userID,
	}
	createdTask, err := h.services.Task.Create(c, &task)
	if err != nil {
		log.Println("createTask error:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"task": createdTask})
}

func (h *Handler) getTaskByID(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	taskID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid task id"})
		return
	}
	task, err := h.services.Task.FindByID(c, taskID, userID)
	if err != nil {
		if errors.Is(err, service.ErrForbidden) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"task": task})
}
func (h *Handler) getTasksByUserID(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	tasks, err := h.services.Task.FindAllByUserID(c, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (h *Handler) updateTask(c *gin.Context) {
	taskID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid task id"})
		return
	}
	userID, err := getUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	var input model.UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid input body"})
		return
	}
	task := model.Task{
		ID:          taskID,
		UserID:      userID,
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
		Priority:    input.Priority,
	}
	err = h.services.Task.Update(c, &task)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidTask):
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid task id"})
			return
		case errors.Is(err, service.ErrForbidden):
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})

}
func (h *Handler) deleteTask(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	taskID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid task id"})
		return
	}
	err = h.services.Task.Delete(c, taskID, userID)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidTask):
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid task id"})
			return
		case errors.Is(err, service.ErrForbidden):
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
