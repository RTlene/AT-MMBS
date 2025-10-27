package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

type DistributorLevelInterfaceImp struct{}

// CreateDistributorLevel 创建分销等级
func (imp *DistributorLevelInterfaceImp) CreateDistributorLevel(level *model.DistributorLevelModel) error {
	cli := db.Get()
	return cli.Create(level).Error
}

// UpdateDistributorLevel 更新分销等级
func (imp *DistributorLevelInterfaceImp) UpdateDistributorLevel(level *model.DistributorLevelModel) error {
	cli := db.Get()
	return cli.Save(level).Error
}

// DeleteDistributorLevel 删除分销等级
func (imp *DistributorLevelInterfaceImp) DeleteDistributorLevel(id string) error {
	cli := db.Get()
	return cli.Delete(&model.DistributorLevelModel{}, "id = ?", id).Error
}

// GetDistributorLevelByID 根据ID获取分销等级
func (imp *DistributorLevelInterfaceImp) GetDistributorLevelByID(id string) (*model.DistributorLevelModel, error) {
	var level model.DistributorLevelModel
	cli := db.Get()
	err := cli.Where("id = ?", id).First(&level).Error
	if err != nil {
		return nil, err
	}
	return &level, nil
}

// GetDistributorLevels 获取所有分销等级
func (imp *DistributorLevelInterfaceImp) GetDistributorLevels() ([]*model.DistributorLevelModel, error) {
	var levels []*model.DistributorLevelModel
	cli := db.Get()
	err := cli.Order("sort asc").Find(&levels).Error
	if err != nil {
		return nil, err
	}
	return levels, nil
}
