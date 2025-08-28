package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

type PaymentInterfaceImp struct{}

// CreatePayment 创建支付记录
func (imp *PaymentInterfaceImp) CreatePayment(payment *model.PaymentModel) error {
	cli := db.Get()
	return cli.Create(payment).Error
}

// UpdatePayment 更新支付记录
func (imp *PaymentInterfaceImp) UpdatePayment(payment *model.PaymentModel) error {
	cli := db.Get()
	return cli.Save(payment).Error
}

// GetPaymentByID 根据ID获取支付记录
func (imp *PaymentInterfaceImp) GetPaymentByID(id string) (*model.PaymentModel, error) {
	var payment model.PaymentModel
	cli := db.Get()
	err := cli.Where("id = ?", id).First(&payment).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

// GetPaymentByOrderID 根据订单ID获取支付记录
func (imp *PaymentInterfaceImp) GetPaymentByOrderID(orderID string) (*model.PaymentModel, error) {
	var payment model.PaymentModel
	cli := db.Get()
	err := cli.Where("order_id = ?", orderID).First(&payment).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

// GetPayments 分页获取支付记录列表
func (imp *PaymentInterfaceImp) GetPayments(page, pageSize int, memberID string) ([]*model.PaymentModel, int64, error) {
	var payments []*model.PaymentModel
	var total int64
	
	cli := db.Get()
	
	// 构建查询条件
	query := cli.Model(&model.PaymentModel{})
	if memberID != "" {
		query = query.Where("member_id = ?", memberID)
	}
	
	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Order("create_time desc").Find(&payments).Error; err != nil {
		return nil, 0, err
	}
	
	return payments, total, nil
}

// UpdatePaymentStatus 更新支付状态
func (imp *PaymentInterfaceImp) UpdatePaymentStatus(id string, status int) error {
	cli := db.Get()
	return cli.Model(&model.PaymentModel{}).Where("id = ?", id).Update("status", status).Error
}
