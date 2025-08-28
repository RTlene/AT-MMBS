package dao

import (
	"fmt"
	"time"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

type ProductInterfaceImp struct{}

// CreateProduct 创建商品
func (imp *ProductInterfaceImp) CreateProduct(product *model.ProductModel) error {
	cli := db.Get()
	return cli.Create(product).Error
}

// UpdateProduct 更新商品
func (imp *ProductInterfaceImp) UpdateProduct(product *model.ProductModel) error {
	cli := db.Get()
	
	// 确保update_time不为零值
	if product.UpdateTime.IsZero() {
		product.UpdateTime = time.Now()
	}
	
	// 只更新特定字段，避免覆盖create_time
	return cli.Model(&model.ProductModel{}).
		Where("_id = ?", product.ID).
		Updates(map[string]interface{}{
			"name":           product.Name,
			"category":       product.Category,
			"price":          product.Price,
			"kucun":          product.Kucun,
			"tp":             product.Tp,
			"sp":             product.Sp,
			"overlay_style":  product.OverlayStyle,
			"content":        product.Content,
			"status":         product.Status,
			"update_time":    product.UpdateTime,
		}).Error
}

// DeleteProduct 删除商品
func (imp *ProductInterfaceImp) DeleteProduct(id string) error {
	cli := db.Get()
	return cli.Delete(&model.ProductModel{}, "_id = ?", id).Error
}

// GetProductByID 根据ID获取商品
func (imp *ProductInterfaceImp) GetProductByID(id string) (*model.ProductModel, error) {
	var product model.ProductModel
	cli := db.Get()
	err := cli.Where("_id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// GetProducts 分页获取商品列表
func (imp *ProductInterfaceImp) GetProducts(page, pageSize int, category string) ([]*model.ProductModel, int64, error) {
	var products []*model.ProductModel
	var total int64
	
	cli := db.Get()
	
	// 构建查询条件
	query := cli.Model(&model.ProductModel{})
	if category != "" {
		query = query.Where("category = ?", category)
	}
	
	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Order("create_time desc").Find(&products).Error; err != nil {
		return nil, 0, err
	}
	
	return products, total, nil
}

// UpdateProductStatus 更新商品上下架状态
func (imp *ProductInterfaceImp) UpdateProductStatus(id string, status int) error {
	cli := db.Get()
	
	// 添加调试日志
	fmt.Printf("开始更新商品状态，ID: %s, 新状态: %d\n", id, status)
	
	// 先检查商品是否存在
	var existingProduct model.ProductModel
	err := cli.Where("_id = ?", id).First(&existingProduct).Error
	if err != nil {
		fmt.Printf("商品不存在，ID: %s, 错误: %v\n", id, err)
		return err
	}
	
	fmt.Printf("找到商品，当前状态: %d\n", existingProduct.Status)
	
	// 使用原生SQL更新，确保更新成功
	updateSQL := "UPDATE products SET status = ?, update_time = ? WHERE _id = ?"
	result := cli.Exec(updateSQL, status, time.Now(), id)
	
	if result.Error != nil {
		fmt.Printf("更新失败，错误: %v\n", result.Error)
		return result.Error
	}
	
	// 检查更新影响的行数
	if result.RowsAffected == 0 {
		fmt.Printf("更新影响行数为0，可能没有实际更新\n")
		return fmt.Errorf("没有记录被更新")
	}
	
	fmt.Printf("状态更新成功，影响行数: %d\n", result.RowsAffected)
	
	// 验证更新结果
	var updatedProduct model.ProductModel
	err = cli.Where("_id = ?", id).First(&updatedProduct).Error
	if err != nil {
		fmt.Printf("验证更新结果失败，错误: %v\n", err)
		return err
	}
	
	fmt.Printf("验证结果，更新后状态: %d\n", updatedProduct.Status)
	
	return nil
}
