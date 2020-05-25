package models

import "github.com/jinzhu/gorm"

//PromotionActivity struct  促销活动表
type PromotionActivity struct {
	gorm.Model
	StartTime  				string  	`json:"start_time"`    			//开始时间
	StopTime  				string  	`json:"stop_time"`				//结束时间
	PromotionPatternId  	int  		`json:"promotion_pattern_id"`	//促销规则
	PromotionTheme  		string  	`json:"promotion_theme"`		//活动主题
	PromotionDescribe  		string  	`json:"promotion_describe"`		//活动描述
	PromotionDiscount  		int  		`json:"promotion_discount"`		//活动折扣，金额
	PromotionCash  			int  		`json:"promotion_cash"`			//活动条件
	PromotionCount  		int  		`json:"promotion_count"`		//发行数量
	ShopId  				int  		`json:"shop_id"`				//店铺id
	ShopName  				string  	`json:"shop_name"`				//店铺名称
	ComplimentaryPatternId  int  		`json:"complimentary_pattern_id"`		//赠品方式
	ComplimentaryCash  		int  		`json:"complimentary_cash"`		//赠品条件
	PromotionType  			int  		`json:"promotion_type"`			//发布状态
}

//PromotionActivityForm  接受表单
type PromotionActivityForm struct {
	gorm.Model
	StartTime  				string  	`json:"start_time"`    			//开始时间
	StopTime  				string  	`json:"stop_time"`				//结束时间
	PromotionPatternId  	int  		`json:"promotion_pattern_id"`	//促销方式
	PromotionTheme  		string  	`json:"promotion_theme"`		//活动主题
	PromotionDescribe  		string  	`json:"promotion_describe"`		//活动描述
	PromotionDiscount  		int  		`json:"promotion_discount"`		//活动折扣，金额
	PromotionCash  			int  		`json:"promotion_cash"`			//活动条件
	PromotionCount  		int  		`json:"promotion_count"`		//发行数量
	ShopId  				int  		`json:"shop_id"`				//店铺id
	ShopName  				string  	`json:"shop_name"`				//店铺名称
	ComplimentaryPatternId  int  		`json:"complimentary_pattern_id"`		//赠品方式
	ComplimentaryCash  		int  		`json:"complimentary_cash"`		//赠品条件
	PromotionType  			int  		`json:"promotion_type"`			//发布状态
	CategoryId  			[]int  		`json:"category_id"`			//类目id
	ProductId  				[]int  		`json:"product_id"`				//商品id
}

//PromotionActivityAbs  活动详情
type PromotionActivityAbs struct {
	gorm.Model
	StartTime  				string  	`json:"start_time"`          //开始时间
	StopTime  				string  	`json:"stop_time"`            //结束时间
	PromotionPatternId  	int  		`json:"promotion_pattern_id"`  //促销方式
	PromotionTheme  		string  	`json:"promotion_theme"`        //活动主题
	PromotionDescribe  		string  	`json:"promotion_describe"`  //活动描述
	PromotionDiscount      	int  		`json:"promotion_discount"`       //活动折扣，金额
	PromotionCash          	int  		`json:"promotion_cash"`           //活动条件
	PromotionCount         	int  		`json:"promotion_count"`          //发行数量
	ShopId                 	int  		`json:"shop_id"`                  //店铺id
	ShopName               	string  	`json:"shop_name"`                 //店铺名称
	ComplimentaryPatternId 	int  		`json:"complimentary_pattern_id"` //赠品方式
	ComplimentaryCash      	int  		`json:"complimentary_cash"`       //赠品条件
	PromotionType          	int  		`json:"promotion_type"`           //发布状态
	Products              	[]*Products                                 //商品
	Complimentary          	[]*Products                                //赠品
}

// Products ...商品/赠品详情
type Products struct {
	Id 					int 			 `json:"id"`
	ProductName         string           `json:"product_name"`        // 商品名称
}
//DiscountCoupon struct  优惠券表
type DiscountCoupon struct {
	PromotionActivityId int 		`json:"promotion_activity_id"`
	StartTime  			string  	`json:"start_time"`    			//开始时间
	StopTime  			string  	`json:"stop_time"`				//结束时间
	PromotionPatternId  int  		`json:"promotion_pattern_id"`	//促销方式
	PromotionDiscount  	int  		`json:"promotion_discount"`		//活动折扣，金额
	PromotionCash  		int  		`json:"promotion_cash"`			//活动条件
	AccountId  			int  		`json:"account_id"`				//用户id
	CompanyId  			int  		`json:"company_id"`				//企业id
	CompanyName  		string  	`json:"company_name"`			//企业名称
	PromotionType  		int  		`json:"promotion_type"`			//发布状态
}

//活动商品关系表
type ProductRelation struct {
	PromotionActivityId  int `json:"promotion_activity_id"`
	CategoryId 			 string `json:"category_id"`
	ProductId  			 string `json:"product_id"`
	ShopId  			 int `json:"shop_id"`
}

//赠品关系表
type ComplimentaryRelation struct {
	PromotionActivityId  int `json:"promotion_activity_id"`
	CategoryId 			 string `json:"category_id"`
	ProductId  			 string `json:"product_id"`
	ShopId  			 int `json:"shop_id"`
}

