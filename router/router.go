package router

import (
	"net/http"
	"wxcloudrun-golang/middleware"
	"wxcloudrun-golang/service"
)

// Router 路由器结构
type Router struct {
	mux *http.ServeMux
}

// NewRouter 创建新的路由器
func NewRouter() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

// SetupRoutes 设置所有路由
func (r *Router) SetupRoutes() {
	// 静态文件和页面（不需要认证）
	r.mux.HandleFunc("/", service.StaticHandler)
	r.mux.HandleFunc("/login", service.StaticHandler)
	r.mux.HandleFunc("/admin", middleware.AuthMiddleware(service.AdminHandler))
	
	// 认证相关路由（不需要认证中间件）
	r.mux.HandleFunc("/api/auth/", service.AuthHandler)
	
	// 文件上传路由（需要认证）
	r.mux.HandleFunc("/uploads/", service.FileServerHandler)
	
	// API路由（需要认证）
	r.setupAPIRoutes()
}

// setupAPIRoutes 设置API路由
func (r *Router) setupAPIRoutes() {
	// 使用中间件包装需要认证的路由
	authWrapper := func(handler http.HandlerFunc) http.HandlerFunc {
		return middleware.AuthMiddleware(handler)
	}
	
	// 计数器
	r.mux.HandleFunc("/api/count", authWrapper(service.CounterHandler))
	
	// 用户管理
	r.mux.HandleFunc("/api/users", authWrapper(service.UserHandler))
	r.mux.HandleFunc("/api/users/", authWrapper(service.UserRESTHandler))
	
	// 商品管理
	r.mux.HandleFunc("/api/products", authWrapper(service.ProductHandler))
	r.mux.HandleFunc("/api/products/", authWrapper(service.ProductRESTHandler))
	r.mux.HandleFunc("/api/products/status", authWrapper(service.ProductStatusHandler))
	
	// 分类管理
	r.mux.HandleFunc("/api/categories", authWrapper(service.CategoryHandler))
	r.mux.HandleFunc("/api/categories/", authWrapper(service.CategoryRESTHandler))
	r.mux.HandleFunc("/api/categories/parent", authWrapper(service.ParentCategoryHandler))
	
	// 会员管理
	r.mux.HandleFunc("/api/members", authWrapper(service.MemberHandler))
	r.mux.HandleFunc("/api/members/", authWrapper(service.MemberRESTHandler))
	r.mux.HandleFunc("/api/members/network", authWrapper(service.MemberNetworkHandler))
	r.mux.HandleFunc("/api/members/distribution", authWrapper(service.MemberDistributionHandler))
	r.mux.HandleFunc("/api/members/consumption", authWrapper(service.MemberConsumptionHandler))
	
	// 会员等级
	r.mux.HandleFunc("/api/member-levels", authWrapper(service.MemberLevelHandler))
	r.mux.HandleFunc("/api/member-levels/", authWrapper(service.MemberLevelRESTHandler))
	
	// 分销商等级
	r.mux.HandleFunc("/api/distributor-levels", authWrapper(service.DistributorLevelHandler))
	r.mux.HandleFunc("/api/distributor-levels/", authWrapper(service.DistributorLevelRESTHandler))
	
	// 订单管理
	r.mux.HandleFunc("/api/orders", authWrapper(service.OrderHandler))
	r.mux.HandleFunc("/api/orders/", authWrapper(service.OrderRESTHandler))
	r.mux.HandleFunc("/api/orders/status", authWrapper(service.OrderStatusHandler))
	
	// 支付管理
	r.mux.HandleFunc("/api/payments", authWrapper(service.PaymentHandler))
	r.mux.HandleFunc("/api/payments/", authWrapper(service.PaymentRESTHandler))
	r.mux.HandleFunc("/api/payments/status", authWrapper(service.PaymentStatusHandler))
	
	// 文件上传
	r.mux.HandleFunc("/api/upload", authWrapper(service.FileUploadHandler))
	r.mux.HandleFunc("/api/upload/delete", authWrapper(service.DeleteFile))
}

// ServeHTTP 实现http.Handler接口
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 添加CORS头
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	
	// 处理OPTIONS请求
	if req.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	
	// 记录请求日志
	middleware.LogRequest(req)
	
	// 处理请求
	r.mux.ServeHTTP(w, req)
}