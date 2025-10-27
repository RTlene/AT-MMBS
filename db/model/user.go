package model

import "time"

// UserModel 用户模型
type UserModel struct {
	ID       string    `gorm:"column:id;primaryKey;type:varchar(36)" json:"id"`
	Username string    `gorm:"column:username;type:varchar(50);not null;uniqueIndex" json:"username"`
	Password string    `gorm:"column:password;type:varchar(255);not null" json:"password"`
	Role     string    `gorm:"column:role;type:varchar(20);default:'user'" json:"role"`
	CreateAt time.Time `gorm:"column:create_at;autoCreateTime" json:"create_at"`
	UpdateAt time.Time `gorm:"column:update_at;autoUpdateTime" json:"update_at"`
}

func (UserModel) TableName() string {
	return "users"
}
