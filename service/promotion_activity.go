package service

import "C"
import (
	"fmt"
	"gomall-center/models"
	"gomall-center/pkg/e"
	"gomall-center/pkg/web"
	"strings"
	"time"
)
//PromotionActivity ...
type PromotionActivity struct {
	Service
}

const (
	// 列表查询返回字段
	selectListField = "id,start_time,stop_time,shop_name,promotion_theme,promotion_describe"
)

//NewPromotionActivity ...
func NewPromotionActivity(cotent *web.RequestContext) *PromotionActivity  {
	c := &PromotionActivity{InitService(cotent)}
	return c
}

//InsertPromotionActivity ...添加促销活动（商家/平台）
func (c *PromotionActivity) InsertPromotionActivity (form *models.PromotionActivityForm) int {
	tx :=c.Begin()
	//去掉栏目数组，商品数组的[]
	categoryIda := strings.Replace(strings.Trim(fmt.Sprint(form.CategoryId), ""), "[", "", -1)
	categoryId := strings.Replace(categoryIda, "]", "", -1 )
	productIda := strings.Replace(strings.Trim(fmt.Sprint(form.ProductId), ""), "[", " ", -1)
	productId := strings.Replace(productIda, "]", "", -1 )
	//字段赋值
	var  promotionactivity = models.PromotionActivity{StartTime:form.StartTime,StopTime:form.StopTime,PromotionPatternId:form.PromotionPatternId,PromotionTheme:form.PromotionTheme,
		PromotionDescribe:form.PromotionDescribe,PromotionDiscount:form.PromotionDiscount,PromotionCash:form.PromotionCash,PromotionCount:form.PromotionCount,ShopId:form.ShopId,
		ShopName:form.ShopName,ComplimentaryPatternId:form.ComplimentaryPatternId,ComplimentaryCash:form.ComplimentaryCash,PromotionType:form.PromotionType}
	if err := c.Table("t_promotion_activity").Create(&promotionactivity).Error ; err != nil{
		tx.Rollback()
		return e.ERROR
	}else{
		//商品关系表
		c.Last(&promotionactivity)
		if promotionactivity.ShopId == form.ShopId{
			var productrelation = models.ProductRelation{PromotionActivityId:int(promotionactivity.ID),CategoryId:categoryId,ProductId:productId,ShopId:form.ShopId}
			if err := c.Table("t_product_relation").Create(&productrelation).Error; err!=nil{
				tx.Rollback()
				return e.ERROR
			}
		}
		//判断此活动中是否是优惠券活动,如果是添加增加优惠卷表,（商家如果有修改活动,已被用户领取的优惠卷需保留未改变之前的属性）
		if form.PromotionPatternId == 11 || form.PromotionPatternId == 12{
			var discountCoupon = models.DiscountCoupon{PromotionActivityId:int(promotionactivity.ID),StartTime:form.StartTime,StopTime:form.StopTime,
				PromotionPatternId:form.PromotionPatternId,PromotionDiscount:form.PromotionDiscount,PromotionCash:form.PromotionCash,PromotionType:form.PromotionType}
			for i := 1; i <= form.PromotionCount; i++ {
				if err := c.Table("t_discount_coupon").Create(&discountCoupon).Error; err!=nil{
					tx.Rollback()
					return e.ERROR
				}
			}
		}
		//判断此活动中是否存在赠品活动
		if form.ComplimentaryPatternId != 0{
			var complimentaryrelation = models.ProductRelation{PromotionActivityId:int(promotionactivity.ID),CategoryId:categoryId,ProductId:productId,ShopId:form.ShopId}
			if err := c.Table("t_complimentary_relation").Create(&complimentaryrelation).Error; err!=nil{
				tx.Rollback()
				return e.ERROR
			}
		}
		tx.Commit()
		return e.SUCCESS
	}
}

//InsertPromotionActivityAbs ...根据平台发布的活动id，获取发布活动的详情，在活动表中重新添加一行，附带上店铺的id，商家叠加平台活动
func (c *PromotionActivity) InsertPromotionActivityAbs (promotionActivityId int,shopId int) int {
	tx :=c.Begin()
	results := &models.PromotionActivity{}
	if err := c.Where("id = ?" ,promotionActivityId).First(&results).Error; err!=nil{
		return e.ERROR
	}
	//添加促销活动，附带店铺id
	var  promotionactivity = models.PromotionActivity{StartTime:results.StartTime,StopTime:results.StopTime,PromotionPatternId:results.PromotionPatternId,PromotionTheme:results.PromotionTheme,
		PromotionDescribe:results.PromotionDescribe,PromotionDiscount:results.PromotionDiscount,PromotionCash:results.PromotionCash,PromotionCount:results.PromotionCount,ShopId:shopId,
		ComplimentaryPatternId:results.ComplimentaryPatternId,ComplimentaryCash:results.ComplimentaryCash,PromotionType:results.PromotionType}
	if err := c.Table("t_promotion_activity").Create(&promotionactivity).Error ; err != nil{
		tx.Rollback()
		return e.ERROR
	}
	//平台只发布优惠券，叠加平台发布的活动，也需在优惠券表添加
	var discountCoupon = models.DiscountCoupon{PromotionActivityId:int(promotionactivity.ID),StartTime:results.StartTime,StopTime:results.StopTime,
		PromotionPatternId:results.PromotionPatternId,PromotionDiscount:results.PromotionDiscount,PromotionCash:results.PromotionCash,PromotionType:results.PromotionType}
	for i := 1; i <= results.PromotionCount; i++ {
		if err := c.Table("t_discount_coupon").Create(&discountCoupon).Error; err!=nil{
			tx.Rollback()
			return e.ERROR
		}
	}
	tx.Commit()
	return e.SUCCESS
}

//DeleteById ...根据活动id删除促销活动
func (c *PromotionActivity) DeleteById (promotionActivityId int) int  {
	tx :=c.Begin()
	results := &models.PromotionActivity{}
	db := c.Table("t_promotion_activity")
	if err := c.Where("id = ?" ,promotionActivityId).First(&results).Error; err!=nil{
		return e.ERROR
	}
	//促销活动表
	if err :=db.Where("id = ?" , promotionActivityId).Delete(&models.PromotionActivity{}).Error; err!=nil{
		tx.Rollback()
		return e.ERROR
	}
	//商品关系表
	if err :=db.Where("promotion_activity_id = ?",promotionActivityId).Delete(&models.ProductRelation{}).Error; err !=nil{
		tx.Rollback()
		return e.ERROR
	}
	//判断此活动中是否存在优惠券
	if results.PromotionPatternId != 0{
		if err :=db.Where("promotion_activity_id = ? and account_id = ?",promotionActivityId,0).Delete(&models.DiscountCoupon{}).Error; err!=nil{
			tx.Rollback()
			return e.ERROR
		}
	}
	//判断此活动中是否存在赠品
	if results.ComplimentaryPatternId != 0{
		if err := db.Where("promotion_activity_id = ?",promotionActivityId).Delete(&models.ComplimentaryRelation{}).Error; err!=nil{
			tx.Rollback()
			return e.ERROR
		}
	}
	tx.Commit()
	return e.SUCCESS
}

//PomotionActivity ...根据活动id修改促销活动
func (c *PromotionActivity) pomotionActivity(promotionActivityId int,form *models.PromotionActivityForm) int {
	tx :=c.Begin()
	categoryIda := strings.Replace(strings.Trim(fmt.Sprint(form.CategoryId), ""), "[", "", -1)
	categoryId := strings.Replace(categoryIda, "]", "", -1 )
	productIda := strings.Replace(strings.Trim(fmt.Sprint(form.ProductId), ""), "[", " ", -1)
	productId := strings.Replace(productIda, "]", "", -1 )
	result := &models.PromotionActivity{}
	c.First(result)
	count := result.PromotionCount
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
	//促销活动表
	if err :=c.Table("t_promotion_activity").Save(&result).Error;err!=nil{
		tx.Rollback()
		return e.ERROR
	}
	//商品关系表
	if err :=c.Table("t_product_relation").Where("promotion_activity_id = ?",promotionActivityId).Update(&productrelation).Error;err!=nil {
		tx.Rollback()
		return e.ERROR
	}
	//判断此活动中是否存在优惠券，存在优惠卷则更改优惠卷，已被领取的优惠卷不能进行修改,不存在优惠券，则删除优惠券表中的值
	db := c.Table("t_discount_coupon")
	if form.PromotionPatternId == 11 || form.PromotionPatternId == 12{
		var discountCoupon = models.DiscountCoupon{PromotionActivityId:promotionActivityId,StartTime:form.StartTime,StopTime:form.StopTime,
			PromotionPatternId:form.PromotionPatternId,PromotionDiscount:form.PromotionDiscount,PromotionCash:form.PromotionCash}
		//判断有没有更改优惠卷发行数量
		if count != form.PromotionCount{
			//更改了优惠卷发行数量的字段，需先删除未被用户领取的，重新在优惠券表中添加
			if err := db.Where("promotion_activity_id = ? and account_id = ?",promotionActivityId,0).Delete(&models.DiscountCoupon{}).Error; err!=nil{
				tx.Rollback()
				return e.ERROR
			}
			for i := 1; i <= form.PromotionCount; i++ {
				if err :=db.Create(&discountCoupon).Error;err!=nil {
					tx.Rollback()
					return e.ERROR
				}
			}
		}
		if err :=db.Where("promotion_activity_id = ? and account_id = ?",promotionActivityId,0).Update(&discountCoupon).Error; err!=nil{
			tx.Rollback()
			return e.ERROR
		}
	}else{
		//更改活动，从优惠券切换为别的活动，删除之前存储未被领取的优惠卷
		if err :=db.Where("promotion_activity_id = ? and account_id = ?",promotionActivityId,0).Delete(&models.DiscountCoupon{}).Error; err!=nil{
			tx.Rollback()
			return e.ERROR
		}
	}
	//判断活动中是否存在赠品,存在赠品则更改赠品，不存在赠品，则删除赠品表中的值
	if form.ComplimentaryPatternId != 0{
		if err :=c.Table("t_complimentary_relation").Where("promotion_activity_id = ?",promotionActivityId).Update(&complimentaryrelation).Error;err!=nil {
			tx.Rollback()
			return e.ERROR
		}
	}else{
		if err := c.Where("promotion_activity_id = ?",promotionActivityId).Delete(&models.ComplimentaryRelation{}).Error; err!=nil{
			tx.Rollback()
			return e.ERROR
		}
	}
	tx.Commit()
	return e.SUCCESS
}

//UpdatePomotionActivity ...根据活动id修改促销活动
func (c *PromotionActivity) UpdatePomotionActivity (promotionActivityId int,form *models.PromotionActivityForm) int {
	var timeUnixB int64
	timeUnixa :=time.Now().Unix()
	theTime, _ := time.ParseInLocation("2006-01-02 15:04:05",form.StartTime , time.Local)
	timeUnixB =theTime.Unix()
	//判断是限时活动还是不限时
	if form.StartTime != "0000-00-00 00:00:00"{
		//判断活动是否已进行,能否进行修改操作，判断条件放在页面中，查出活动详情的同时，就对其活动开始时间和当前时间进行比较判断
		if timeUnixa < timeUnixB {
			code :=c.pomotionActivity(promotionActivityId,form)
			return code
		}
		return e.Time_ERROR
	}
	code :=c.pomotionActivity(promotionActivityId,form)
	return code
}

func (c *PromotionActivity) findByProduct(productId string,categoryId string,shopId int) []*models.Products {
	products :=make([]*models.Products, 0)
	productIds := strings.Split(productId," ")
	categoryIds := strings.Split(productId," ")
	db := c.Table("t_product")
	if productId != ""{
		db.Where("id in (?)",productIds).Find(&products)
	} else if categoryId != ""{
		db.Where("category_id in (?)",categoryIds).Find(&products)
	} else if shopId != 0{
		db.Where("shop_id = ?",shopId).Find(&products)
	}
	return products
}

//FindPomotionActivityById ...根据活动id获取活动详情
func (c *PromotionActivity) FindPomotionActivityById (promotionActivityId int) *models.PromotionActivityAbs  {
	promotionActivity := &models.PromotionActivity{}
	productRelation := &models.ProductRelation{}
	complimentaryRelation := &models.ComplimentaryRelation{}
	products := make([]*models.Products, 0)
	complimentary := make([]*models.Products, 0)
	//判断该促销活动是否有促销商品
	c.Table("t_product_relation").Where("promotion_activity_id = ?",promotionActivityId).First(&productRelation)
	if productRelation.PromotionActivityId != 0 {
		products = c.findByProduct(productRelation.ProductId, productRelation.CategoryId, productRelation.ShopId)
	}
	//判断该促销活动是否有赠品
	c.Table("t_complimentary_relation").Where("promotion_activity_id = ?", promotionActivityId).First(&complimentaryRelation)
	if complimentaryRelation.PromotionActivityId != 0 {
		complimentary = c.findByProduct(complimentaryRelation.ProductId, complimentaryRelation.CategoryId, complimentaryRelation.ShopId)
	}
	c.Table("t_promotion_activity").Where("id = ?",promotionActivityId).First(&promotionActivity)
	results := &models.PromotionActivityAbs{}
	results.ID = uint(promotionActivityId)
	results.StartTime = promotionActivity.StartTime
	results.StopTime = promotionActivity.StopTime
	results.PromotionPatternId = promotionActivity.PromotionPatternId
	results.PromotionTheme = promotionActivity.PromotionTheme
	results.PromotionDescribe = promotionActivity.PromotionDescribe
	results.PromotionDiscount = promotionActivity.PromotionDiscount
	results.PromotionCash = promotionActivity.PromotionCash
	results.PromotionCount = promotionActivity.PromotionCount
	results.ComplimentaryPatternId = promotionActivity.ComplimentaryPatternId
	results.ComplimentaryCash = promotionActivity.ComplimentaryCash
	results.Products = products
	results.Complimentary = complimentary
	return results
}

//FindPomotionActivityList ...根据店铺状态id获取活动列表
func (c *PromotionActivity) FindPomotionActivityList (shopId int,args *models.PagerArgs) *models.PagerData  {
	var promotionActivity []*models.PromotionActivity
	var count int
	db := c.Table("t_promotion_activity")
	if args.KeyWord != "" {
		db.Where("promotion_theme like ?", "'%"+args.KeyWord+"%'")
	}
	if shopId != 0{
		db.Select(selectListField).Where("shop_id = ?",shopId)
		if err := db.Offset((args.PageNum - 1) * args.PageSize).Limit(args.PageSize).Find(&promotionActivity).Count(&count).Error ;err !=nil{
			return nil
		}
	}
	db.Select(selectListField).Where("promotion_type = ?",2)
	if err := db.Offset((args.PageNum - 1) * args.PageSize).Limit(args.PageSize).Find(&promotionActivity).Count(&count).Error ;err !=nil{
		return nil
	}
	results := &models.PagerData{
		Count: count,
		Data:  promotionActivity,
	}
	return results
}