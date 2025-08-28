package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

type OrderInterfaceImp struct{}

// CreateOrder 创建订单
func (imp *OrderInterfaceImp) CreateOrder(order *model.OrderModel) error {
	cli := db.Get()
	return cli.Create(order).Error
}

// UpdateOrder 更新订单
func (imp *OrderInterfaceImp) UpdateOrder(order *model.OrderModel) error {
	cli := db.Get()
	return cli.Save(order).Error
}

// DeleteOrder 删除订单
func (imp *OrderInterfaceImp) DeleteOrder(id string) error {
	cli := db.Get()
	return cli.Delete(&model.OrderModel{}, "id = ?", id).Error
}

// GetOrderByID 根据ID获取订单
func (imp *OrderInterfaceImp) GetOrderByID(id string) (*model.OrderModel, error) {
	var order model.OrderModel
	cli := db.Get()
	err := cli.Where("id = ?", id).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// GetOrderByOrderNo 根据订单号获取订单
func (imp *OrderInterfaceImp) GetOrderByOrderNo(orderNo string) (*model.OrderModel, error) {
	var order model.OrderModel
	cli := db.Get()
	err := cli.Where("order_no = ?", orderNo).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// GetOrders 分页获取订单列表
func (imp *OrderInterfaceImp) GetOrders(page, pageSize int, memberID string) ([]*model.OrderModel, int64, error) {
	var orders []*model.OrderModel
	var total int64
	
	cli := db.Get()
	
	// 构建查询条件
	query := cli.Model(&model.OrderModel{})
	if memberID != "" {
		query = query.Where("member_id = ?", memberID)
	}
	
	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Order("create_time desc").Find(&orders).Error; err != nil {
		return nil, 0, err
	}
	
	return orders, total, nil
}

// UpdateOrderStatus 更新订单状态
func (imp *OrderInterfaceImp) UpdateOrderStatus(id string, status int) error {
	cli := db.Get()
	return cli.Model(&model.OrderModel{}).Where("id = ?", id).Update("order_status", status).Error
}
