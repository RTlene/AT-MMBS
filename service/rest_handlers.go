package service

import (
	"fmt"
	"net/http"
	"strings"
	"wxcloudrun-golang/db/model"
)

// UserRESTHandler 处理单个用户的RESTful操作
func UserRESTHandler(w http.ResponseWriter, r *http.Request) {
	// 从URL中提取用户ID
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		WriteError(w, -1, "无效的请求路径")
		return
	}
	
	userID := parts[3]
	if userID == "" {
		WriteError(w, -1, "缺少用户ID")
		return
	}
	
	switch r.Method {
	case http.MethodGet:
		handleGetUser(w, r, userID)
	case http.MethodPut:
		handleUpdateUser(w, r, userID)
	case http.MethodDelete:
		handleDeleteUser(w, r, userID)
	default:
		WriteError(w, -1, fmt.Sprintf("不支持的方法: %s", r.Method))
	}
}

// ProductRESTHandler 处理单个商品的RESTful操作
func ProductRESTHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		WriteError(w, -1, "无效的请求路径")
		return
	}
	
	productID := parts[3]
	if productID == "" {
		WriteError(w, -1, "缺少商品ID")
		return
	}
	
	switch r.Method {
	case http.MethodGet:
		handleGetProduct(w, r, productID)
	case http.MethodPut:
		handleUpdateProduct(w, r, productID)
	case http.MethodDelete:
		handleDeleteProduct(w, r, productID)
	default:
		WriteError(w, -1, fmt.Sprintf("不支持的方法: %s", r.Method))
	}
}

// CategoryRESTHandler 处理单个分类的RESTful操作
func CategoryRESTHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		WriteError(w, -1, "无效的请求路径")
		return
	}
	
	categoryID := parts[3]
	if categoryID == "" {
		WriteError(w, -1, "缺少分类ID")
		return
	}
	
	switch r.Method {
	case http.MethodGet:
		handleGetCategory(w, r, categoryID)
	case http.MethodPut:
		handleUpdateCategory(w, r, categoryID)
	case http.MethodDelete:
		handleDeleteCategory(w, r, categoryID)
	default:
		WriteError(w, -1, fmt.Sprintf("不支持的方法: %s", r.Method))
	}
}

// MemberRESTHandler 处理单个会员的RESTful操作
func MemberRESTHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		WriteError(w, -1, "无效的请求路径")
		return
	}
	
	memberID := parts[3]
	if memberID == "" {
		WriteError(w, -1, "缺少会员ID")
		return
	}
	
	switch r.Method {
	case http.MethodGet:
		handleGetMember(w, r, memberID)
	case http.MethodPut:
		handleUpdateMember(w, r, memberID)
	case http.MethodDelete:
		handleDeleteMember(w, r, memberID)
	default:
		WriteError(w, -1, fmt.Sprintf("不支持的方法: %s", r.Method))
	}
}

// MemberLevelRESTHandler 处理单个会员等级的RESTful操作
func MemberLevelRESTHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		WriteError(w, -1, "无效的请求路径")
		return
	}
	
	levelID := parts[3]
	if levelID == "" {
		WriteError(w, -1, "缺少等级ID")
		return
	}
	
	switch r.Method {
	case http.MethodGet:
		handleGetMemberLevel(w, r, levelID)
	case http.MethodPut:
		handleUpdateMemberLevel(w, r, levelID)
	case http.MethodDelete:
		handleDeleteMemberLevel(w, r, levelID)
	default:
		WriteError(w, -1, fmt.Sprintf("不支持的方法: %s", r.Method))
	}
}

// DistributorLevelRESTHandler 处理单个分销商等级的RESTful操作
func DistributorLevelRESTHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		WriteError(w, -1, "无效的请求路径")
		return
	}
	
	levelID := parts[3]
	if levelID == "" {
		WriteError(w, -1, "缺少等级ID")
		return
	}
	
	switch r.Method {
	case http.MethodGet:
		handleGetDistributorLevel(w, r, levelID)
	case http.MethodPut:
		handleUpdateDistributorLevel(w, r, levelID)
	case http.MethodDelete:
		handleDeleteDistributorLevel(w, r, levelID)
	default:
		WriteError(w, -1, fmt.Sprintf("不支持的方法: %s", r.Method))
	}
}

// OrderRESTHandler 处理单个订单的RESTful操作
func OrderRESTHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		WriteError(w, -1, "无效的请求路径")
		return
	}
	
	orderID := parts[3]
	if orderID == "" {
		WriteError(w, -1, "缺少订单ID")
		return
	}
	
	switch r.Method {
	case http.MethodGet:
		handleGetOrder(w, r, orderID)
	case http.MethodPut:
		handleUpdateOrder(w, r, orderID)
	case http.MethodDelete:
		handleDeleteOrder(w, r, orderID)
	default:
		WriteError(w, -1, fmt.Sprintf("不支持的方法: %s", r.Method))
	}
}

// PaymentRESTHandler 处理单个支付的RESTful操作
func PaymentRESTHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		WriteError(w, -1, "无效的请求路径")
		return
	}
	
	paymentID := parts[3]
	if paymentID == "" {
		WriteError(w, -1, "缺少支付ID")
		return
	}
	
	switch r.Method {
	case http.MethodGet:
		handleGetPayment(w, r, paymentID)
	case http.MethodPut:
		handleUpdatePayment(w, r, paymentID)
	case http.MethodDelete:
		handleDeletePayment(w, r, paymentID)
	default:
		WriteError(w, -1, fmt.Sprintf("不支持的方法: %s", r.Method))
	}
}

// 具体的处理函数
func handleGetUser(w http.ResponseWriter, r *http.Request, userID string) {
	user, err := GetUserByID(userID)
	if err != nil {
		WriteError(w, -1, err.Error())
		return
	}
	WriteSuccess(w, user)
}

func handleUpdateUser(w http.ResponseWriter, r *http.Request, userID string) {
	var userData map[string]interface{}
	if err := DecodeJSON(r, &userData); err != nil {
		WriteError(w, -1, "请求参数解析失败")
		return
	}
	
	// 只更新提供的字段
	updateData := make(map[string]interface{})
	if username, ok := userData["username"].(string); ok && username != "" {
		updateData["username"] = username
	}
	if password, ok := userData["password"].(string); ok && password != "" {
		updateData["password"] = HashPassword(password)
	}
	if role, ok := userData["role"].(string); ok && role != "" {
		updateData["role"] = role
	}
	
	if len(updateData) == 0 {
		WriteError(w, -1, "没有提供要更新的字段")
		return
	}
	
	updateData["id"] = userID
	if err := UpdateUserPartial(userID, updateData); err != nil {
		WriteError(w, -1, err.Error())
		return
	}
	
	WriteSuccess(w, "用户更新成功")
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request, userID string) {
	if err := DeleteUser(userID); err != nil {
		WriteError(w, -1, err.Error())
		return
	}
	WriteSuccess(w, "删除成功")
}

func handleGetProduct(w http.ResponseWriter, r *http.Request, productID string) {
	product, err := GetProductByID(productID)
	if err != nil {
		WriteError(w, -1, err.Error())
		return
	}
	WriteSuccess(w, product)
}

func handleUpdateProduct(w http.ResponseWriter, r *http.Request, productID string) {
	var product model.ProductModel
	if err := DecodeJSON(r, &product); err != nil {
		WriteError(w, -1, "请求参数解析失败")
		return
	}
	
	product.ID = productID
	if err := UpdateProduct(&product); err != nil {
		WriteError(w, -1, err.Error())
		return
	}
	
	WriteSuccess(w, product)
}

func handleDeleteProduct(w http.ResponseWriter, r *http.Request, productID string) {
	if err := DeleteProduct(productID); err != nil {
		WriteError(w, -1, err.Error())
		return
	}
	WriteSuccess(w, "删除成功")
}

func handleGetCategory(w http.ResponseWriter, r *http.Request, categoryID string) {
	category, err := GetCategoryByID(categoryID)
	if err != nil {
		WriteError(w, -1, err.Error())
		return
	}
	WriteSuccess(w, category)
}

func handleUpdateCategory(w http.ResponseWriter, r *http.Request, categoryID string) {
	var category model.CategoryModel
	if err := DecodeJSON(r, &category); err != nil {
		WriteError(w, -1, "请求参数解析失败")
		return
	}
	
	category.ID = categoryID
	if err := UpdateCategory(&category); err != nil {
		WriteError(w, -1, err.Error())
		return
	}
	
	WriteSuccess(w, category)
}

func handleDeleteCategory(w http.ResponseWriter, r *http.Request, categoryID string) {
	if err := DeleteCategory(categoryID); err != nil {
		WriteError(w, -1, err.Error())
		return
	}
	WriteSuccess(w, "删除成功")
}

func handleGetMember(w http.ResponseWriter, r *http.Request, memberID string) {
	member, err := GetMemberByID(memberID)
	if err != nil {
		WriteError(w, -1, err.Error())
		return
	}
	WriteSuccess(w, member)
}

func handleUpdateMember(w http.ResponseWriter, r *http.Request, memberID string) {
	var member model.MemberModel
	if err := DecodeJSON(r, &member); err != nil {
		WriteError(w, -1, "请求参数解析失败")
		return
	}
	
	member.ID = memberID
	if err := UpdateMember(&member); err != nil {
		WriteError(w, -1, err.Error())
		return
	}
	
	WriteSuccess(w, member)
}

func handleDeleteMember(w http.ResponseWriter, r *http.Request, memberID string) {
	if err := DeleteMember(memberID); err != nil {
		WriteError(w, -1, err.Error())
		return
	}
	WriteSuccess(w, "删除成功")
}

func handleGetMemberLevel(w http.ResponseWriter, r *http.Request, levelID string) {
	level, err := GetMemberLevelByID(levelID)
	if err != nil {
		WriteError(w, -1, err.Error())
		return
	}
	WriteSuccess(w, level)
}

func handleUpdateMemberLevel(w http.ResponseWriter, r *http.Request, levelID string) {
	var level model.MemberLevelModel
	if err := DecodeJSON(r, &level); err != nil {
		WriteError(w, -1, "请求参数解析失败")
		return
	}
	
	level.ID = levelID
	if err := UpdateMemberLevel(&level); err != nil {
		WriteError(w, -1, err.Error())
		return
	}
	
	WriteSuccess(w, level)
}

func handleDeleteMemberLevel(w http.ResponseWriter, r *http.Request, levelID string) {
	if err := DeleteMemberLevel(levelID); err != nil {
		WriteError(w, -1, err.Error())
		return
	}
	WriteSuccess(w, "删除成功")
}

func handleGetDistributorLevel(w http.ResponseWriter, r *http.Request, levelID string) {
	level, err := GetDistributorLevelByID(levelID)
	if err != nil {
		WriteError(w, -1, err.Error())
		return
	}
	WriteSuccess(w, level)
}

func handleUpdateDistributorLevel(w http.ResponseWriter, r *http.Request, levelID string) {
	var level model.DistributorLevelModel
	if err := DecodeJSON(r, &level); err != nil {
		WriteError(w, -1, "请求参数解析失败")
		return
	}
	
	level.ID = levelID
	if err := UpdateDistributorLevel(&level); err != nil {
		WriteError(w, -1, err.Error())
		return
	}
	
	WriteSuccess(w, level)
}

func handleDeleteDistributorLevel(w http.ResponseWriter, r *http.Request, levelID string) {
	if err := DeleteDistributorLevel(levelID); err != nil {
		WriteError(w, -1, err.Error())
		return
	}
	WriteSuccess(w, "删除成功")
}

func handleGetOrder(w http.ResponseWriter, r *http.Request, orderID string) {
	order, err := GetOrderByID(orderID)
	if err != nil {
		WriteError(w, -1, err.Error())
		return
	}
	WriteSuccess(w, order)
}

func handleUpdateOrder(w http.ResponseWriter, r *http.Request, orderID string) {
	var order model.OrderModel
	if err := DecodeJSON(r, &order); err != nil {
		WriteError(w, -1, "请求参数解析失败")
		return
	}
	
	order.ID = orderID
	if err := UpdateOrder(&order); err != nil {
		WriteError(w, -1, err.Error())
		return
	}
	
	WriteSuccess(w, order)
}

func handleDeleteOrder(w http.ResponseWriter, r *http.Request, orderID string) {
	if err := DeleteOrder(orderID); err != nil {
		WriteError(w, -1, err.Error())
		return
	}
	WriteSuccess(w, "删除成功")
}

func handleGetPayment(w http.ResponseWriter, r *http.Request, paymentID string) {
	payment, err := GetPaymentByID(paymentID)
	if err != nil {
		WriteError(w, -1, err.Error())
		return
	}
	WriteSuccess(w, payment)
}

func handleUpdatePayment(w http.ResponseWriter, r *http.Request, paymentID string) {
	var payment model.PaymentModel
	if err := DecodeJSON(r, &payment); err != nil {
		WriteError(w, -1, "请求参数解析失败")
		return
	}
	
	payment.ID = paymentID
	if err := UpdatePayment(&payment); err != nil {
		WriteError(w, -1, err.Error())
		return
	}
	
	WriteSuccess(w, payment)
}

func handleDeletePayment(w http.ResponseWriter, r *http.Request, paymentID string) {
	if err := DeletePayment(paymentID); err != nil {
		WriteError(w, -1, err.Error())
		return
	}
	WriteSuccess(w, "删除成功")
}