package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yourname/oxo-game-api/models"
	"github.com/yourname/oxo-game-api/services"
)

// ListPlayers 获取玩家列表
func ListPlayers(c *gin.Context) {
	// 获取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	// 获取数据
	players, total, err := services.Player.GetAll(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch players",
		})
		return
	}

	// 计算总页数
	totalPages := int(total) / pageSize
	if int(total)%pageSize > 0 {
		totalPages++
	}

	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    players,
		"meta": gin.H{
			"page":        page,
			"page_size":   pageSize,
			"total":       total,
			"total_pages": totalPages,
		},
	})
}

// CreatePlayer 创建玩家
func CreatePlayer(c *gin.Context) {
	var req struct {
		Name    string    `json:"name" binding:"required,min=2,max=50"`
		LevelID uuid.UUID `json:"level_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 创建玩家
	player := &models.Player{
		Name:    req.Name,
		LevelID: req.LevelID,
	}

	if err := services.Player.Create(player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    player,
	})
}

// GetPlayer 获取单个玩家
func GetPlayer(c *gin.Context) {
	// 解析ID
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid player ID",
		})
		return
	}

	// 获取玩家
	player, err := services.Player.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    player,
	})
}

// UpdatePlayer 更新玩家
func UpdatePlayer(c *gin.Context) {
	// 解析ID
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid player ID",
		})
		return
	}

	// 解析请求
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
	if levelID, ok := req["level_id"].(string); ok {
		if uid, err := uuid.Parse(levelID); err == nil {
			updates["level_id"] = uid
		}
	}

	if len(updates) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No valid fields to update",
		})
		return
	}

	// 更新玩家
	if err := services.Player.Update(id, updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 获取更新后的玩家信息
	player, _ := services.Player.GetByID(id)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    player,
	})
}

// DeletePlayer 删除玩家
func DeletePlayer(c *gin.Context) {
	// 解析ID
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid player ID",
		})
		return
	}

	// 删除玩家
	if err := services.Player.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Player deleted successfully",
	})
}

// ListLevels 获取等级列表
func ListLevels(c *gin.Context) {
	levels, err := services.Level.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch levels",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    levels,
	})
}

// CreateLevel 创建等级
func CreateLevel(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required,min=2,max=30"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	level := &models.Level{
		Name: req.Name,
	}

	if err := services.Level.Create(level); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    level,
	})
}