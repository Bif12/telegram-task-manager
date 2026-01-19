package handlers

import (
	"net/http"
	"strconv"

	"github.com/bif12/telegram-task-manager/internal/database"
	"github.com/bif12/telegram-task-manager/internal/middleware"
	"github.com/bif12/telegram-task-manager/internal/models"
	"github.com/gin-gonic/gin"
)

type MilestoneHandler struct{}

func NewMilestoneHandler() *MilestoneHandler {
	return &MilestoneHandler{}
}

func (h *MilestoneHandler) Create(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)
	var milestone models.Milestone

	if err := c.ShouldBindJSON(&milestone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify goal belongs to user (optional but good practice)
	var goal models.Goal
	if err := database.DB.Where("id = ? AND user_id = ?", milestone.GoalID, userID).First(&goal).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Goal not found"})
		return
	}

	if err := database.DB.Create(&milestone).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create milestone"})
		return
	}

	c.JSON(http.StatusCreated, milestone)
}

func (h *MilestoneHandler) Update(c *gin.Context) {
	// userID, _ := middleware.GetUserID(c) // Use if we want to verify ownership via goal
	id, _ := strconv.Atoi(c.Param("id"))

	var milestone models.Milestone
	if err := database.DB.First(&milestone, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Milestone not found"})
		return
	}

	if err := c.ShouldBindJSON(&milestone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&milestone).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update milestone"})
		return
	}

	c.JSON(http.StatusOK, milestone)
}

func (h *MilestoneHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := database.DB.Delete(&models.Milestone{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete milestone"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Milestone deleted"})
}
