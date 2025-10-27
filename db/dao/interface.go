package dao

import (
	"wxcloudrun-golang/db/model"
)

// CounterInterface 计数器接口
type CounterInterface interface {
	ClearCounter(id int32) error
	UpsertCounter(counter *model.CounterModel) error
	GetCounter(id int32) (*model.CounterModel, error)
}

// UserInterface 用户接口
type UserInterface interface {
	CreateUser(user *model.UserModel) error
	UpdateUser(user *model.UserModel) error
	UpdateUserPartial(id string, updateData map[string]interface{}) error
	DeleteUser(id string) error
	GetUserByID(id string) (*model.UserModel, error)
	GetUserByUsername(username string) (*model.UserModel, error)
	GetUsers(page, pageSize int) ([]*model.UserModel, int64, error)
}

// ProductInterface 商品接口
type ProductInterface interface {
	CreateProduct(product *model.ProductModel) error
	UpdateProduct(product *model.ProductModel) error
	DeleteProduct(id string) error
	GetProductByID(id string) (*model.ProductModel, error)
	GetProducts(page, pageSize int, category string) ([]*model.ProductModel, int64, error)
	UpdateProductStatus(id string, status int) error
}

// CategoryInterface 分类接口
type CategoryInterface interface {
	CreateCategory(category *model.CategoryModel) error
	UpdateCategory(category *model.CategoryModel) error
	DeleteCategory(id string) error
	GetCategoryByID(id string) (*model.CategoryModel, error)
	GetCategories() ([]*model.CategoryModel, error)
	GetParentCategories() ([]*model.CategoryModel, error)
}

// MemberInterface 会员接口
type MemberInterface interface {
	CreateMember(member *model.MemberModel) error
	UpdateMember(member *model.MemberModel) error
	DeleteMember(id string) error
	GetMemberByID(id string) (*model.MemberModel, error)
	GetMemberByOpenID(openid string) (*model.MemberModel, error)
	GetMembers(page, pageSize int) ([]*model.MemberModel, int64, error)
	GetMemberNetwork(id string) (map[string]interface{}, error)
	GetMemberDistribution(id string) (map[string]interface{}, error)
	GetMemberConsumption(id string) ([]*model.OrderModel, error)
}

// MemberLevelInterface 会员等级接口
type MemberLevelInterface interface {
	CreateMemberLevel(level *model.MemberLevelModel) error
	UpdateMemberLevel(level *model.MemberLevelModel) error
	DeleteMemberLevel(id string) error
	GetMemberLevelByID(id string) (*model.MemberLevelModel, error)
	GetMemberLevels() ([]*model.MemberLevelModel, error)
}

// DistributorLevelInterface 分销等级接口
type DistributorLevelInterface interface {
	CreateDistributorLevel(level *model.DistributorLevelModel) error
	UpdateDistributorLevel(level *model.DistributorLevelModel) error
	DeleteDistributorLevel(id string) error
	GetDistributorLevelByID(id string) (*model.DistributorLevelModel, error)
	GetDistributorLevels() ([]*model.DistributorLevelModel, error)
}

// OrderInterface 订单接口
type OrderInterface interface {
	CreateOrder(order *model.OrderModel) error
	UpdateOrder(order *model.OrderModel) error
	DeleteOrder(id string) error
	GetOrderByID(id string) (*model.OrderModel, error)
	GetOrderByOrderNo(orderNo string) (*model.OrderModel, error)
	GetOrders(page, pageSize int, memberID string) ([]*model.OrderModel, int64, error)
	UpdateOrderStatus(id string, status int) error
}

// PaymentInterface 支付接口
type PaymentInterface interface {
	CreatePayment(payment *model.PaymentModel) error
	UpdatePayment(payment *model.PaymentModel) error
	DeletePayment(id string) error
	GetPaymentByID(id string) (*model.PaymentModel, error)
	GetPaymentByOrderID(orderID string) (*model.PaymentModel, error)
	GetPayments(page, pageSize int, memberID string) ([]*model.PaymentModel, int64, error)
	UpdatePaymentStatus(id string, status int) error
}

// Imp 接口实现实例
var Imp = &InterfaceImp{}

// InterfaceImp 接口实现
type InterfaceImp struct {
	CounterInterface
	UserInterface
	ProductInterface
	CategoryInterface
	MemberInterface
	MemberLevelInterface
	DistributorLevelInterface
	OrderInterface
	PaymentInterface
}

func init() {
	Imp.CounterInterface = &CounterInterfaceImp{}
	Imp.UserInterface = &UserInterfaceImp{}
	Imp.ProductInterface = &ProductInterfaceImp{}
	Imp.CategoryInterface = &CategoryInterfaceImp{}
	Imp.MemberInterface = &MemberInterfaceImp{}
	Imp.MemberLevelInterface = &MemberLevelInterfaceImp{}
	Imp.DistributorLevelInterface = &DistributorLevelInterfaceImp{}
	Imp.OrderInterface = &OrderInterfaceImp{}
	Imp.PaymentInterface = &PaymentInterfaceImp{}
}
