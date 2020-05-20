package models

import (
	"gomall-center/pkg/enum"

	"github.com/jinzhu/gorm"
)

//Category ...商品种类
type Category struct {
	gorm.Model
	ParentID       int                 `gorm:"column:parent_id;type:int(10)" json:"parent_id"`
	CategoryName   string              `gorm:"column:category_name;type:varchar(45)" json:"category_name"`
	Icon           string              `gorm:"column:icon;type:varchar(100)" json:"icon"`
	CategoryStatus enum.CategoryStatus `gorm:"column:category_status;type:tinyint(1)" json:"category_status"` // 状态：0->下线；1- >正常
	Level          int                 `gorm:"column:level;type:tinyint(2)" json:"level"`                     // 层级
	Path           string              `gorm:"column:path;type:varchar(200)" json:"path"`
}

//CategoryNode ...分类树
type CategoryNode struct {
	ID           uint            `json:"id"`
	ParentID     int             `json:"parent_id"`
	CategoryName string          `json:"category_name"`
	Icon         string          `json:"icon"`
	Level        int             `json:"level"`    // 层级
	Children     []*CategoryNode `json:"children"` // 子节点
}

//CategoryProp ...分类属性关联表categoryProp
type CategoryProp struct {
	ID         int `gorm:"primary_key;column:id;type:bigint(20);not null"`
	CategoryID int `gorm:"column:category_id;type:bigint(20)"`
	PropID     int `gorm:"column:prop_id;type:bigint(20)"`
}
