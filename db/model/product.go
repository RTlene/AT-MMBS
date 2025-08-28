package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// StringArray 自定义类型，用于处理JSON数组
type StringArray []string

// Value 实现driver.Valuer接口
func (sa StringArray) Value() (driver.Value, error) {
	if sa == nil {
		return nil, nil
	}
	return json.Marshal(sa)
}

// Scan 实现sql.Scanner接口
func (sa *StringArray) Scan(value interface{}) error {
	if value == nil {
		*sa = nil
		return nil
	}
	
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, sa)
	case string:
		return json.Unmarshal([]byte(v), sa)
	default:
		return errors.New("cannot scan non-string value into StringArray")
	}
}

// ProductModel 商品模型
type ProductModel struct {
	ID           string      `gorm:"column:_id;primaryKey;type:varchar(36)" json:"id"`
	Name         string      `gorm:"column:name;type:varchar(100);not null" json:"name"`
	Category     string      `gorm:"column:category;type:varchar(50);not null" json:"category"`
	Price        float64     `gorm:"column:price;type:decimal(10,2);not null" json:"price"`
	Kucun        int         `gorm:"column:kucun;type:int;not null;default:0" json:"kucun"`
	Tp           StringArray `gorm:"column:tp;type:json" json:"tp"`
	Sp           string      `gorm:"column:sp;type:varchar(500)" json:"sp"`
	OverlayStyle string      `gorm:"column:overlay_style;type:text" json:"overlay_style"`
	Content      string      `gorm:"column:content;type:text" json:"content"`
	Status       int         `gorm:"column:status;type:int;default:1;comment:'1:上架 0:下架'" json:"status"`
	UpdateTime   time.Time   `gorm:"column:update_time;autoUpdateTime:milli" json:"update_time"`
	CreateTime   time.Time   `gorm:"column:create_time;autoCreateTime:milli" json:"create_time"`
}

func (ProductModel) TableName() string {
	return "products"
}
