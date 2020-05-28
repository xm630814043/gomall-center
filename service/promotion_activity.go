package service

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
	selectListField = "id,start_time,stop_time,shop_id,shop_name,promotion_theme,promotion_describe"
)
//NewPromotionActivity ...
func NewPromotionActivity(content *web.RequestContext) *PromotionActivity  {
	c := &PromotionActivity{InitService(content)}
	return c
}
//InsertPromotionActivity ...添加促销活动（商家/平台）
func (c *PromotionActivity) InsertPromotionActivity (form *models.PromotionActivityForm) int {
	id := c.strReplace(form)
	var  promotionActivity = models.PromotionActivity{StartTime:form.StartTime,StopTime:form.StopTime,PromotionPatternId:form.PromotionPatternId,PromotionTheme:form.PromotionTheme,
		PromotionDescribe:form.PromotionDescribe,PromotionDiscount:form.PromotionDiscount,PromotionCash:form.PromotionCash,PromotionCount:form.PromotionCount,ShopId:form.ShopId,
		ShopName:form.ShopName,ComplimentaryPatternId:form.ComplimentaryPatternId,ComplimentaryCash:form.ComplimentaryCash,PromotionType:form.PromotionType}
	if err := c.insertPromotionActivity(promotionActivity); err != e.SUCCESS{
		return e.ERROR
	}else{
		//商品关系表
		c.Last(&promotionActivity)
		if promotionActivity.ShopId == form.ShopId{
			var productRelation = models.ProductRelation{PromotionActivityId:int(promotionActivity.ID),CategoryId:id[0],ProductId:id[1],ShopId:form.ShopId}
			if err := c.insertProductRelation(productRelation); err != e.SUCCESS{
				return e.ERROR
			}
		}
		//判断此活动中是否是优惠券活动
		if form.PromotionPatternId == 11 || form.PromotionPatternId == 12{
			var discountCoupon = models.DiscountCoupon{PromotionActivityId:int(promotionActivity.ID),StartTime:form.StartTime,StopTime:form.StopTime,
				PromotionPatternId:form.PromotionPatternId,PromotionDiscount:form.PromotionDiscount,ShopId:form.ShopId,PromotionCash:form.PromotionCash,PromotionType:form.PromotionType}
			counts :=promotionActivity.PromotionCount
			if err := c.insertDiscountCoupon(discountCoupon,counts); err != e.SUCCESS{
				return e.ERROR
			}
		}
		//判断此活动中是否存在赠品活动
		if form.ComplimentaryPatternId != 0{
			var complimentaryRelation = models.ComplimentaryRelation{PromotionActivityId:int(promotionActivity.ID),CategoryId:id[0],ProductId:id[1],ShopId:form.ShopId}
			if err := c.insertComplimentaryPattern(complimentaryRelation); err != e.SUCCESS{
				return e.ERROR
			}
		}
		return e.SUCCESS
	}
}
//InsertPromotionActivityAbs ...InsertPromotionActivityAbs
func (c *PromotionActivity) InsertPromotionActivityAbs (promotionActivityId int,shopId int) int {
	results := &models.PromotionActivity{}
	if err := c.Where("id = ?" ,promotionActivityId).First(&results).Error; err!=nil{
		return e.ERROR
	}
	//添加平台促销活动，附带店铺id
	var  promotionActivity = models.PromotionActivity{StartTime: results.StartTime,StopTime:results.StopTime,PromotionPatternId:results.PromotionPatternId,PromotionTheme:results.PromotionTheme,
		PromotionDescribe:results.PromotionDescribe,PromotionDiscount:results.PromotionDiscount,PromotionCash:results.PromotionCash,PromotionCount:results.PromotionCount,ShopId:shopId,
		ComplimentaryPatternId:results.ComplimentaryPatternId,ComplimentaryCash:results.ComplimentaryCash,PromotionType:results.PromotionType,PromotionActivityId:int(results.ID)}
	if err := c.insertPromotionActivity(promotionActivity); err != e.SUCCESS{
		return e.ERROR
	}
	return e.SUCCESS
}
//DeleteById ...根据促销活动id删除
func (c *PromotionActivity) DeleteById (promotionActivityId int) int  {
	results := &models.PromotionActivity{}
	db := c.Table("t_promotion_activity")
	if err := db.Where("id = ?" ,promotionActivityId).First(&results).Error; err!=nil{
		return e.ERROR
	}
	//判断是否存在活动
	if results.PromotionPatternId != 0{
		//判断是不是优惠券活动？是:删除(优惠券表)
		if results.PromotionPatternId == 11 || results.PromotionPatternId == 12{
			if err := c.deleteDiscountCoupon(promotionActivityId); err != e.SUCCESS{
				return e.ERROR
			}
		}else{
			//不是优惠券活动，那就是直减和满减活动，删除（商品关系表）
			if err := c.deleteProductRelation(promotionActivityId); err != e.SUCCESS{
				return e.ERROR
			}
		}
		//判断是否存在赠品活动,是:删除(赠品关系表)
		if results.ComplimentaryPatternId != 0{
			if err := c.deleteComplimentaryPattern(promotionActivityId); err != e.SUCCESS{
				return e.ERROR
			}
		}
		if err := c.deletePromotionActivity(promotionActivityId); err != e.SUCCESS{
			return e.ERROR
		}
	}
	if err := c.deletePromotionActivity(promotionActivityId); err != e.SUCCESS{
		return e.ERROR
	}
	return e.SUCCESS
}
//PromotionActivity ...UpdatePromotionActivity调用的函数
func (c *PromotionActivity) promotionActivity(promotionActivityId int,form *models.PromotionActivityForm) int {
	tx :=c.Begin()
	id := c.strReplace(form)
	result := &models.PromotionActivity{}
	c.First(result)
	promotionPatternIda := result.PromotionPatternId
	complimentaryPatternIda := result.ComplimentaryPatternId
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
	var productRelation = models.ProductRelation{PromotionActivityId:int(result.ID),CategoryId:id[0],ProductId:id[1],ShopId:form.ShopId}
	var complimentaryRelation = models.ComplimentaryRelation{PromotionActivityId:int(result.ID),CategoryId:id[0],ProductId:id[1],ShopId:form.ShopId}
	//促销活动表
	if err :=c.Table("t_promotion_activity").Where("id = ?",promotionActivityId).Update(&result).Error ; err != nil {
		tx.Rollback()
		return e.ERROR
	}
	var discountCoupon = models.DiscountCoupon{PromotionActivityId:promotionActivityId,StartTime:form.StartTime,StopTime:form.StopTime,PromotionPatternId:form.PromotionPatternId,
		PromotionDiscount:form.PromotionDiscount,PromotionCash:form.PromotionCash,ShopId:form.ShopId,PromotionType:form.PromotionType}
	counts := form.PromotionCount
	//判断未修改前是什么活动
	if promotionPatternIda == 11 || promotionPatternIda == 12{
		if form.PromotionPatternId == 11 || form.PromotionPatternId == 12{
			db := c.Table("t_discount_coupon")
			//判断有没有更改优惠卷发行数量
			if count != form.PromotionCount{
				//更改了优惠卷发行数量的字段，需先删除未被用户领取的，重新在优惠券表中添加
				if err := c.deleteDiscountCoupon(promotionActivityId); err != e.SUCCESS{
					return e.ERROR
				}
				if err := c.insertDiscountCoupon(discountCoupon,counts); err != e.SUCCESS{
					return e.ERROR
				}
			}else{
				if err :=db.Where("promotion_activity_id = ? and account_id = ?",promotionActivityId,0).Update(&discountCoupon).Error; err!=nil{
					tx.Rollback()
					return e.ERROR
				}
			}
		}else{
			//更改活动，从优惠券切换为别的活动，删除之前存储未被领取的优惠卷
			if err := c.deleteDiscountCoupon(promotionActivityId); err != e.SUCCESS{
				return e.ERROR
			}
			if err := c.insertProductRelation(productRelation); err != e.SUCCESS{
				return e.ERROR
			}
		}
	}else{
		if form.PromotionPatternId == 11 || form.PromotionPatternId == 12{
			if err := c.deleteProductRelation(promotionActivityId); err != e.SUCCESS{
				return e.ERROR
			}
			if err := c.insertDiscountCoupon(discountCoupon,counts); err != e.SUCCESS{
				return e.ERROR
			}
		}else{
			if err :=c.Table("t_product_relation").Where("promotion_activity_id = ?",promotionActivityId).Update(&productRelation).Error;err!=nil {
				tx.Rollback()
				return e.ERROR
			}
		}
	}
	//判断活动中是否存在赠品,存在赠品则更改赠品，不存在赠品，则删除赠品表中的值
	if complimentaryPatternIda !=0{
		if form.ComplimentaryPatternId != 0{
			if err :=c.Table("t_complimentary_relation").Where("promotion_activity_id = ?",promotionActivityId).Update(&complimentaryRelation).Error;err!=nil {
				tx.Rollback()
				return e.ERROR
			}
		}else{
			if err := c.deleteComplimentaryPattern(promotionActivityId); err != e.SUCCESS{
				return e.ERROR
			}
		}
	}else{
		if form.ComplimentaryPatternId != 0{
			if err := c.insertComplimentaryPattern(complimentaryRelation); err != e.SUCCESS{
				return e.ERROR
			}
		}
	}
	tx.Commit()
	return e.SUCCESS
}
//UpdatePromotionActivity ...根据活动id修改促销活动
func (c *PromotionActivity) UpdatePromotionActivity (promotionActivityId int,form *models.PromotionActivityForm) int {
	promotionActivity := &models.PromotionActivity{}
	if err := c.Where("id = ?",promotionActivityId).First(&promotionActivity).Error;err!=nil{
		return e.ERROR
	}
	startTimes,flag := c.timeBefore(promotionActivity.StartTime)
	//判断是限时活动还是不限时
	if startTimes != "0001-01-01 00:00:00"{
		//判断活动是否已进行,能否进行修改操作
		if flag {
			code :=c.promotionActivity(promotionActivityId,form)
			return code
		}
		return e.Time_ERROR
	}
	code := c.promotionActivity(promotionActivityId,form)
	return code
}
//FindPromotionActivityById ...根据活动id获取活动详情
func (c *PromotionActivity) PromotionActivityById(promotionActivityId int) *models.PromotionActivityAbs  {
	promotionActivity := &models.PromotionActivity{}
	//判断该促销活动是否存在促销商品
	promotionProduct := c.findPromotionActivityById(promotionActivityId)
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
	results.Products = promotionProduct[0]
	results.Complimentary = promotionProduct[1]
	return results
}
//FindPromotionActivityList ...根据店铺id获取活动列表
func (c *PromotionActivity) PromotionActivityList(shopId int,args *models.PagerArgs) *models.PagerData  {
	var promotionActivity []*models.PromotionActivity
	var count int
	db := c.Table("t_promotion_activity").Select(selectListField)
	if args.KeyWord != "" {
		db = db.Where("promotion_theme like ?", "%"+args.KeyWord+"%")
	}
	if shopId != 0{
		db = db.Where("shop_id = ? and deleted_at is null ",shopId)
	}else{
		db = db.Where("promotion_type = ? and deleted_at is null ",2)
	}
	db.Count(&count)
	db.Offset((args.PageNum - 1) * args.PageSize).Limit(args.PageSize).Find(&promotionActivity)
	results := &models.PagerData{
		Count: count,
		Data:  promotionActivity,
	}
	return results
}
//FindPromotionActivityList ...根据活动id获取商品列表
func (c *PromotionActivity) PromotionProductList(promotionActivityId int,args *models.PagerArgs) *models.PagerData  {
	promotionActivity := &models.PromotionActivity{}
	var count int
	c.Table("t_promotion_activity").Where("id = ?",promotionActivityId).First(&promotionActivity)
	stopTimes,flag := c.timeBefore(promotionActivity.StopTime)
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
	errors := &models.Products{
		Id:0,
		ProductName:"活动已过期",
	}
	if stopTimes != "0001-01-01 00:00:00" {
		fmt.Println(stopTimes,flag)
		if flag {
			promotionProduct := c.findPromotionActivityById(promotionActivityId)
			if len(promotionProduct[0]) > 0{
				results.Products = promotionProduct[0]
				count = len(promotionProduct[0])
			}else{
				promotionProduct[0] = append(promotionProduct[0],errors)
				results.Products = promotionProduct[0]
			}
			results.Complimentary = promotionProduct[1]
		}else{
			return nil
		}
	}else{
		promotionProduct := c.findPromotionActivityById(promotionActivityId)
		if len(promotionProduct[0]) > 0{
			results.Products = promotionProduct[0]
			count = len(promotionProduct[0])
		}else{
			promotionProduct[0] = append(promotionProduct[0],errors)
			results.Products = promotionProduct[0]
		}
		results.Complimentary = promotionProduct[1]
	}
	fmt.Println(results)
	result := &models.PagerData{
		Count: count,
		Data:  results,
	}
	return result
}
//DiscountCouponList ...根据活动id获取优惠券列表
func (c *PromotionActivity) DiscountCouponList(shopId int,args *models.PagerArgs) (int,*models.PagerData)  {
	var promotionActivity []*models.PromotionActivity
	var discountCoupon  []*models.DiscountCoupon
	var count int
	var promotionActivityId []int
	//获取促销活动表中店铺创建的所有优惠券活动
	db := c.Table("t_discount_coupon").Select("distinct promotion_activity_id,stop_time")
	if shopId != 0{
		 db = db.Where("shop_id = ? and account_id = ?",shopId,0)
	}else{
		db = db.Where("promotion_type = ? and account_id = ?",2,0)
	}
	db = db.Find(&discountCoupon)
	//根据优惠券表中的shopId获取优惠券表中信息
	for _,f :=range discountCoupon{
		stopTimes,flag := c.timeBefore(f.StopTime)
		if stopTimes != "0001-01-01 00:00:00" {
			if flag {
				promotionActivityId = append(promotionActivityId, f.PromotionActivityId)
			}
		}else{
			promotionActivityId = append(promotionActivityId, f.PromotionActivityId)
		}
	}
	//根据促销活动表中的id获取优惠券描述
	dba := c.Table("t_promotion_activity")
	if args.KeyWord != "" {
		dba = dba.Where("promotion_theme like ?", "%"+args.KeyWord+"%")
	}
	dba = dba.Where("id in (?) and deleted_at is null ",promotionActivityId)
	dba.Count(&count)
	dba.Offset((args.PageNum - 1) * args.PageSize).Limit(args.PageSize).Find(&promotionActivity)
	results := &models.PagerData{
		Count: count,
		Data:  promotionActivity,
	}
	return e.SUCCESS,results
}
//insertPromotionActivity ...添加促销活动
func (c *PromotionActivity) insertPromotionActivity (promotionActivity models.PromotionActivity) int {
	if err := c.Table("t_promotion_activity").Create(&promotionActivity).Error ; err != nil{
		return e.ERROR
	}
	return e.SUCCESS
}
//insertProductRelation ...添加商品关系
func (c *PromotionActivity) insertProductRelation (productRelation models.ProductRelation) int {
	if err := c.Table("t_product_relation").Create(&productRelation).Error; err!=nil{
		return e.ERROR
	}
	return e.SUCCESS
}
//insertDiscountCoupon ...添加优惠券
func (c *PromotionActivity) insertDiscountCoupon (discountCoupon models.DiscountCoupon,counts int) int {
	for i := 1; i <= counts; i++ {
		if err := c.Table("t_discount_coupon").Create(&discountCoupon).Error; err!=nil{
			return e.ERROR
		}
	}
	return e.SUCCESS
}
//insertComplimentaryPattern ...添加赠品
func (c *PromotionActivity) insertComplimentaryPattern (complimentaryRelation models.ComplimentaryRelation) int {
	if err := c.Table("t_complimentary_relation").Create(&complimentaryRelation).Error; err!=nil{
		return e.ERROR
	}
	return e.SUCCESS
}
//deleteProductRelation ...删除促销活动
func (c *PromotionActivity) deletePromotionActivity (promotionActivityId int) int  {
	if err := c.Where("id = ?" ,promotionActivityId).Delete(&models.PromotionActivity{}).Error; err!=nil{
		return e.ERROR
	}
	return e.SUCCESS
}
//deleteProductRelation ...删除商品关系
func (c *PromotionActivity) deleteProductRelation (promotionActivityId int) int  {
	if err :=c.Where("promotion_activity_id = ?",promotionActivityId).Delete(&models.ProductRelation{}).Error; err !=nil{
		return e.ERROR
	}
	return e.SUCCESS
}
//deleteDiscountCoupon ...删除优惠券
func (c *PromotionActivity) deleteDiscountCoupon (promotionActivityId int) int  {
	if err :=c.Where("promotion_activity_id = ? and account_id = ?",promotionActivityId,0).Delete(&models.DiscountCoupon{}).Error; err!=nil{
		return e.ERROR
	}
	return e.SUCCESS
}
//deleteComplimentaryPattern ...删除赠品
func (c *PromotionActivity) deleteComplimentaryPattern (promotionActivityId int) int  {
	if err := c.Where("promotion_activity_id = ?",promotionActivityId).Delete(&models.ComplimentaryRelation{}).Error; err!=nil{
		return e.ERROR
	}
	return e.SUCCESS
}
//findByProduct ...获取活动商品/赠品列表
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
//findPromotionActivityById ...根据活动id获取商品/赠品
func (c *PromotionActivity)  findPromotionActivityById (promotionActivityId int) [2][]*models.Products{
	var promotionProduct [2][]*models.Products
	products := make([]*models.Products, 0)
	complimentary := make([]*models.Products, 0)
	productRelation := &models.ProductRelation{}
	complimentaryRelation := &models.ComplimentaryRelation{}
	c.Table("t_product_relation").Where("promotion_activity_id = ?",promotionActivityId).First(&productRelation)
	if productRelation.PromotionActivityId != 0 {
		products = c.findByProduct(productRelation.ProductId, productRelation.CategoryId, productRelation.ShopId)
		promotionProduct[0] = products
	}
	//判断该促销活动是否存在赠品
	c.Table("t_complimentary_relation").Where("promotion_activity_id = ?", promotionActivityId).First(&complimentaryRelation)
	if complimentaryRelation.PromotionActivityId != 0 {
		complimentary = c.findByProduct(complimentaryRelation.ProductId, complimentaryRelation.CategoryId, complimentaryRelation.ShopId)
		promotionProduct[1] = complimentary
	}
	return promotionProduct
}
//strReplace ...商品数组转换字符串
func (c *PromotionActivity) strReplace(form *models.PromotionActivityForm) []string  {
	var id []string
	categoryIda := strings.Replace(strings.Trim(fmt.Sprint(form.CategoryId), ""), "[", "", -1)
	categoryId := strings.Replace(categoryIda, "]", "", -1 )
	id = append(id, categoryId)
	productIda := strings.Replace(strings.Trim(fmt.Sprint(form.ProductId), ""), "[", " ", -1)
	productId := strings.Replace(productIda, "]", "", -1 )
	id = append(id, productId)
	return id
}
//timeBefore ...mysql,datetime类型转换time.Time
func (c *PromotionActivity) timeBefore(times string) (string,bool){
	var startTimes []string
	startTime := strings.Replace(times, "T", " ", -1 )
	start :=strings.Index(startTime, "Z")
	if start != -1{
		startTimes = strings.Split(startTime,"Z")
	}else{
		startTimes = strings.Split(startTime,"+")
	}
	theTime, _ := time.ParseInLocation("2006-01-02 15:04:05",startTimes[0] , time.Local)
	flag := time.Now().Before(theTime)
	return startTimes[0],flag
}