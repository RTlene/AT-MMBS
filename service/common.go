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

// WriteSuccess 写入成功响应
func WriteSuccess(w http.ResponseWriter, data interface{}) {
	WriteJSON(w, &JsonResult{
		Code: 0,
		Data: data,
	})
}

// WriteError 写入错误响应
func WriteError(w http.ResponseWriter, code int, errorMsg string) {
	WriteJSON(w, &JsonResult{
		Code:     code,
		ErrorMsg: errorMsg,
	})
}

// WriteUnauthorized 写入未授权响应
func WriteUnauthorized(w http.ResponseWriter, errorMsg string) {
	w.WriteHeader(http.StatusUnauthorized)
	WriteJSON(w, &JsonResult{
		Code:     401,
		ErrorMsg: errorMsg,
	})
}

// WriteForbidden 写入禁止访问响应
func WriteForbidden(w http.ResponseWriter, errorMsg string) {
	w.WriteHeader(http.StatusForbidden)
	WriteJSON(w, &JsonResult{
		Code:     403,
		ErrorMsg: errorMsg,
	})
}

// GetMemberByID 根据ID获取会员
func GetMemberByID(memberID string) (*model.MemberModel, error) {
	return dao.Imp.GetMemberByID(memberID)
}

// UpdateMember 更新会员
func UpdateMember(member *model.MemberModel) error {
	return dao.Imp.UpdateMember(member)
}

// DeleteMember 删除会员
func DeleteMember(memberID string) error {
	return dao.Imp.DeleteMember(memberID)
}

// GetMemberLevelByID 根据ID获取会员等级
func GetMemberLevelByID(levelID string) (*model.MemberLevelModel, error) {
	return dao.Imp.GetMemberLevelByID(levelID)
}

// UpdateMemberLevel 更新会员等级
func UpdateMemberLevel(level *model.MemberLevelModel) error {
	return dao.Imp.UpdateMemberLevel(level)
}

// DeleteMemberLevel 删除会员等级
func DeleteMemberLevel(levelID string) error {
	return dao.Imp.DeleteMemberLevel(levelID)
}

// GetDistributorLevelByID 根据ID获取分销商等级
func GetDistributorLevelByID(levelID string) (*model.DistributorLevelModel, error) {
	return dao.Imp.GetDistributorLevelByID(levelID)
}

// UpdateDistributorLevel 更新分销商等级
func UpdateDistributorLevel(level *model.DistributorLevelModel) error {
	return dao.Imp.UpdateDistributorLevel(level)
}

// DeleteDistributorLevel 删除分销商等级
func DeleteDistributorLevel(levelID string) error {
	return dao.Imp.DeleteDistributorLevel(levelID)
}

// GetOrderByID 根据ID获取订单
func GetOrderByID(orderID string) (*model.OrderModel, error) {
	return dao.Imp.GetOrderByID(orderID)
}

// UpdateOrder 更新订单
func UpdateOrder(order *model.OrderModel) error {
	return dao.Imp.UpdateOrder(order)
}

// DeleteOrder 删除订单
func DeleteOrder(orderID string) error {
	return dao.Imp.DeleteOrder(orderID)
}

// GetPaymentByID 根据ID获取支付记录
func GetPaymentByID(paymentID string) (*model.PaymentModel, error) {
	return dao.Imp.GetPaymentByID(paymentID)
}

// UpdatePayment 更新支付记录
func UpdatePayment(payment *model.PaymentModel) error {
	return dao.Imp.UpdatePayment(payment)
}

// DeletePayment 删除支付记录
func DeletePayment(paymentID string) error {
	return dao.Imp.DeletePayment(paymentID)
}
