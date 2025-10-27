package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

type MemberLevelInterfaceImp struct{}

// CreateMemberLevel 创建会员等级
func (imp *MemberLevelInterfaceImp) CreateMemberLevel(level *model.MemberLevelModel) error {
	cli := db.Get()
	return cli.Create(level).Error
}

// UpdateMemberLevel 更新会员等级
func (imp *MemberLevelInterfaceImp) UpdateMemberLevel(level *model.MemberLevelModel) error {
	cli := db.Get()
	return cli.Save(level).Error
}

// DeleteMemberLevel 删除会员等级
func (imp *MemberLevelInterfaceImp) DeleteMemberLevel(id string) error {
	cli := db.Get()
	return cli.Delete(&model.MemberLevelModel{}, "id = ?", id).Error
}

// GetMemberLevelByID 根据ID获取会员等级
func (imp *MemberLevelInterfaceImp) GetMemberLevelByID(id string) (*model.MemberLevelModel, error) {
	var level model.MemberLevelModel
	cli := db.Get()
	err := cli.Where("id = ?", id).First(&level).Error
	if err != nil {
		return nil, err
	}
	return &level, nil
}

// GetMemberLevels 获取所有会员等级
func (imp *MemberLevelInterfaceImp) GetMemberLevels() ([]*model.MemberLevelModel, error) {
	var levels []*model.MemberLevelModel
	cli := db.Get()
	err := cli.Order("sort asc").Find(&levels).Error
	if err != nil {
		return nil, err
	}
	return levels, nil
}
