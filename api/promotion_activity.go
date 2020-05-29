package api

import (
	"gomall-center/models"
	"gomall-center/pkg/e"
	"gomall-center/pkg/web"
	"gomall-center/service"
	"gopkg.in/go-playground/validator.v9"
	"strconv"
)

//AddPromotionActivity ...添加促销活动
func AddPromotionActivity (ctx *web.Context){
	form := &models.PromotionActivityForm{}
	if err := ctx.BindJSON(form);err != nil{
		ctx.Response(e.BAD_REQUEST)
		return
	}
	validate := validator.New()
	if err :=validate.Struct(form);err != nil{
		ctx.Response(e.BAD_REQUEST)
		return
	}
	srv := service.NewPromotionActivity(ctx.RequestContext)
	code := srv.InsertPromotionActivity(form)
	ctx.Response(code)
}

//AddPromotionActivityAbs ...叠加平台活动
func AddPromotionActivityAbs (ctx *web.Context){
	shopId,_ := strconv.Atoi(ctx.Query("shopId"))
	promotionActivityId,_ := strconv.Atoi(ctx.Param("id"))
	srv := service.NewPromotionActivity(ctx.RequestContext)
	code := srv.InsertPromotionActivityAbs(promotionActivityId,shopId)
	ctx.Response(code)
}

//RemovePromotionActivity ...删除促销活动
func RemovePromotionActivity(ctx *web.Context)  {
	promotionActivityId,_ := strconv.Atoi(ctx.Param("id"))
	srv := service.NewPromotionActivity(ctx.RequestContext)
	code := srv.DeleteById(promotionActivityId)
	ctx.Response(code)
}

//ModifyPromotionActivity ...修改促销活动
func ModifyPromotionActivity(ctx *web.Context)  {
	form := &models.PromotionActivityForm{}
	promotionActivityId,_ := strconv.Atoi(ctx.Param("id"))
	if err := ctx.BindJSON(form); err!=nil{
		ctx.Response(e.BAD_REQUEST)
		return
	}
	validate := validator.New()
	if  err := validate.Struct(form); err !=nil{
		ctx.Response(e.BAD_REQUEST)
		return
	}
	srv := service.NewPromotionActivity(ctx.RequestContext)
	code := srv.UpdatePromotionActivity(promotionActivityId,form)
	ctx.Response(code)
}

//FindByPromotionActivity ...获取促销活动详情
func FindByPromotionActivity(ctx *web.Context){
	promotionActivityId,_ := strconv.Atoi(ctx.Param("id"))
	srv :=service.NewPromotionActivity(ctx.RequestContext)
	data := srv.PromotionActivityById(promotionActivityId)
	ctx.ResponseData(e.SUCCESS,data)
}

//FindPromotionActivityList ...获取促销活动列表
func FindPromotionActivityList(ctx *web.Context){
	shopId,_ := strconv.Atoi(ctx.Param("id"))
	pager := &models.PagerArgs{}
	_ = pager.Bind(ctx)
	srv :=service.NewPromotionActivity(ctx.RequestContext)
	data := srv.PromotionActivityList(shopId,pager)
	ctx.ResponseData(e.SUCCESS,data)
}

//FindPromotionProductList ...获取促销活动商品
func FindPromotionProductList(ctx *web.Context){
	promotionPatternId,_ := strconv.Atoi(ctx.Query("id"))
	productType,_ := strconv.Atoi(ctx.Query("typeId"))
	pager := &models.PagerArgs{}
	_ = pager.Bind(ctx)
	srv :=service.NewPromotionActivity(ctx.RequestContext)
	data := srv.PromotionProductList(promotionPatternId,productType,pager)
	ctx.ResponseData(e.SUCCESS,data)
}
//FindPromotionProductList ...获取促销优惠券
func FindDiscountCouponList(ctx *web.Context){
	shopId,_ := strconv.Atoi(ctx.Query("id"))
	pager := &models.PagerArgs{}
	_ = pager.Bind(ctx)
	srv :=service.NewPromotionActivity(ctx.RequestContext)
	e,data := srv.DiscountCouponList(shopId,pager)
	ctx.ResponseData(e,data)
}