package model

import "time"

// OrderModel 订单模型
type OrderModel struct {
	ID           string    `gorm:"column:id;primaryKey;type:varchar(36)" json:"id"`
	OrderNo      string    `gorm:"column:order_no;type:varchar(50);not null;uniqueIndex" json:"order_no"`
	MemberID     string    `gorm:"column:member_id;type:varchar(36);not null" json:"member_id"`
	Products     []string  `gorm:"column:products;type:json;not null;comment:'订购商品列表'" json:"products"`
	DeliveryType string    `gorm:"column:delivery_type;type:varchar(20);not null;comment:'配送方式'" json:"delivery_type"`
	Address      string    `gorm:"column:address;type:text;comment:'配送/门店地址'" json:"address"`
	TotalAmount  float64   `gorm:"column:total_amount;type:decimal(10,2);not null;comment:'订单金额'" json:"total_amount"`
	PaidAmount   float64   `gorm:"column:paid_amount;type:decimal(10,2);not null;comment:'实付金额'" json:"paid_amount"`
	PaymentStatus int      `gorm:"column:payment_status;type:int;default:0;comment:'0:未支付 1:已支付 2:已退款'" json:"payment_status"`
	OrderStatus  int      `gorm:"column:order_status;type:int;default:0;comment:'0:待确认 1:已确认 2:已发货 3:已完成 4:已取消'" json:"order_status"`
	CreateTime   time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`
	UpdateTime   time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time"`
}

func (OrderModel) TableName() string {
	return "orders"
}
