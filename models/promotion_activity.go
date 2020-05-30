package models

import "github.com/jinzhu/gorm"

//PromotionActivity struct ...促销活动表
type PromotionActivity struct {
	gorm.Model
	StartTime  				string  	`json:"start_time"`    				//开始时间
	StopTime  				string  	`json:"stop_time"`					//结束时间
	PromotionPatternId  	int  		`json:"promotion_pattern_id"`		//促销规则
	PromotionTheme  		string  	`json:"promotion_theme"`			//活动主题
	PromotionDescribe  		string  	`json:"promotion_describe"`			//活动描述
	PromotionDiscount  		string  	`json:"promotion_discount"`			//活动折扣，金额
	PromotionCash  			int  		`json:"promotion_cash"`				//活动条件
	PromotionCount  		int  		`json:"promotion_count"`			//发行数量
	ShopId  				int  		`json:"shop_id"`					//店铺id
	ShopName  				string  	`json:"shop_name"`					//店铺名称
	ComplimentaryPatternId  int  		`json:"complimentary_pattern_id"`	//赠品方式
	ComplimentaryCash  		int  		`json:"complimentary_cash"`			//赠品条件
	PromotionType  			int  		`json:"promotion_type"`				//发布状态
	PromotionActivityId     int 		`json:"promotion_activity_id"`
}

//PromotionActivityForm ...促销活动表单
type PromotionActivityForm struct {
	gorm.Model
	StartTime  				string  	`json:"start_time"`    				//开始时间
	StopTime  				string  	`json:"stop_time"`					//结束时间
	PromotionPatternId  	int  		`json:"promotion_pattern_id"`		//促销方式
	PromotionTheme  		string  	`json:"promotion_theme"`			//活动主题
	PromotionDescribe  		string  	`json:"promotion_describe"`			//活动描述
	PromotionDiscount  		string  	`json:"promotion_discount"`			//活动折扣，金额
	PromotionCash  			int  		`json:"promotion_cash"`				//活动条件
	PromotionCount  		int  		`json:"promotion_count"`			//发行数量
	ShopId  				int  		`json:"shop_id"`					//店铺id
	ShopName  				string  	`json:"shop_name"`					//店铺名称
	ComplimentaryPatternId  int  		`json:"complimentary_pattern_id"`	//赠品方式
	ComplimentaryCash  		int  		`json:"complimentary_cash"`			//赠品条件
	PromotionType  			int  		`json:"promotion_type"`				//发布状态
	CategoryId  			[]int  		`json:"category_id"`				//类目id
	ProductId  				[]int  		`json:"product_id"`					//商品id
	ProductSkuId  		 	[]int 		`json:"product_sku_id"`
}

//PromotionActivityAbs ...促销活动详情
type PromotionActivityAbs struct {
	gorm.Model
	StartTime  				string  	`json:"start_time"`          		//开始时间
	StopTime  				string  	`json:"stop_time"`            		//结束时间
	PromotionPatternId  	int  		`json:"promotion_pattern_id"`  		//促销方式
	PromotionTheme  		string  	`json:"promotion_theme"`        	//活动主题
	PromotionDescribe  		string  	`json:"promotion_describe"`  		//活动描述
	PromotionDiscount      	string  	`json:"promotion_discount"`     	//活动折扣，金额
	PromotionCash          	int  		`json:"promotion_cash"`         	//活动条件
	PromotionCount         	int  		`json:"promotion_count"`        	//发行数量
	ShopId                 	int  		`json:"shop_id"`                	//店铺id
	ShopName               	string  	`json:"shop_name"`              	//店铺名称
	ComplimentaryPatternId 	int  		`json:"complimentary_pattern_id"`   //赠品方式
	ComplimentaryCash      	int  		`json:"complimentary_cash"`         //赠品条件
	PromotionType          	int  		`json:"promotion_type"`             //发布状态
	Products              	[]*PromotionProductSku                            //促销商品
	Complimentary          	[]*PromotionProductSku                            //促销赠品
}
//PromotionProducts ...促销活动商品详情
type PromotionProductSku struct {
	Id 				   int 			 	`json:"id"`
	SkuName        	   string  			`json:"sku_name"`               // sku名称
	RetailPrice    	   float64 			`json:"retail_price"`  			// 零售价
	WholesalePrice     float64 			`json:"wholesale_price"` 		// 批发价
	CurrentPrice       float64          `json:"current_price"`			// 现价（活动价）
	Pic                string           `json:"pic"`
	Properties     	   string  			`json:"properties"`             // 销售规格属性集合格式为-> p1:v1;p2:v2
	Stocks         	   int     			`json:"stocks"`            		// 库存量
	PackagingType  	   string  			`json:"packaging_type"`         // sku 包装方式
	SkuWeight          float64 			`json:"sku_weight"`             // sku 重量
}
//PromotionProductList ...促销活动商品列表
type PromotionProductList struct {
	Products              	[]*PromotionProductSku                            //促销商品
	Complimentary          	[]*PromotionProductSku                            //促销赠品
	PromotionActivity		[]*PromotionActivity
}
//DiscountCoupon struct ...优惠券表
type DiscountCoupon struct {
	PromotionActivityId int 		`json:"promotion_activity_id"`
	StartTime  			string  	`json:"start_time"`    			//开始时间
	StopTime  			string  	`json:"stop_time"`				//结束时间
	PromotionPatternId  int  		`json:"promotion_pattern_id"`	//促销方式
	PromotionDiscount  	string  	`json:"promotion_discount"`		//活动折扣，金额
	PromotionCash  		int  		`json:"promotion_cash"`			//活动条件
	ShopId              int      	`json:"shop_id"`                  //店铺id
	PromotionType  		int  		`json:"promotion_type"`			//发布状态
	AccountId  			int  		`json:"account_id"`				//用户id
	CompanyId  			int  		`json:"company_id"`				//企业id
	CompanyName  		string  	`json:"company_name"`			//企业名称
}
//ProductRelation struct ...活动商品关系表
type ProductRelation struct {
	PromotionActivityId  int 	`json:"promotion_activity_id"`
	ShopId  			 int 	`json:"shop_id"`
	CategoryId 			 string `json:"category_id"`
	ProductId  			 string `json:"product_id"`
	ProductSkuId  		 string `json:"product_sku_id"`
}
//ComplimentaryRelation struct ...赠品关系表
type ComplimentaryRelation struct {
	PromotionActivityId  int 	`json:"promotion_activity_id"`
	ShopId  			 int 	`json:"shop_id"`
	CategoryId 			 string `json:"category_id"`
	ProductId  			 string `json:"product_id"`
	ProductSkuId  		 string `json:"product_sku_id"`
}

