package model

import "time"

// MemberModel 会员模型
type MemberModel struct {
	ID              string    `gorm:"column:_id;primaryKey;type:varchar(36)" json:"_id"`
	NickName        string    `gorm:"column:nick_name;type:varchar(50)" json:"nick_name"`
	OpenID          string    `gorm:"column:openid;type:varchar(100);uniqueIndex" json:"openid"`
	UpdateTime      time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time"`
	CreateTime      time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`
	Birthday        time.Time `gorm:"column:birthday;type:date" json:"birthday"`
	LastLoginTime   time.Time `gorm:"column:last_login_time" json:"last_login_time"`
	Gender          string    `gorm:"column:gender;type:varchar(10)" json:"gender"`
	Phone           string    `gorm:"column:phone;type:varchar(20)" json:"phone"`
	PhoneVerified   bool      `gorm:"column:phone_verified;type:bool;default:false" json:"phone_verified"`
	PhoneVerifyTime time.Time `gorm:"column:phone_verify_time" json:"phone_verify_time"`
	Region          []string  `gorm:"column:region;type:json" json:"region"`
	Addresses       []string  `gorm:"column:addresses;type:json" json:"addresses"`
	ReferrerID      string    `gorm:"column:referrer_id;type:varchar(36)" json:"referrer_id"`
	DirectFans      []string  `gorm:"column:direct_fans;type:json" json:"direct_fans"`
	Orders          []string  `gorm:"column:orders;type:json" json:"orders"`
	LevelID         string    `gorm:"column:level_id;type:varchar(36)" json:"level_id"`
	DistributorID   string    `gorm:"column:distributor_id;type:varchar(36)" json:"distributor_id"`
	Points          int       `gorm:"column:points;type:int;default:0" json:"points"`
	Commission      float64   `gorm:"column:commission;type:decimal(10,2);default:0" json:"commission"`
	CommissionLogs  []string  `gorm:"column:commission_logs;type:json" json:"commission_logs"`
	PointsLogs      []string  `gorm:"column:points_logs;type:json" json:"points_logs"`
	AutoUpgrade     bool      `gorm:"column:auto_upgrade;type:bool;default:true" json:"auto_upgrade"`
}

func (MemberModel) TableName() string {
	return "members"
}
