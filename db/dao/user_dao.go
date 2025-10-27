package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

type UserInterfaceImp struct{}

// CreateUser 创建用户
func (imp *UserInterfaceImp) CreateUser(user *model.UserModel) error {
	cli := db.Get()
	return cli.Create(user).Error
}

// UpdateUser 更新用户
func (imp *UserInterfaceImp) UpdateUser(user *model.UserModel) error {
	cli := db.Get()
	// 使用Updates避免覆盖时间字段，只更新指定字段
	updateData := map[string]interface{}{
		"username": user.Username,
		"password": user.Password,
	}
	return cli.Model(&model.UserModel{}).Where("id = ?", user.ID).Updates(updateData).Error
}

// UpdateUserPartial 部分更新用户字段
func (imp *UserInterfaceImp) UpdateUserPartial(userID string, updateData map[string]interface{}) error {
	cli := db.Get()
	// 移除id字段，避免更新主键
	delete(updateData, "id")
	
	// 确保只更新允许的字段
	allowedFields := []string{"username", "password", "role"}
	filteredData := make(map[string]interface{})
	
	for _, field := range allowedFields {
		if value, exists := updateData[field]; exists {
			filteredData[field] = value
		}
	}
	
	return cli.Model(&model.UserModel{}).Where("id = ?", userID).Updates(filteredData).Error
}

// DeleteUser 删除用户
func (imp *UserInterfaceImp) DeleteUser(id string) error {
	cli := db.Get()
	return cli.Delete(&model.UserModel{}, "id = ?", id).Error
}

// GetUserByID 根据ID获取用户
func (imp *UserInterfaceImp) GetUserByID(id string) (*model.UserModel, error) {
	var user model.UserModel
	cli := db.Get()
	err := cli.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByUsername 根据用户名获取用户
func (imp *UserInterfaceImp) GetUserByUsername(username string) (*model.UserModel, error) {
	var user model.UserModel
	cli := db.Get()
	err := cli.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUsers 分页获取用户列表
func (imp *UserInterfaceImp) GetUsers(page, pageSize int) ([]*model.UserModel, int64, error) {
	var users []*model.UserModel
	var total int64
	
	cli := db.Get()
	
	// 获取总数
	if err := cli.Model(&model.UserModel{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	// 分页查询
	offset := (page - 1) * pageSize
	if err := cli.Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	
	return users, total, nil
}
