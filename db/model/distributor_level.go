package model

// DistributorLevelModel 分销等级模型
type DistributorLevelModel struct {
	ID              string   `gorm:"column:id;primaryKey;type:varchar(36)" json:"id"`
	Name            string   `gorm:"column:name;type:varchar(50);not null" json:"name"`
	Description     string   `gorm:"column:description;type:text" json:"description"`
	PickupDiscount  float64  `gorm:"column:pickup_discount;type:decimal(3,2);default:1.00;comment:'提货折扣'" json:"pickup_discount"`
	DirectRate      float64  `gorm:"column:direct_rate;type:decimal(5,4);default:0.0000;comment:'直接佣金比例'" json:"direct_rate"`
	IndirectRate    float64  `gorm:"column:indirect_rate;type:decimal(5,4);default:0.0000;comment:'间接佣金比例'" json:"indirect_rate"`
	UpgradeConditions []string `gorm:"column:upgrade_conditions;type:json;comment:'升级条件列表'" json:"upgrade_conditions"`
	AutoUpgrade     bool     `gorm:"column:auto_upgrade;type:bool;default:false" json:"auto_upgrade"`
	DistributorMode string   `gorm:"column:distributor_mode;type:varchar(20);default:'normal';comment:'分销模式'" json:"distributor_mode"`
	Sort            int      `gorm:"column:sort;type:int;default:0" json:"sort"`
}

func (DistributorLevelModel) TableName() string {
	return "distributor_levels"
}
