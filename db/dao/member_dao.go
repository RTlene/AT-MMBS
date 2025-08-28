package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

type MemberInterfaceImp struct{}

// CreateMember 创建会员
func (imp *MemberInterfaceImp) CreateMember(member *model.MemberModel) error {
	cli := db.Get()
	return cli.Create(member).Error
}

// UpdateMember 更新会员
func (imp *MemberInterfaceImp) UpdateMember(member *model.MemberModel) error {
	cli := db.Get()
	return cli.Save(member).Error
}

// DeleteMember 删除会员
func (imp *MemberInterfaceImp) DeleteMember(id string) error {
	cli := db.Get()
	return cli.Delete(&model.MemberModel{}, "_id = ?", id).Error
}

// GetMemberByID 根据ID获取会员
func (imp *MemberInterfaceImp) GetMemberByID(id string) (*model.MemberModel, error) {
	var member model.MemberModel
	cli := db.Get()
	err := cli.Where("_id = ?", id).First(&member).Error
	if err != nil {
		return nil, err
	}
	return &member, nil
}

// GetMemberByOpenID 根据OpenID获取会员
func (imp *MemberInterfaceImp) GetMemberByOpenID(openid string) (*model.MemberModel, error) {
	var member model.MemberModel
	cli := db.Get()
	err := cli.Where("openid = ?", openid).First(&member).Error
	if err != nil {
		return nil, err
	}
	return &member, nil
}

// GetMembers 分页获取会员列表
func (imp *MemberInterfaceImp) GetMembers(page, pageSize int) ([]*model.MemberModel, int64, error) {
	var members []*model.MemberModel
	var total int64
	
	cli := db.Get()
	
	// 获取总数
	if err := cli.Model(&model.MemberModel{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	// 分页查询
	offset := (page - 1) * pageSize
	if err := cli.Offset(offset).Limit(pageSize).Order("create_time desc").Find(&members).Error; err != nil {
		return nil, 0, err
	}
	
	return members, total, nil
}

// GetMemberNetwork 获取会员关系网
func (imp *MemberInterfaceImp) GetMemberNetwork(id string) (map[string]interface{}, error) {
	cli := db.Get()
	
	// 获取当前会员
	var member model.MemberModel
	if err := cli.Where("_id = ?", id).First(&member).Error; err != nil {
		return nil, err
	}
	
	// 获取推荐人
	var referrer *model.MemberModel
	if member.ReferrerID != "" {
		referrer = &model.MemberModel{}
		if err := cli.Where("_id = ?", member.ReferrerID).First(referrer).Error; err != nil {
			referrer = nil
		}
	}
	
	// 获取直接粉丝（一代）
	var directFans []*model.MemberModel
	if err := cli.Where("referrer_id = ?", id).Find(&directFans).Error; err != nil {
		directFans = []*model.MemberModel{}
	}
	
	// 获取间接粉丝（二代）
	var indirectFans []*model.MemberModel
	if len(directFans) > 0 {
		var directFanIDs []string
		for _, fan := range directFans {
			directFanIDs = append(directFanIDs, fan.ID)
		}
		if err := cli.Where("referrer_id IN ?", directFanIDs).Find(&indirectFans).Error; err != nil {
			indirectFans = []*model.MemberModel{}
		}
	}
	
	return map[string]interface{}{
		"referrer":      referrer,
		"direct_fans":   directFans,
		"indirect_fans": indirectFans,
	}, nil
}

// GetMemberDistribution 获取会员分销信息
func (imp *MemberInterfaceImp) GetMemberDistribution(id string) (map[string]interface{}, error) {
	cli := db.Get()
	
	// 获取当前会员
	var member model.MemberModel
	if err := cli.Where("_id = ?", id).First(&member).Error; err != nil {
		return nil, err
	}
	
	// 获取分销等级信息
	var distributorLevel *model.DistributorLevelModel
	if member.DistributorID != "" {
		distributorLevel = &model.DistributorLevelModel{}
		if err := cli.Where("id = ?", member.DistributorID).First(distributorLevel).Error; err != nil {
			distributorLevel = nil
		}
	}
	
	// 这里可以添加更多分销相关的统计逻辑
	// 比如销售额、佣金记录等
	
	return map[string]interface{}{
		"member":            member,
		"distributor_level": distributorLevel,
		"commission":        member.Commission,
		// 可以添加更多分销信息
	}, nil
}

// GetMemberConsumption 获取会员消费记录
func (imp *MemberInterfaceImp) GetMemberConsumption(id string) ([]*model.OrderModel, error) {
	var orders []*model.OrderModel
	cli := db.Get()
	
	err := cli.Where("member_id = ? AND payment_status = ?", id, 1).Order("create_time desc").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	
	return orders, nil
}
