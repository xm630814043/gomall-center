package models

import (
	"gomall-center/pkg/enum"

	"github.com/jinzhu/gorm"
)

// Company 企业信息表
type Company struct {
	gorm.Model
	CompanyName     string             `json:"company_name"`
	BusinessLicense string             `json:"business_license"`  // 营业执照号
	LegalPersonName string             `json:"legal_person_name"` // 法人姓名
	LegalPersonID   string             `json:"legal_person_id"`   // 法人身份证
	Taxpayer        string             `json:"taxpayer"`          // 纳税识别号
	Contacts        string             `json:"contacts"`          // 联系人
	Tel             string             `json:"tel"`               // 联系电话
	Lat             string             `json:"lat"`               // 经度
	Lng             string             `json:"lng"`               // 维度
	Address         string             `json:"address"`           // 地址
	Province        string             `json:"province"`          // 省份
	City            string             `json:"city"`              // 城市
	Area            string             `json:"area"`              // 地区
	Logo            string             `json:"logo"`              // 企业logo
	CompanyType     enum.CompanyType   `json:"company_type"`      // 企业类型
	CompanyStatus   enum.CompanyStatus `json:"company_status"`
	Shop            *Shop              `json:"shop"`
	Accounts        []*Account
}

// Shop 店铺信息
type Shop struct {
	gorm.Model
	ShopName    string `json:"shop_name"`
	ShopType    string `json:"shop_type"`
	Intro       string `json:"intro"`
	ShopNotice  string `json:"shop_notice"`
	ShopStatus  int    `json:"shop_status"`
	CompanyID   int64  `json:"company_id"`
	InvoiceType string `json:"invoice_type"`
}
