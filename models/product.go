package models

import "github.com/jinzhu/gorm"

// Product 产品表
type Product struct {
	gorm.Model
	ShopID             int64            `gorm:"column:shop_id;type:bigint(20)"`
	PublishStatus      int8             `gorm:"column:publish_status;type:tinyint(1)"`   // 发布状态：0->上架；1->下架
	IsRecommand        int8             `gorm:"column:is_recommand;type:tinyint(1)"`     // 是否推荐：0->不推荐；1->推荐
	Price              string           `gorm:"column:price;type:varchar(45)"`           // 原价（零售价）
	WholesalePrice     string           `gorm:"column:wholesale_price;type:varchar(45)"` // 批发价
	BrandID            string           `gorm:"column:brand_id;type:varchar(45)"`
	CategoryID         int64            `gorm:"column:category_id;type:bigint(20)"`
	Description        string           `gorm:"column:description;type:text"`
	Unit               string           `gorm:"column:unit;type:varchar(20)"`
	Tags               string           `gorm:"column:tags;type:varchar(100)"`
	Pic                string           `gorm:"column:pic;type:varchar(100)"`
	AlbumPics          string           `gorm:"column:album_pics;type:varchar(500)"`         // 图片，逗号分割，限定5张
	BrandName          string           `gorm:"column:brand_name;type:varchar(45)"`          // 品牌名称
	CategoryName       string           `gorm:"column:category_name;type:varchar(45)"`       // 分类名称
	TotalStocks        int              `gorm:"column:total_stocks;type:int(11)"`            // 总库存
	DeliveryMode       string           `gorm:"column:delivery_mode;type:varchar(1000)"`     // 配送方式
	DeliveryTemplateID int64            `gorm:"column:delivery_template_id;type:bigint(20)"` // 运费模板id
	StandardID         int64            `gorm:"column:standard_id;type:bigint(20)"`          // 标准产品库
	ClickNum           int              `gorm:"column:click_num;type:int(11)"`               // 点击查看次数
	SellNum            int              `gorm:"column:sell_num;type:int(11)"`                // 销售数量
	ProductName        string           `gorm:"column:product_name;type:varchar(50)"`        // 商品名称
	MinBuyNum          int              `gorm:"column:min_buy_num;type:int(11)"`             // 起订量（最小购买量）
	SaveData           string           `gorm:"column:save_data;type:varchar(10000)"`        // 前台需要的属性相关信息
	Producer           string           `gorm:"column:producer;type:varchar(200)"`           // 生产厂家名称
	ProductSkus        []*ProductSku    `json:"product_skus"`
	ProductAttrs       []*ProductAttrVo `json:"product_attrs"`
}

// ProductSku ...Sku 产品sku表
type ProductSku struct {
	ID             int64   `gorm:"primary_key;column:id;type:bigint(20);not null"`
	ProductID      int64   `gorm:"column:product_id;type:bigint(20);not null"`
	RetailPrice    float64 `gorm:"column:retail_price;type:decimal(15,2);not null"`  // 零售价
	WholesalePrice float64 `gorm:"column:wholesale_price;type:varchar(45);not null"` // 批发价
	Properties     string  `gorm:"column:properties;type:varchar(2000)"`             // 销售规格属性集合格式为-> p1:v1;p2:v2
	Stocks         int     `gorm:"column:stocks;type:int(11);not null"`              // 库存量
	PartyCode      string  `gorm:"column:party_code;type:varchar(50)"`               // 商家编码
	ModelID        string  `gorm:"column:model_id;type:varchar(50)"`                 // 商品条形码
	Pic            string  `gorm:"column:pic;type:varchar(500)"`                     // sku图片
	SkuName        string  `gorm:"column:sku_name;type:varchar(255)"`                // sku名称
	SkuStatus      int     `gorm:"column:sku_status;type:tinyint(2)"`                // 0禁用 1 启用
	PackagingType  string  `gorm:"column:packaging_type;type:varchar(20)"`           // sku 包装方式
	SkuWeight      float64 `gorm:"column:sku_weight;type:decimal(15,2)"`             // sku 重量
	SaveData       string  `gorm:"column:save_data;type:varchar(10000)"`             // 前台需要的属性相关信息
}

//ProductAttrVo ... 后台product返回时候 返回的attr
type ProductAttrVo struct {
	ID        int64  `json:"id"`
	PropID    int64  `json:"prop_id"`
	PropName  string `json:"prop_name"`
	PropValue string `json:"prop_value"`
	PropType  int8   `json:"prop_type"`
	ProductID int64  `json:"product_id"`
	InputType int8   `json:"input_type"`
	IsDefault int8   `json:"is_default"`
}

//ProductBasicInfo ...产品基本摘要信息  用于列表显示
type ProductBasicInfo struct {
	ID          int    `gorm:"column:id" json:"id"`
	ProductName string `gorm:"column:product_name" json:"product_name"`
	Producer    string `gorm:"column:producer" json:"producer"`
}

//ProductAudit ...商品详情,要审核的信息
type ProductAudit struct {
	ID            int    `gorm:"column:id" json:"id"`
	PublishStatus int    `gorm:"column:publish_status" json:"publish_status"`
	Description   string `gorm:"column:description" json:"description"`
	Pic           string `gorm:"column:pic" json:"pic"`
}

//ProductProp ...商品prop类别
type ProductProp struct {
	ID        int    `json:"id" `
	PropName  string `json:"prop_name" binding:"required"`
	PropType  int    `json:"prop_type"`
	ShopID    int    `json:"shop_id"`
	InputType int    `json:"input_type"`
	IsDefault int    `json:"is_default"`
}

//PropForm ...添加属性
type PropForm struct {
	PropName  string `json:"prop_name" binding:"required" validate:"required" `
	PropType  string `json:"prop_type" validate:"required"`
	InputType string `json:"input_type" validate:"required"`
}

//ProductPropValue ...商品prop_value
type ProductPropValue struct {
	ID        int    ` json:"id"`
	PropID    int    ` json:"prop_id" `
	PropValue string ` json:"prop_value" validate:"required"`
}

//ProductPutStatic ...接收商品发布状态
type ProductPutStatic struct {
	PublishStatus int `gorm:"column:publish_status" json:"publish_status" validate:"required"`
}
