package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"wxcloudrun-golang/config"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
	"wxcloudrun-golang/middleware"
	"wxcloudrun-golang/router"
	"wxcloudrun-golang/service"

	"gorm.io/gorm"
)

func main() {
	// 加载配置
	if err := config.Load(); err != nil {
		log.Printf("Warning: failed to load config: %v, using default values", err)
	}
	cfg := config.Get()

	// 初始化数据库
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	// 初始化数据库表
	if err := initTables(); err != nil {
		panic(fmt.Sprintf("init tables failed with %+v", err))
	}

	// 创建路由器
	r := router.NewRouter()
	r.SetupRoutes()

	// 启动服务器
	srv := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      middleware.LoggerMiddleware(r),
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
	}

	log.Printf("Server starting on port %s...", cfg.Server.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

// initTables 初始化数据库表
func initTables() error {
	dbInstance := db.Get()

	// 自动迁移所有模型
	models := []interface{}{
		&model.UserModel{},
		&model.ProductModel{},
		&model.CategoryModel{},
		&model.MemberModel{},
		&model.MemberLevelModel{},
		&model.DistributorLevelModel{},
		&model.OrderModel{},
		&model.PaymentModel{},
	}

	for _, model := range models {
		if err := dbInstance.AutoMigrate(model); err != nil {
			return fmt.Errorf("failed to migrate %T: %v", model, err)
		}
	}

	// 创建默认管理员用户
	if err := createDefaultAdminUser(dbInstance); err != nil {
		log.Printf("Warning: failed to create default admin user: %v", err)
	}

	log.Println("Database tables initialized successfully")
	return nil
}

// createDefaultAdminUser 创建默认管理员用户
func createDefaultAdminUser(dbInstance *gorm.DB) error {
	// 检查是否已存在管理员用户
	var count int64
	if err := dbInstance.Model(&model.UserModel{}).Where("username = ?", "admin").Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		log.Println("Default admin user already exists")
		return nil
	}

	// 创建默认管理员用户
	adminUser := &model.UserModel{
		ID:       "admin-" + time.Now().Format("20060102150405"),
		Username: "admin",
		Password: service.HashPassword("admin123"), // 默认密码：admin123
		Role:     "admin",
	}

	if err := dbInstance.Create(adminUser).Error; err != nil {
		return err
	}

	log.Println("Default admin user created successfully: username=admin, password=admin123")
	return nil
}