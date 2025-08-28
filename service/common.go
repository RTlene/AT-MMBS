package service

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
)

// JsonResult 通用JSON响应结构
type JsonResult struct {
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
	ErrorMsg  string      `json:"errorMsg,omitempty"`
}

// WriteJSON 写入JSON响应
func WriteJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}

// DecodeJSON 解析JSON请求体
func DecodeJSON(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

// GetUserByID 根据ID获取用户
func GetUserByID(userID string) (*model.UserModel, error) {
	return dao.Imp.GetUserByID(userID)
}

// UpdateUser 更新用户
func UpdateUser(user *model.UserModel) error {
	return dao.Imp.UpdateUser(user)
}

// UpdateUserPartial 部分更新用户字段
func UpdateUserPartial(userID string, updateData map[string]interface{}) error {
	return dao.Imp.UpdateUserPartial(userID, updateData)
}

// DeleteUser 删除用户
func DeleteUser(userID string) error {
	return dao.Imp.DeleteUser(userID)
}

// GetProductByID 根据ID获取商品
func GetProductByID(productID string) (*model.ProductModel, error) {
	return dao.Imp.GetProductByID(productID)
}

// UpdateProduct 更新商品
func UpdateProduct(product *model.ProductModel) error {
	return dao.Imp.UpdateProduct(product)
}

// DeleteProduct 删除商品
func DeleteProduct(productID string) error {
	return dao.Imp.DeleteProduct(productID)
}

// GetCategoryByID 根据ID获取分类
func GetCategoryByID(categoryID string) (*model.CategoryModel, error) {
	return dao.Imp.GetCategoryByID(categoryID)
}

// UpdateCategory 更新分类
func UpdateCategory(category *model.CategoryModel) error {
	return dao.Imp.UpdateCategory(category)
}

// DeleteCategory 删除分类
func DeleteCategory(categoryID string) error {
	return dao.Imp.DeleteCategory(categoryID)
}

// HashPassword 对密码进行MD5加密
func HashPassword(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}
