package handler

import (
	"net/http"
	"time"

	"timenote/db"
	"timenote/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateActivityLogRequest struct {
	Title     string     `json:"title"`
	StartTime time.Time  `json:"start_time"`
	EndTime   time.Time  `json:"end_time"`
	Notes     string     `json:"notes"`
	ParentID  *uuid.UUID `json:"parent_id"`
	Tags      []string   `json:"tags"`
}

func CreateActivityLogHandler(c *gin.Context) {
	var req CreateActivityLogRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userID := uuid.New()

	activityLog := model.ActivityLog{
		ID:        uuid.New(),
		UserID:    userID,
		Title:     req.Title,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Notes:     req.Notes,
		ParentID:  req.ParentID,
		CreatedAt: time.Now(),
		Path:      "",
	}

	if err := db.DB.Create(&activityLog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "faild to save!!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      activityLog.ID,
		"message": "ActivityLog created successfully",
	})
}
