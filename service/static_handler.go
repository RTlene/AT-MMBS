package service

import (
	"fmt"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// StaticHandler 静态文件处理器
func StaticHandler(w http.ResponseWriter, r *http.Request) {
	// 获取请求的文件路径
	path := r.URL.Path
	
	// 如果路径是根路径，默认返回index.html
	if path == "/" {
		path = "/index.html"
	}
	
	// 如果路径是/admin，返回admin.html
	if path == "/admin" {
		path = "/admin.html"
	}
	
	// 如果路径是/login，返回login.html
	if path == "/login" {
		path = "/login.html"
	}
	
	// 移除开头的斜杠
	filePath := strings.TrimPrefix(path, "/")
	
	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}
	
	// 根据文件扩展名设置正确的Content-Type
	ext := filepath.Ext(filePath)
	switch ext {
	case ".html":
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	case ".css":
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
	case ".js":
		w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	case ".json":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	case ".png":
		w.Header().Set("Content-Type", "image/png")
	case ".jpg", ".jpeg":
		w.Header().Set("Content-Type", "image/jpeg")
	case ".gif":
		w.Header().Set("Content-Type", "image/gif")
	case ".svg":
		w.Header().Set("Content-Type", "image/svg+xml")
	default:
		// 尝试使用mime包检测类型
		if mimeType := mime.TypeByExtension(ext); mimeType != "" {
			w.Header().Set("Content-Type", mimeType)
		}
	}
	
	// 读取并返回文件内容
	content, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("读取文件失败: %v", err), http.StatusInternalServerError)
		return
	}
	
	w.Write(content)
}

// AdminHandler 管理后台接口（兼容性保留）
func AdminHandler(w http.ResponseWriter, r *http.Request) {
	// 直接调用StaticHandler处理/admin路径
	StaticHandler(w, r)
}

// IndexHandler 首页接口（兼容性保留）
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// 直接调用StaticHandler处理根路径
	StaticHandler(w, r)
}

// FileServerHandler 处理上传文件的访问
func FileServerHandler(w http.ResponseWriter, r *http.Request) {
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
}
