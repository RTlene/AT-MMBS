package model

import "time"

// PaymentModel 支付模型
type PaymentModel struct {
	ID            string    `gorm:"column:id;primaryKey;type:varchar(36)" json:"id"`
	OrderID       string    `gorm:"column:order_id;type:varchar(36);not null;uniqueIndex" json:"order_id"`
	OrderNo       string    `gorm:"column:order_no;type:varchar(50);not null" json:"order_no"`
	MemberID      string    `gorm:"column:member_id;type:varchar(36);not null" json:"member_id"`
	PaymentMethod string    `gorm:"column:payment_method;type:varchar(20);not null;comment:'支付方式'" json:"payment_method"`
	Amount        float64   `gorm:"column:amount;type:decimal(10,2);not null" json:"amount"`
	Status        int       `gorm:"column:status;type:int;default:0;comment:'0:待支付 1:支付中 2:支付成功 3:支付失败 4:已退款'" json:"status"`
	TransactionID string    `gorm:"column:transaction_id;type:varchar(100);comment:'第三方支付交易号'" json:"transaction_id"`
	PayTime       time.Time `gorm:"column:pay_time" json:"pay_time"`
	RefundTime    time.Time `gorm:"column:refund_time" json:"refund_time"`
	RefundAmount  float64   `gorm:"column:refund_amount;type:decimal(10,2);default:0" json:"refund_amount"`
	CreateTime    time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`
	UpdateTime    time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time"`
}

func (PaymentModel) TableName() string {
	return "payments"
}
