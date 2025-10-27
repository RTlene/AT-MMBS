package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// Config 应用配置结构
type Config struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
	JWT      JWTConfig      `json:"jwt"`
	Upload   UploadConfig   `json:"upload"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port         string `json:"port"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteTimeout int    `json:"write_timeout"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
}

// JWTConfig JWT配置
type JWTConfig struct {
	Secret     string `json:"secret"`
	ExpireHour int    `json:"expire_hour"`
}

// UploadConfig 上传配置
type UploadConfig struct {
	MaxSize      int64    `json:"max_size"`       // 最大文件大小（MB）
	AllowedTypes []string `json:"allowed_types"`  // 允许的文件类型
	UploadDir    string   `json:"upload_dir"`     // 上传目录
}

var cfg *Config

// Load 加载配置
func Load() error {
	// 优先从环境变量加载
	cfg = &Config{
		Server: ServerConfig{
			Port:         getEnv("PORT", "80"),
			ReadTimeout:  getEnvAsInt("READ_TIMEOUT", 10),
			WriteTimeout: getEnvAsInt("WRITE_TIMEOUT", 10),
		},
		Database: DatabaseConfig{
			Host:     getEnv("MYSQL_ADDRESS", "localhost:3306"),
			Port:     getEnvAsInt("MYSQL_PORT", 3306),
			User:     getEnv("MYSQL_USERNAME", "root"),
			Password: getEnv("MYSQL_PASSWORD", ""),
			DBName:   getEnv("MYSQL_DATABASE", "golang_demo"),
		},
		JWT: JWTConfig{
			Secret:     getEnv("JWT_SECRET", "your-secret-key-change-this"),
			ExpireHour: getEnvAsInt("JWT_EXPIRE_HOUR", 24),
		},
		Upload: UploadConfig{
			MaxSize:      getEnvAsInt64("UPLOAD_MAX_SIZE", 10),
			AllowedTypes: []string{".jpg", ".jpeg", ".png", ".gif", ".webp"},
			UploadDir:    getEnv("UPLOAD_DIR", "uploads"),
		},
	}
	
	// 尝试从配置文件加载
	configFile := getEnv("CONFIG_FILE", "config.json")
	if _, err := os.Stat(configFile); err == nil {
		data, err := ioutil.ReadFile(configFile)
		if err != nil {
			return fmt.Errorf("读取配置文件失败: %v", err)
		}
		
		if err := json.Unmarshal(data, cfg); err != nil {
			return fmt.Errorf("解析配置文件失败: %v", err)
		}
	}
	
	return nil
}

// Get 获取配置
func Get() *Config {
	if cfg == nil {
		if err := Load(); err != nil {
			panic(fmt.Sprintf("加载配置失败: %v", err))
		}
	}
	return cfg
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt 获取环境变量并转换为整数
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

// getEnvAsInt64 获取环境变量并转换为int64
func getEnvAsInt64(key string, defaultValue int64) int64 {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseInt(valueStr, 10, 64); err == nil {
		return value
	}
	return defaultValue
}