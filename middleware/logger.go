package middleware

import (
	"log"
	"net/http"
	"time"
)

// LogRequest 记录请求日志
func LogRequest(r *http.Request) {
	log.Printf("[%s] %s %s %s", 
		time.Now().Format("2006-01-02 15:04:05"),
		r.Method,
		r.URL.Path,
		r.RemoteAddr,
	)
}

// LoggerMiddleware 日志中间件
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// 创建响应记录器
		rw := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		
		// 处理请求
		next.ServeHTTP(rw, r)
		
		// 记录响应日志
		duration := time.Since(start)
		log.Printf("[%s] %s %s %d %v",
			time.Now().Format("2006-01-02 15:04:05"),
			r.Method,
			r.URL.Path,
			rw.statusCode,
			duration,
		)
	})
}

// responseWriter 包装http.ResponseWriter以记录状态码
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}