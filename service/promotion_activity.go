package service

import (
	"fmt"
	"gomall-center/models"
	"gomall-center/pkg/e"
	"gomall-center/pkg/web"
	"strings"
	"time"
)

type PromotionActivity struct {
	Service
}

func NewPromotionActivity(cotent *web.RequestContext) *PromotionActivity  {
	c := &PromotionActivity{InitService(cotent)}
	return c
}

func (c *PromotionActivity) InsertPromotionActivity (form *models.PromotionActivityForm) int {
	categoryIda := strings.Replace(strings.Trim(fmt.Sprint(form.CategoryId), ""), "[", "", -1)
	categoryId := strings.Replace(categoryIda, "]", "", -1 )
	productIda := strings.Replace(strings.Trim(fmt.Sprint(form.ProductId), ""), "[", " ", -1)
	productId := strings.Replace(productIda, "]", "", -1 )
	var  promotionactivity = models.PromotionActivity{StartTime:form.StartTime,StopTime:form.StopTime,PromotionPatternId:form.PromotionPatternId,PromotionTheme:form.PromotionTheme,
		PromotionDescribe:form.PromotionDescribe,PromotionDiscount:form.PromotionDiscount,PromotionCash:form.PromotionCash,PromotionCount:form.PromotionCount,ShopId:form.ShopId,
		ShopName:form.ShopName,ComplimentaryPatternId:form.ComplimentaryPatternId,ComplimentaryCash:form.ComplimentaryCash,PromotionType:form.PromotionType}
	if err := c.Create(&promotionactivity).Error ; err != nil{
		return e.ERROR
	}else{
		c.Last(&promotionactivity)
		if promotionactivity.ShopId == form.ShopId{
			var productrelation = models.ProductRelation{PromotionActivityId:int(promotionactivity.ID),CategoryId:categoryId,ProductId:productId,ShopId:form.ShopId}
			c.Table("t_product_relation").Create(&productrelation)
		}
		if form.ComplimentaryPatternId != 0{
			var complimentaryrelation = models.ProductRelation{PromotionActivityId:int(promotionactivity.ID),CategoryId:categoryId,ProductId:productId,ShopId:form.ShopId}
			c.Table("t_complimentary_relation").Create(&complimentaryrelation)
		}
		return e.SUCCESS
	}
}

func (c *PromotionActivity) DeleteById (promotionActivityId int) int  {
	results := &models.PromotionActivity{}
	c.Where("id = ?" ,promotionActivityId).Find(&results)
	if results.ComplimentaryPatternId != 0{
		if err := c.Where("promotion_activity_id = ?",promotionActivityId).Delete(&models.ComplimentaryRelation{}).Error; err!=nil{
			return e.ERROR
		}
	}
	if err :=c.Where("id = ?" , promotionActivityId).Delete(&models.PromotionActivity{}).Error; err!=nil{
	 	return e.ERROR
	 }
	if err :=c.Where("promotion_activity_id = ?",promotionActivityId).Delete(&models.ProductRelation{}).Error; err !=nil{
	 	return e.ERROR
	 }
	return e.SUCCESS
}

func (c *PromotionActivity) UpdatePomotionActivity (promotionActivityId int,form *models.PromotionActivityForm) int {
	var timeUnixb  int64
	timeUnixa :=time.Now().Unix()
	theTime, _ := time.ParseInLocation("2006-01-02 15:04:05",form.StartTime , time.Local)
	timeUnixb =theTime.Unix()
	categoryIda := strings.Replace(strings.Trim(fmt.Sprint(form.CategoryId), ""), "[", "", -1)
	categoryId := strings.Replace(categoryIda, "]", "", -1 )
	productIda := strings.Replace(strings.Trim(fmt.Sprint(form.ProductId), ""), "[", " ", -1)
	productId := strings.Replace(productIda, "]", "", -1 )
	result := &models.PromotionActivity{}
	c.First(result)
	result.ID = uint(promotionActivityId)
	result.StartTime=form.StartTime
	result.StopTime=form.StopTime
	result.PromotionPatternId=form.PromotionPatternId
	result.PromotionTheme=form.PromotionTheme
	result.PromotionDescribe=form.PromotionDescribe
	result.PromotionDiscount=form.PromotionDiscount
	result.PromotionCash=form.PromotionCash
	result.PromotionCount=form.PromotionCount
	result.ComplimentaryPatternId=form.ComplimentaryPatternId
	result.ComplimentaryCash=form.ComplimentaryCash
	var productrelation = models.ProductRelation{CategoryId:categoryId,ProductId:productId,ShopId:form.ShopId}
	var complimentaryrelation = models.ProductRelation{CategoryId:categoryId,ProductId:productId,ShopId:form.ShopId}
	//判断是限时活动还是不限时
	if form.StartTime != "0000-00-00 00:00:00"{
		//判断活动是否进行中
		fmt.Println("判断活动是否进行中")
		if timeUnixa < timeUnixb{
			fmt.Println("活动未开始可以进行修改")
			if err :=c.Table("t_promotion_activity").Save(&result).Error;err!=nil{
				return e.ERROR
			}
			if err :=c.Table("t_product_relation").Where("promotion_activity_id = ?",promotionActivityId).Update(&productrelation).Error;err!=nil {
				return e.ERROR
			}
			if form.ComplimentaryPatternId != 0{
				if err :=c.Table("t_complimentary_relation").Where("promotion_activity_id = ?",promotionActivityId).Update(&complimentaryrelation).Error;err!=nil {
					return e.ERROR
				}
			}
			return e.SUCCESS
		}
		return e.Time_ERROR
	}
	fmt.Println("不限时修改")
	if err :=c.Table("t_promotion_activity").Save(&result).Error;err!=nil{
		return e.ERROR
	}
	if err :=c.Table("t_product_relation").Where("promotion_activity_id = ?",promotionActivityId).Update(&productrelation).Error;err!=nil {
		return e.ERROR
	}
	if form.ComplimentaryPatternId != 0{
		if err :=c.Table("t_complimentary_relation").Where("promotion_activity_id = ?",promotionActivityId).Update(&complimentaryrelation).Error;err!=nil {
			return e.ERROR
		}
	}
	return e.SUCCESS
}


