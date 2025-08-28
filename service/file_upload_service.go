package service

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// FileUploadHandler 处理文件上传
func FileUploadHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	if r.Method != http.MethodPost {
		res.Code = -1
		res.ErrorMsg = "只支持POST方法"
		WriteJSON(w, res)
		return
	}

	// 解析multipart表单
	if err := r.ParseMultipartForm(32 << 20); err != nil { // 32MB max
		res.Code = -1
		res.ErrorMsg = "文件解析失败: " + err.Error()
		WriteJSON(w, res)
		return
	}

	// 获取模块类型和产品ID
	module := r.FormValue("module")
	productID := r.FormValue("product_id")

	if module == "" {
		res.Code = -1
		res.ErrorMsg = "缺少模块类型参数"
		WriteJSON(w, res)
		return
	}

	// 创建上传目录
	uploadDir := fmt.Sprintf("uploads/%s", module)
	if productID != "" {
		uploadDir = fmt.Sprintf("%s/%s", uploadDir, productID)
	}

	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		res.Code = -1
		res.ErrorMsg = "创建目录失败: " + err.Error()
		WriteJSON(w, res)
		return
	}

	var uploadedFiles []string
	files := r.MultipartForm.File["files"]

	for _, fileHeader := range files {
		// 检查文件类型
		if !isValidImageFile(fileHeader.Filename) {
			continue
		}

		// 生成唯一文件名
		ext := filepath.Ext(fileHeader.Filename)
		filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
		filePath := filepath.Join(uploadDir, filename)

		// 打开源文件
		src, err := fileHeader.Open()
		if err != nil {
			continue
		}
		defer src.Close()

		// 创建目标文件
		dst, err := os.Create(filePath)
		if err != nil {
			src.Close()
			continue
		}
		defer dst.Close()

		// 复制文件内容
		if _, err := io.Copy(dst, src); err != nil {
			continue
		}

		// 记录上传的文件路径（相对于uploads目录）
		relativePath := strings.TrimPrefix(filePath, "uploads/")
		uploadedFiles = append(uploadedFiles, relativePath)
	}

	if len(uploadedFiles) == 0 {
		res.Code = -1
		res.ErrorMsg = "没有成功上传的文件"
	} else {
		res.Code = 0
		res.Data = map[string]interface{}{
			"files": uploadedFiles,
			"count": len(uploadedFiles),
		}
	}

	WriteJSON(w, res)
}

// isValidImageFile 检查是否为有效的图片文件
func isValidImageFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	validExts := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp"}
	
	for _, validExt := range validExts {
		if ext == validExt {
			return true
		}
	}
	return false
}

// DeleteFile 删除文件
func DeleteFile(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	if r.Method != http.MethodDelete {
		res.Code = -1
		res.ErrorMsg = "只支持DELETE方法"
		WriteJSON(w, res)
		return
	}

	filePath := r.URL.Query().Get("path")
	if filePath == "" {
		res.Code = -1
		res.ErrorMsg = "缺少文件路径参数"
		WriteJSON(w, res)
		return
	}

	// 构建完整文件路径
	fullPath := filepath.Join("uploads", filePath)
	
	// 安全检查：确保路径在uploads目录内
	if !strings.HasPrefix(fullPath, "uploads") {
		res.Code = -1
		res.ErrorMsg = "无效的文件路径"
		WriteJSON(w, res)
		return
	}

	// 删除文件
	if err := os.Remove(fullPath); err != nil {
		res.Code = -1
		res.ErrorMsg = "删除文件失败: " + err.Error()
	} else {
		res.Code = 0
		res.Data = "文件删除成功"
	}

	WriteJSON(w, res)
}
