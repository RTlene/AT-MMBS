package service

import (
	"fmt"
	"net/http"
	"time"

	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
)

// LoginRequest 登录请求结构
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse 登录响应结构
type LoginResponse struct {
	Token string      `json:"token"`
	User  *model.UserModel `json:"user"`
}

// AuthHandler 处理认证相关的请求
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	
	switch path {
	case "/api/auth/login":
		if r.Method == http.MethodPost {
			handleLogin(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	case "/api/auth/validate":
		if r.Method == http.MethodGet {
			handleValidateToken(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	case "/api/auth/logout":
		if r.Method == http.MethodPost {
			handleLogout(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	default:
		http.NotFound(w, r)
	}
}

// handleLogin 处理登录请求
func handleLogin(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginRequest
	if err := DecodeJSON(r, &loginReq); err != nil {
		WriteJSON(w, &JsonResult{
			Code:     -1,
			ErrorMsg: "请求参数解析失败",
		})
		return
	}

	// 验证用户名和密码
	if loginReq.Username == "" || loginReq.Password == "" {
		WriteJSON(w, &JsonResult{
			Code:     -1,
			ErrorMsg: "用户名和密码不能为空",
		})
		return
	}

	// 查询用户
	user, err := dao.Imp.GetUserByUsername(loginReq.Username)
	if err != nil {
		WriteJSON(w, &JsonResult{
			Code:     -1,
			ErrorMsg: "用户名或密码错误",
		})
		return
	}

	// 验证密码
	if !validatePassword(loginReq.Password, user.Password) {
		WriteJSON(w, &JsonResult{
			Code:     -1,
			ErrorMsg: "用户名或密码错误",
		})
		return
	}

	// 生成token
	token := generateToken(user.ID, user.Username)

	// 返回登录成功信息
	WriteJSON(w, &JsonResult{
		Code: 0,
		Data: LoginResponse{
			Token: token,
			User:  user,
		},
	})
}

// handleValidateToken 验证token
func handleValidateToken(w http.ResponseWriter, r *http.Request) {
	// 从请求头获取token
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		WriteJSON(w, &JsonResult{
			Code:     -1,
			ErrorMsg: "缺少认证token",
		})
		return
	}

	// 简单的token验证（实际项目中应该使用JWT等更安全的方式）
	if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
		WriteJSON(w, &JsonResult{
			Code:     -1,
			ErrorMsg: "token格式错误",
		})
		return
	}

	token := authHeader[7:]
	
	// 这里简化处理，实际应该验证token的有效性
	// 暂时只要token不为空就认为有效
	if token == "" {
		WriteJSON(w, &JsonResult{
			Code:     -1,
			ErrorMsg: "token无效",
		})
		return
	}

	WriteJSON(w, &JsonResult{
		Code: 0,
		Data: "token有效",
	})
}

// handleLogout 处理登出请求
func handleLogout(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, &JsonResult{
		Code: 0,
		Data: "登出成功",
	})
}

// validatePassword 验证密码
func validatePassword(inputPassword, storedPassword string) bool {
	// 对输入的密码进行MD5加密
	hashedInput := HashPassword(inputPassword)
	return hashedInput == storedPassword
}



// generateToken 生成简单的token（实际项目中应该使用JWT）
func generateToken(userID, username string) string {
	// 简单的token生成，实际项目中应该使用更安全的方式
	timestamp := time.Now().Unix()
	tokenData := fmt.Sprintf("%s_%s_%d", userID, username, timestamp)
	return HashPassword(tokenData)
}

// RequireAuth 中间件：要求认证
func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 检查认证头
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			// 如果没有认证头，检查是否是登录页面
			if r.URL.Path == "/login" || r.URL.Path == "/api/auth/login" {
				next.ServeHTTP(w, r)
				return
			}
			
			// 重定向到登录页面
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		// 验证token
		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		token := authHeader[7:]
		if token == "" {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		// token验证通过，继续处理请求
		next.ServeHTTP(w, r)
	}
}
