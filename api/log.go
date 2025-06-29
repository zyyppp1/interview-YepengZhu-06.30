package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zyyppp1/interview-YepengZhu-06.30/models"
	"github.com/zyyppp1/interview-YepengZhu-06.30/services"
)

// ListGameLogs 获取游戏日志列表
func ListGameLogs(c *gin.Context) {
	// 解析查询参数
	var playerID *uuid.UUID
	var startTime, endTime *time.Time
	
	if playerIDStr := c.Query("player_id"); playerIDStr != "" {
		if id, err := uuid.Parse(playerIDStr); err == nil {
			playerID = &id
		}
	}
	
	actionType := c.Query("action")
	
	if startTimeStr := c.Query("start_time"); startTimeStr != "" {
		if t, err := time.Parse(time.RFC3339, startTimeStr); err == nil {
			startTime = &t
		}
	}
	
	if endTimeStr := c.Query("end_time"); endTimeStr != "" {
		if t, err := time.Parse(time.RFC3339, endTimeStr); err == nil {
			endTime = &t
		}
	}
	
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
	if limit <= 0 || limit > 100 {
		limit = 50
	}

	// 获取日志
	logs, err := services.GameLog.GetAll(playerID, actionType, startTime, endTime, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch logs",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    logs,
	})
}

// CreateGameLog 创建游戏日志
func CreateGameLog(c *gin.Context) {
	var req struct {
		PlayerID   *uuid.UUID             `json:"player_id,omitempty"`
		ActionType string                 `json:"action_type" binding:"required"`
		Details    map[string]interface{} `json:"details,omitempty"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 创建日志
	log := &models.GameLog{
		PlayerID:   req.PlayerID,
		ActionType: req.ActionType,
		Details:    req.Details,
		IP:         c.ClientIP(),
		UserAgent:  c.GetHeader("User-Agent"),
	}

	if err := services.GameLog.Create(log); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data": gin.H{
			"id": log.ID,
		},
	})
}