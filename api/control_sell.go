package api

import (
	"gomall-center/models"
	"gomall-center/pkg/e"
	"gomall-center/pkg/web"
	"gomall-center/service"
	"gopkg.in/go-playground/validator.v9"
	"strconv"
)

//InsertControlSell ...添加控销模板
func AddControlSell (ctx *web.Context){
	form := &models.ControlSells{}
	if err := ctx.BindJSON(form); err != nil {
		ctx.Response(e.BAD_REQUEST)
		return
	}
	validate := validator.New()
	if err := validate.Struct(form) ; err != nil{
		ctx.Response(e.BAD_REQUEST)
		return
	}
	srv := service.NewControlSellService(ctx.RequestContext)
	code := srv.InsertControlSell(form)
	ctx.Response(code)
}

//InsertControlSell ...删除控销模板
func RemoveControlSell(ctx *web.Context){
	controlSellId,_ := strconv.Atoi(ctx.Param("id"))
	srv := service.NewControlSellService(ctx.RequestContext)
	code := srv.DeleteByID(controlSellId)
	ctx.Response(code)
}

//InsertControlSell ...修改控销模板
func ModifyControlSell(ctx *web.Context){
	form := &models.ControlSells{}
	controlSellId,_ := strconv.Atoi(ctx.Param("id"))
	if err := ctx.BindJSON(form) ; err != nil{
		ctx.Response(e.BAD_REQUEST)
		return
	}
	validate := validator.New()
	if err := validate.Struct(form); err != nil{
		ctx.Response(e.BAD_REQUEST)
		return
	}
	srv := service.NewControlSellService(ctx.RequestContext)
	code := srv.UpdateControlSell(controlSellId,form)
	ctx.Response(code)
}

//InsertControlSell ...获取控销模板详情
func FindByControlSell(ctx *web.Context){
	controlSellId,_ := strconv.Atoi(ctx.Param("id"))
	srv :=service.NewControlSellService(ctx.RequestContext)
	data := srv.ControlSellByID(controlSellId)
	ctx.ResponseData(e.SUCCESS,data)
}

//InsertControlSell ...获取控销模板列表
func FindControlSellList(ctx *web.Context){
	companyId,_ := strconv.Atoi(ctx.Param("id"))
	pager := &models.PagerArgs{}
	_ = pager.Bind(ctx)
	srv :=service.NewControlSellService(ctx.RequestContext)
	data := srv.ControlSellList(companyId,pager)
	ctx.ResponseData(e.SUCCESS,data)
}