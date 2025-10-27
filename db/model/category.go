package model

// CategoryModel 商品分类模型
type CategoryModel struct {
	ID           string      `gorm:"column:id;primaryKey;type:varchar(36)" json:"id"`
	Name         string      `gorm:"column:name;type:varchar(50);not null" json:"name"`
	Description  string      `gorm:"column:description;type:text" json:"description"`
	SubCategory  string `gorm:"column:sub_category;type:text" json:"sub_category"`
	Status       int         `gorm:"column:status;type:int;default:1;comment:'1:启用 0:禁用'" json:"status"`
	IsParent     bool        `gorm:"column:is_parent;type:bool;default:false" json:"is_parent"`
	ParentID     string      `gorm:"column:parent_id;type:varchar(36)" json:"parent_id"`
	Sort         int         `gorm:"column:sort;type:int;default:0" json:"sort"`
}

func (CategoryModel) TableName() string {
	return "categories"
}
