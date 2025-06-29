package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourname/oxo-game-api/api"
	"github.com/yourname/oxo-game-api/middleware"
)

// SetupRouter 配置路由
func SetupRouter() *gin.Engine {
	router := gin.New()

	// 全局中间件
	router.Use(middleware.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORS())

	// 健康检查
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "OXO Game API",
		})
	})

	// API v1 路由组
	v1 := router.Group("/api/v1")
	{
		// 玩家管理
		players := v1.Group("/players")
		{
			players.GET("", api.ListPlayers)
			players.POST("", api.CreatePlayer)
			players.GET("/:id", api.GetPlayer)
			players.PUT("/:id", api.UpdatePlayer)
			players.DELETE("/:id", api.DeletePlayer)
		}

		// 等级管理
		levels := v1.Group("/levels")
		{
			levels.GET("", api.ListLevels)
			levels.POST("", api.CreateLevel)
		}

		// 房间管理
		rooms := v1.Group("/rooms")
		{
			rooms.GET("", api.ListRooms)
			rooms.POST("", api.CreateRoom)
			rooms.GET("/:id", api.GetRoom)
			rooms.PUT("/:id", api.UpdateRoom)
			rooms.DELETE("/:id", api.DeleteRoom)
		}

		// 预约管理
		v1.GET("/reservations", api.ListReservations)
		v1.POST("/reservations", api.CreateReservation)

		// 挑战系统
		v1.POST("/challenges", api.JoinChallenge)
		v1.GET("/challenges/results", api.ListChallengeResults)

		// 游戏日志
		v1.GET("/logs", api.ListGameLogs)
		v1.POST("/logs", api.CreateGameLog)

		// 支付系统
		v1.POST("/payments", api.ProcessPayment)
		v1.GET("/payments/:id", api.GetPayment)
	}

	// 404处理
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Resource not found",
		})
	})

	return router
}