package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config 应用配置结构
type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Server   ServerConfig
}

// AppConfig 应用配置
type AppConfig struct {
	Name string
	Env  string
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port string
	Mode string
}

// Load 加载配置
func Load() (*Config, error) {
	// 加载 .env 文件
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	config := &Config{
		App: AppConfig{
			Name: getEnv("APP_NAME", "OXO Game API"),
			Env:  getEnv("APP_ENV", "development"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "oxogame"),
			Password: getEnv("DB_PASSWORD", "oxogame123"),
			DBName:   getEnv("DB_NAME", "oxogame_db"),
		},
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
			Mode: getEnv("GIN_MODE", "debug"),
		},
	}

	return config, nil
}

// getEnv 获取环境变量
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}