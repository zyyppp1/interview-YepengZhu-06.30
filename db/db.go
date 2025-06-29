package db

import (
	"fmt"
	"log"

	"github.com/yourname/oxo-game-api/config"
	"github.com/yourname/oxo-game-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Init 初始化数据库连接
func Init(cfg *config.Config) error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.Port,
	)

	gormConfig := &gorm.Config{}
	
	if cfg.App.Env == "production" {
		gormConfig.Logger = logger.Default.LogMode(logger.Silent)
	} else {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	// 自动迁移
	if err := AutoMigrate(); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("Database connected successfully")
	return nil
}

// AutoMigrate 自动迁移所有模型
func AutoMigrate() error {
	return DB.AutoMigrate(
		&models.Level{},
		&models.Player{},
		&models.Room{},
		&models.Reservation{},
		&models.Challenge{},
		&models.PrizePool{},
		&models.GameLog{},
		&models.Payment{},
	)
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	return DB
}