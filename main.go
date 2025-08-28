package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
	"wxcloudrun-golang/service"
	"gorm.io/gorm"
	"time"
)

// RESTfulMux 自定义路由处理器
type RESTfulMux struct{}

func (m *RESTfulMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	
	// 处理用户相关的RESTful路由
	if strings.HasPrefix(path, "/api/users") {
		parts := strings.Split(path, "/")
		if len(parts) == 4 && parts[3] != "" {
			// /api/users/{id} - 单个用户操作
			userID := parts[3]
			handleSingleUser(w, r, userID)
			return
		} else if len(parts) == 3 {
			// /api/users - 用户列表操作
			service.UserHandler(w, r)
			return
		}
	}
	
	// 处理商品相关的RESTful路由
	if strings.HasPrefix(path, "/api/products") {
		parts := strings.Split(path, "/")
		if len(parts) == 4 && parts[3] != "" {
			// /api/products/{id} - 单个商品操作
			productID := parts[3]
			handleSingleProduct(w, r, productID)
			return
		} else if len(parts) == 3 {
			// /api/products - 商品列表操作
			service.ProductHandler(w, r)
			return
		}
	}
	
	// 处理分类相关的RESTful路由
	if strings.HasPrefix(path, "/api/categories") {
		parts := strings.Split(path, "/")
		if len(parts) == 4 && parts[3] != "" {
			// /api/categories/{id} - 单个分类操作
			categoryID := parts[3]
			handleSingleCategory(w, r, categoryID)
			return
		} else if len(parts) == 3 {
			// /api/categories - 分类列表操作
			service.CategoryHandler(w, r)
			return
		}
	}
	
	// 处理其他API路由
	switch path {
	case "/api/count":
		service.CounterHandler(w, r)
	case "/api/products/status":
		service.ProductStatusHandler(w, r)
	case "/api/categories/parent":
		service.ParentCategoryHandler(w, r)
	case "/api/members":
		service.MemberHandler(w, r)
	case "/api/members/network":
		service.MemberNetworkHandler(w, r)
	case "/api/members/distribution":
		service.MemberDistributionHandler(w, r)
	case "/api/members/consumption":
		service.MemberConsumptionHandler(w, r)
	case "/api/member-levels":
		service.MemberLevelHandler(w, r)
	case "/api/distributor-levels":
		service.DistributorLevelHandler(w, r)
	case "/api/orders":
		service.OrderHandler(w, r)
	case "/api/orders/status":
		service.OrderStatusHandler(w, r)
	case "/api/payments":
		service.PaymentHandler(w, r)
	case "/api/payments/status":
		service.PaymentStatusHandler(w, r)
	case "/api/upload":
		service.FileUploadHandler(w, r)
	case "/api/upload/delete":
		service.DeleteFile(w, r)
	default:
		http.NotFound(w, r)
	}
}

// handleSingleUser 处理单个用户的操作
func handleSingleUser(w http.ResponseWriter, r *http.Request, userID string) {
	res := &service.JsonResult{}
	
	switch r.Method {
	case http.MethodGet:
		// 获取单个用户
		user, err := service.GetUserByID(userID)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Code = 0
			res.Data = user
		}
		service.WriteJSON(w, res)
		
	case http.MethodPut:
		// 更新用户
		var userData map[string]interface{}
		if err := service.DecodeJSON(r, &userData); err != nil {
			res.Code = -1
			res.ErrorMsg = "请求参数解析失败"
		} else {
			// 只更新提供的字段，避免覆盖时间字段
			updateData := make(map[string]interface{})
			if username, ok := userData["username"].(string); ok && username != "" {
				updateData["username"] = username
			}
			if password, ok := userData["password"].(string); ok && password != "" {
				updateData["password"] = service.HashPassword(password)
			}
			if role, ok := userData["role"].(string); ok && role != "" {
				updateData["role"] = role
			}
			
			if len(updateData) == 0 {
				res.Code = -1
				res.ErrorMsg = "没有提供要更新的字段"
			} else {
				updateData["id"] = userID
				if err := service.UpdateUserPartial(userID, updateData); err != nil {
					res.Code = -1
					res.ErrorMsg = err.Error()
				} else {
					res.Code = 0
					res.Data = "用户更新成功"
				}
			}
		}
		service.WriteJSON(w, res)
		
	case http.MethodDelete:
		// 删除用户
		if err := service.DeleteUser(userID); err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Code = 0
			res.Data = "删除成功"
		}
		service.WriteJSON(w, res)
		
	default:
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
		service.WriteJSON(w, res)
	}
}

// handleSingleProduct 处理单个商品的操作
func handleSingleProduct(w http.ResponseWriter, r *http.Request, productID string) {
	res := &service.JsonResult{}
	
	switch r.Method {
	case http.MethodGet:
		// 获取单个商品
		product, err := service.GetProductByID(productID)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Code = 0
			res.Data = product
		}
		service.WriteJSON(w, res)
		
	case http.MethodPut:
		// 更新商品
		var product model.ProductModel
		if err := service.DecodeJSON(r, &product); err != nil {
			res.Code = -1
			res.ErrorMsg = "请求参数解析失败"
		} else {
			product.ID = productID
			if err := service.UpdateProduct(&product); err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Code = 0
				res.Data = product
			}
		}
		service.WriteJSON(w, res)
		
	case http.MethodDelete:
		// 删除商品
		if err := service.DeleteProduct(productID); err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Code = 0
			res.Data = "删除成功"
		}
		service.WriteJSON(w, res)
		
	default:
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
		service.WriteJSON(w, res)
	}
}

// handleSingleCategory 处理单个分类的操作
func handleSingleCategory(w http.ResponseWriter, r *http.Request, categoryID string) {
	res := &service.JsonResult{}
	
	switch r.Method {
	case http.MethodGet:
		// 获取单个分类
		category, err := service.GetCategoryByID(categoryID)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Code = 0
			res.Data = category
		}
		service.WriteJSON(w, res)
		
	case http.MethodPut:
		// 更新分类
		var category model.CategoryModel
		if err := service.DecodeJSON(r, &category); err != nil {
			res.Code = -1
			res.ErrorMsg = "请求参数解析失败"
		} else {
			category.ID = categoryID
			if err := service.UpdateCategory(&category); err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Code = 0
				res.Data = category
			}
		}
		service.WriteJSON(w, res)
		
	case http.MethodDelete:
		// 删除分类
		if err := service.DeleteCategory(categoryID); err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Code = 0
			res.Data = "删除成功"
		}
		service.WriteJSON(w, res)
		
	default:
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
		service.WriteJSON(w, res)
	}
}

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	// 初始化数据库表
	if err := initTables(); err != nil {
		panic(fmt.Sprintf("init tables failed with %+v", err))
	}

	// 认证路由（必须在其他API路由之前）
	http.HandleFunc("/api/auth/", service.AuthHandler)
	
	// 静态文件服务（登录页面不需要认证）
	http.HandleFunc("/login", service.StaticHandler)
	http.HandleFunc("/", service.StaticHandler)
	http.HandleFunc("/admin", service.AdminHandler)
	
	// 静态文件服务 - 上传的图片等
	http.HandleFunc("/uploads/", func(w http.ResponseWriter, r *http.Request) {
		// 设置CORS头
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		// 移除 /uploads/ 前缀
		filePath := strings.TrimPrefix(r.URL.Path, "/uploads/")
		if filePath == "" {
			http.NotFound(w, r)
			return
		}
		
		// 构建完整文件路径
		fullPath := filepath.Join("uploads", filePath)
		
		// 安全检查：确保路径在uploads目录内
		if !strings.HasPrefix(fullPath, "uploads") {
			http.NotFound(w, r)
			return
		}
		
		// 检查文件是否存在
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
		
		// 根据文件扩展名设置Content-Type
		ext := strings.ToLower(filepath.Ext(filePath))
		switch ext {
		case ".jpg", ".jpeg":
			w.Header().Set("Content-Type", "image/jpeg")
		case ".png":
			w.Header().Set("Content-Type", "image/png")
		case ".gif":
			w.Header().Set("Content-Type", "image/gif")
		case ".bmp":
			w.Header().Set("Content-Type", "image/bmp")
		case ".webp":
			w.Header().Set("Content-Type", "image/webp")
		default:
			w.Header().Set("Content-Type", "application/octet-stream")
		}
		
		// 设置缓存头
		w.Header().Set("Cache-Control", "public, max-age=31536000")
		
		// 提供文件
		http.ServeFile(w, r, fullPath)
	})
	
	// 使用自定义RESTful路由处理器（需要认证）
	http.Handle("/api/", &RESTfulMux{})

	log.Printf("Server starting on port 80...")
	log.Fatal(http.ListenAndServe(":80", nil))
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
			return fmt.Errorf("failed to migrate %T: %v", model)
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
