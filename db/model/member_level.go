package model

// MemberLevelModel 会员等级模型
type MemberLevelModel struct {
	ID          string   `gorm:"column:id;primaryKey;type:varchar(36)" json:"id"`
	Name        string   `gorm:"column:name;type:varchar(50);not null" json:"name"`
	Description string   `gorm:"column:description;type:text" json:"description"`
	Discount    float64  `gorm:"column:discount;type:decimal(3,2);default:1.00;comment:'购物折扣'" json:"discount"`
	PointsRate  float64  `gorm:"column:points_rate;type:decimal(3,2);default:1.00;comment:'积分倍率'" json:"points_rate"`
	UpgradeConditions []string `gorm:"column:upgrade_conditions;type:json;comment:'升级条件列表'" json:"upgrade_conditions"`
	AutoUpgrade bool     `gorm:"column:auto_upgrade;type:bool;default:false" json:"auto_upgrade"`
	Sort        int      `gorm:"column:sort;type:int;default:0" json:"sort"`
}

func (MemberLevelModel) TableName() string {
	return "member_levels"
}
