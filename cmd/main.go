package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yourname/oxo-game-api/config"
	"github.com/yourname/oxo-game-api/db"
	"github.com/yourname/oxo-game-api/routes"
	"github.com/yourname/oxo-game-api/services"
)

func main() {
	// 加载配置
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 设置 Gin 模式
	gin.SetMode(cfg.Server.Mode)

	// 初始化数据库
	if err := db.Init(cfg); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 初始化服务
	services.InitServices()

	// 设置路由
	router := routes.SetupRouter()

	// 启动服务器
	log.Printf("Starting server on port %s", cfg.Server.Port)
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}