package models

import "github.com/jinzhu/gorm"

type PromotionPattern struct {
	gorm.Model
	PromotionTypeName     string   `json:"promotion_type_name"`
	PromotionPatternName  string   `json:"promotion_pattern_name"`
}

type PromotionPatternFrom struct {
	ID         				[]int `json:"id"`
	StartTime  				string  	`json:"start_time"`    				//开始时间
	StopTime  				string  	`json:"stop_time"`					//结束时间
	PromotionTypeName       string      `json:"promotion_type_name"`		//活动类型
	PromotionPatternName    string      `json:"promotion_pattern_name"`		//活动方式
	Price    				[]string    `json:"price"`						//原价
	PromotionDiscount  		int  		`json:"promotion_discount"`			//活动折扣，金额
	PromotionCount  		int  		`json:"promotion_count"`			//发行数量
	PromotionCash  			int  		`json:"promotion_cash"`				//活动条件
	ComplimentaryPatternId  int  		`json:"complimentary_pattern_id"`	//赠品方式
	ComplimentaryCash  		int  		`json:"complimentary_cash"`			//赠品条件
}


