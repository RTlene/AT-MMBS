package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	"wxcloudrun-golang/db/dao"

	"github.com/golang-jwt/jwt/v5"
)

// JWT配置
var (
	jwtSecret   = []byte(getJWTSecret())
	jwtIssuer   = "at-mmbs"
	jwtExpiry   = time.Hour * 24 // 24小时过期
)

// getJWTSecret 获取JWT密钥
func getJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "your-secret-key-change-this-in-production"
	}
	return secret
}

// JWTClaims 自定义JWT Claims
type JWTClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateJWT 生成JWT token
func GenerateJWT(userID, username, role string) (string, error) {
	now := time.Now()
	claims := JWTClaims{
		UserID:   userID,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    jwtIssuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(jwtExpiry)),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return tokenString, nil
}

// ValidateJWT 验证JWT token并返回claims
func ValidateJWT(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// handleJWTLogin 处理登录请求（使用JWT）
func handleJWTLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		WriteJSON(w, &JsonResult{
			Code:     -1,
			ErrorMsg: "只支持POST请求",
		})
		return
	}

	var loginReq LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
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

	// 生成JWT token
	token, err := GenerateJWT(user.ID, user.Username, user.Role)
	if err != nil {
		WriteJSON(w, &JsonResult{
			Code:     -1,
			ErrorMsg: "生成令牌失败",
		})
		return
	}

	// 返回登录成功信息
	WriteJSON(w, &JsonResult{
		Code: 0,
		Data: LoginResponse{
			Token: token,
			User:  user,
		},
	})
}

// handleJWTValidate 验证JWT token
func handleJWTValidate(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		WriteJSON(w, &JsonResult{
			Code:     -1,
			ErrorMsg: "未提供认证信息",
		})
		return
	}

	// 检查Bearer前缀
	if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
		WriteJSON(w, &JsonResult{
			Code:     -1,
			ErrorMsg: "token格式错误",
		})
		return
	}

	token := authHeader[7:]
	
	// 验证JWT token
	claims, err := ValidateJWT(token)
	if err != nil {
		WriteJSON(w, &JsonResult{
			Code:     -1,
			ErrorMsg: "token无效或已过期",
		})
		return
	}

	WriteJSON(w, &JsonResult{
		Code: 0,
		Data: map[string]interface{}{
			"user_id":   claims.UserID,
			"username":  claims.Username,
			"role":      claims.Role,
			"expires":   claims.ExpiresAt.Time.Format(time.RFC3339),
		},
	})
}

// ValidateTokenNew 新的JWT验证函数（替换原有的ValidateToken）
func ValidateTokenNew(token string) (string, error) {
	claims, err := ValidateJWT(token)
	if err != nil {
		return "", err
	}
	return claims.UserID, nil
}

// JWTMiddleware JWT认证中间件
func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 获取Authorization头
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			WriteUnauthorized(w, "未提供认证信息")
			return
		}
		
		// 检查Bearer token格式
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			WriteUnauthorized(w, "认证信息格式错误")
			return
		}
		
		token := parts[1]
		
		// 验证JWT token
		claims, err := ValidateJWT(token)
		if err != nil {
			WriteUnauthorized(w, "认证失败: "+err.Error())
			return
		}
		
		// 将用户信息添加到请求上下文
		r.Header.Set("X-User-ID", claims.UserID)
		r.Header.Set("X-User-Role", claims.Role)
		r.Header.Set("X-Username", claims.Username)
		
		// 继续处理请求
		next(w, r)
	}
}

// JWTAdminMiddleware JWT管理员权限中间件
func JWTAdminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return JWTMiddleware(func(w http.ResponseWriter, r *http.Request) {
		role := r.Header.Get("X-User-Role")
		
		if role != "admin" {
			WriteForbidden(w, "需要管理员权限")
			return
		}
		
		next(w, r)
	})
}