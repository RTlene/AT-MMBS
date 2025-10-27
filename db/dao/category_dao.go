package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

type CategoryInterfaceImp struct{}

// CreateCategory 创建分类
func (imp *CategoryInterfaceImp) CreateCategory(category *model.CategoryModel) error {
	cli := db.Get()
	return cli.Create(category).Error
}

// UpdateCategory 更新分类
func (imp *CategoryInterfaceImp) UpdateCategory(category *model.CategoryModel) error {
	cli := db.Get()
	return cli.Save(category).Error
}

// DeleteCategory 删除分类
func (imp *CategoryInterfaceImp) DeleteCategory(id string) error {
	cli := db.Get()
	return cli.Delete(&model.CategoryModel{}, "id = ?", id).Error
}

// GetCategoryByID 根据ID获取分类
func (imp *CategoryInterfaceImp) GetCategoryByID(id string) (*model.CategoryModel, error) {
	var category model.CategoryModel
	cli := db.Get()
	err := cli.Where("id = ?", id).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// GetCategories 获取所有分类
func (imp *CategoryInterfaceImp) GetCategories() ([]*model.CategoryModel, error) {
	var categories []*model.CategoryModel
	cli := db.Get()
	err := cli.Order("sort asc").Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// GetParentCategories 获取父分类
func (imp *CategoryInterfaceImp) GetParentCategories() ([]*model.CategoryModel, error) {
	var categories []*model.CategoryModel
	cli := db.Get()
	err := cli.Where("is_parent = ?", true).Order("sort asc").Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}
