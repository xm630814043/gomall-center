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
	PromotionDiscount  		int  		`json:"promotion_discount"`			//活动折扣，金额
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
	PromotionDiscount  		int  		`json:"promotion_discount"`			//活动折扣，金额
	PromotionCash  			int  		`json:"promotion_cash"`				//活动条件
	PromotionCount  		int  		`json:"promotion_count"`			//发行数量
	ShopId  				int  		`json:"shop_id"`					//店铺id
	ShopName  				string  	`json:"shop_name"`					//店铺名称
	ComplimentaryPatternId  int  		`json:"complimentary_pattern_id"`	//赠品方式
	ComplimentaryCash  		int  		`json:"complimentary_cash"`			//赠品条件
	PromotionType  			int  		`json:"promotion_type"`				//发布状态
	CategoryId  			[]int  		`json:"category_id"`				//类目id
	ProductId  				[]int  		`json:"product_id"`					//商品id
}

//PromotionActivityAbs ...促销活动详情
type PromotionActivityAbs struct {
	gorm.Model
	StartTime  				string  	`json:"start_time"`          		//开始时间
	StopTime  				string  	`json:"stop_time"`            		//结束时间
	PromotionPatternId  	int  		`json:"promotion_pattern_id"`  		//促销方式
	PromotionTheme  		string  	`json:"promotion_theme"`        	//活动主题
	PromotionDescribe  		string  	`json:"promotion_describe"`  		//活动描述
	PromotionDiscount      	int  		`json:"promotion_discount"`     	//活动折扣，金额
	PromotionCash          	int  		`json:"promotion_cash"`         	//活动条件
	PromotionCount         	int  		`json:"promotion_count"`        	//发行数量
	ShopId                 	int  		`json:"shop_id"`                	//店铺id
	ShopName               	string  	`json:"shop_name"`              	//店铺名称
	ComplimentaryPatternId 	int  		`json:"complimentary_pattern_id"`   //赠品方式
	ComplimentaryCash      	int  		`json:"complimentary_cash"`         //赠品条件
	PromotionType          	int  		`json:"promotion_type"`             //发布状态
	Products              	[]*PromotionProducts                            //促销商品
	Complimentary          	[]*PromotionProducts                            //促销赠品
}
//PromotionProducts ...促销活动商品详情
type PromotionProducts struct {
	Id 				   int 			 `json:"id"`
	ProductName        string           `json:"product_name"`        		// 商品名称
	Price              string           `json:"price"`           			// 原价（零售价）
	Pic                string           `json:"pic;type"`
	AlbumPics          string           `json:"album_pics"`         		// 图片，逗号分割，限定5张
	BrandName          string           `json:"brand_name"`          		// 品牌名称
	CategoryName       string           `json:"category_name"`       		// 分类名称
	TotalStocks        int              `json:"total_stocks"`            	// 总库存
	ClickNum           int              `json:"click_num"`               	// 点击查看次数
	SellNum            int              `json:"sell_num"`                	// 销售数量
}
//PromotionProductList ...促销活动商品列表
type PromotionProductList struct {
	Products              	[]*PromotionProducts                            //促销商品
	Complimentary          	[]*PromotionProducts                            //促销赠品
	PromotionActivity		[]*PromotionActivity
}
//DiscountCoupon struct ...优惠券表
type DiscountCoupon struct {
	PromotionActivityId int 		`json:"promotion_activity_id"`
	StartTime  			string  	`json:"start_time"`    			//开始时间
	StopTime  			string  	`json:"stop_time"`				//结束时间
	PromotionPatternId  int  		`json:"promotion_pattern_id"`	//促销方式
	PromotionDiscount  	int  		`json:"promotion_discount"`		//活动折扣，金额
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
	CategoryId 			 string `json:"category_id"`
	ProductId  			 string `json:"product_id"`
	ShopId  			 int 	`json:"shop_id"`
}
//ComplimentaryRelation struct ...赠品关系表
type ComplimentaryRelation struct {
	PromotionActivityId  int 	`json:"promotion_activity_id"`
	CategoryId 			 string `json:"category_id"`
	ProductId  			 string `json:"product_id"`
	ShopId  			 int 	`json:"shop_id"`
}

