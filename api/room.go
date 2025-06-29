package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yourname/oxo-game-api/models"
	"github.com/yourname/oxo-game-api/services"
)

// ListRooms 获取房间列表
func ListRooms(c *gin.Context) {
	status := c.Query("status")
	
	rooms, err := services.Room.GetAll(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch rooms",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    rooms,
	})
}

// CreateRoom 创建房间
func CreateRoom(c *gin.Context) {
	var req struct {
		Name        string `json:"name" binding:"required,min=2,max=50"`
		Description string `json:"description"`
		MaxPlayers  int    `json:"max_players,omitempty"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	room := &models.Room{
		Name:        req.Name,
		Description: req.Description,
		MaxPlayers:  req.MaxPlayers,
		Status:      "available",
	}

	if err := services.Room.Create(room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    room,
	})
}

// GetRoom 获取单个房间
func GetRoom(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid room ID",
		})
		return
	}

	room, err := services.Room.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    room,
	})
}

// UpdateRoom 更新房间
func UpdateRoom(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid room ID",
		})
		return
	}

	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 只允许更新特定字段
	updates := make(map[string]interface{})
	if name, ok := req["name"].(string); ok && name != "" {
		updates["name"] = name
	}
	if description, ok := req["description"].(string); ok {
		updates["description"] = description
	}
	if status, ok := req["status"].(string); ok {
		updates["status"] = status
	}
	if maxPlayers, ok := req["max_players"].(float64); ok {
		updates["max_players"] = int(maxPlayers)
	}

	if len(updates) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No valid fields to update",
		})
		return
	}

	if err := services.Room.Update(id, updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	room, _ := services.Room.GetByID(id)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    room,
	})
}

// DeleteRoom 删除房间
func DeleteRoom(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid room ID",
		})
		return
	}

	if err := services.Room.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Room deleted successfully",
	})
}

// ListReservations 获取预约列表
func ListReservations(c *gin.Context) {
	// 解析查询参数
	var roomID *uuid.UUID
	var date *time.Time
	
	if roomIDStr := c.Query("room_id"); roomIDStr != "" {
		if id, err := uuid.Parse(roomIDStr); err == nil {
			roomID = &id
		}
	}
	
	if dateStr := c.Query("date"); dateStr != "" {
		if d, err := time.Parse("2006-01-02", dateStr); err == nil {
			date = &d
		}
	}
	
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if limit <= 0 || limit > 100 {
		limit = 20
	}

	reservations, err := services.Reservation.GetAll(roomID, date, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch reservations",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    reservations,
	})
}

// CreateReservation 创建预约
func CreateReservation(c *gin.Context) {
	var req struct {
		RoomID          uuid.UUID `json:"room_id" binding:"required"`
		PlayerID        uuid.UUID `json:"player_id" binding:"required"`
		ReservationDate time.Time `json:"reservation_date" binding:"required"`
		StartTime       string    `json:"start_time" binding:"required"`
		EndTime         string    `json:"end_time" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 验证时间格式
	if !isValidTimeFormat(req.StartTime) || !isValidTimeFormat(req.EndTime) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid time format, use HH:MM",
		})
		return
	}

	// 验证时间顺序
	if req.StartTime >= req.EndTime {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "End time must be after start time",
		})
		return
	}

	reservation := &models.Reservation{
		RoomID:          req.RoomID,
		PlayerID:        req.PlayerID,
		ReservationDate: req.ReservationDate,
		StartTime:       req.StartTime,
		EndTime:         req.EndTime,
		Status:          "active",
	}

	if err := services.Reservation.Create(reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    reservation,
	})
}

// isValidTimeFormat 验证时间格式是否为 HH:MM
func isValidTimeFormat(timeStr string) bool {
	_, err := time.Parse("15:04", timeStr)
	return err == nil
}