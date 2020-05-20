package models

import "github.com/jinzhu/gorm"

// Shop 店铺信息
type Shop struct {
	gorm.Model
	ShopName    string `gorm:"type:varchar(50);unique_index" json:"shop_name"`
	ShopType    string `gorm:"type:varchar(10)" json:"shop_type"`
	Intro       string `gorm:"type:varchar(500)" json:"intro"`
	ShopNotice  string `gorm:"type:varchar(500)" json:"shop_notice"`
	ShopStatus  int    `gorm:"type:tinyint" json:"shop_status"`
	CompanyID   int64  `gorm:"type:bigint(20)" json:"company_id"`
	InvoiceType string `gorm:"type:varchar(50)"json:"invoice_type"`
}

// ShopVo ...推荐店铺
type ShopVo struct {
	ID        int64  `json:"id"`
	CompanyID int64  `json:"company_id"`
	ShopName  string `json:"shop_name"`
	Logo      string `json:"logo"`
}
