package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yourname/oxo-game-api/services"
)

// JoinChallenge 参加挑战
func JoinChallenge(c *gin.Context) {
	var req struct {
		PlayerID uuid.UUID `json:"player_id" binding:"required"`
		Amount   float64   `json:"amount" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 参加挑战
	challenge, err := services.Challenge.JoinChallenge(req.PlayerID, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 检查下次可参加时间
	_, nextTime, _ := services.Challenge.CheckCooldown(req.PlayerID)

	c.JSON(http.StatusCreated, gin.H{
		"challenge_id":   challenge.ID.String(),
		"status":         "started",
		"message":        "Challenge started, result will be available in 30 seconds",
		"start_time":     challenge.StartedAt,
		"next_available": nextTime,
	})
}

// ListChallengeResults 获取挑战结果列表
func ListChallengeResults(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	challenges, err := services.Challenge.GetRecentResults(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch challenge results",
		})
		return
	}

	// 格式化响应
	results := make([]gin.H, len(challenges))
	for i, challenge := range challenges {
		result := gin.H{
			"id":           challenge.ID,
			"player_id":    challenge.PlayerID,
			"is_winner":    challenge.IsWinner,
			"prize_amount": challenge.PrizeAmount,
			"started_at":   challenge.StartedAt,
			"ended_at":     challenge.EndedAt,
		}
		
		// 如果有玩家信息，添加玩家名称
		if challenge.Player != nil {
			result["player_name"] = challenge.Player.Name
		}
		
		results[i] = result
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    results,
	})
}