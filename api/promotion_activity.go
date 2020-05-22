package api

import (
	"gomall-center/models"
	"gomall-center/pkg/e"
	"gomall-center/pkg/web"
	"gomall-center/service"
	"gopkg.in/go-playground/validator.v9"
	"strconv"
)

func AddPomotionActivity (ctx *web.Context){
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

func RemovePomotionActivity(ctx *web.Context)  {
	promotionActivityId,_ := strconv.Atoi(ctx.Param("id"))
	srv := service.NewPromotionActivity(ctx.RequestContext)
	code := srv.DeleteById(promotionActivityId)
	ctx.Response(code)

}

func ModifyPomotionActivity(ctx *web.Context)  {
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
	code := srv.UpdatePomotionActivity(promotionActivityId,form)
	ctx.Response(code)
}