package middleware

import (
	"net/http"
	"strings"
	"wxcloudrun-golang/service"
)

// AuthMiddleware 认证中间件
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 获取Authorization头
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			service.WriteUnauthorized(w, "未提供认证信息")
			return
		}
		
		// 检查Bearer token格式
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			service.WriteUnauthorized(w, "认证信息格式错误")
			return
		}
		
		token := parts[1]
		
		// 验证token
		userID, err := service.ValidateToken(token)
		if err != nil {
			service.WriteUnauthorized(w, "认证失败: "+err.Error())
			return
		}
		
		// 将用户ID添加到请求上下文
		r.Header.Set("X-User-ID", userID)
		
		// 继续处理请求
		next(w, r)
	}
}

// AdminMiddleware 管理员权限中间件
func AdminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("X-User-ID")
		
		// 检查用户是否为管理员
		isAdmin, err := service.IsUserAdmin(userID)
		if err != nil || !isAdmin {
			service.WriteForbidden(w, "需要管理员权限")
			return
		}
		
		next(w, r)
	})
}