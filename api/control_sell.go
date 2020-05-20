package api

import (
	"fmt"
	"gomall-center/models"
	"gomall-center/pkg/e"
	"gomall-center/pkg/web"
	"gomall-center/service"
	"gopkg.in/go-playground/validator.v9"
	"strconv"
)

//InsertControlSell ...添加促销模板
func AddControlSell (ctx *web.Context){
	form := &models.ControlSells{}
	if err := ctx.BindJSON(form); err != nil {
		ctx.Response(e.BAD_REQUEST)
		fmt.Println("错误添加")
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

//InsertControlSell ...删除促销模板
func RemoveControlSell(ctx *web.Context){
	controlsellId,_ := strconv.Atoi(ctx.Param("id"))
	srv := service.NewControlSellService(ctx.RequestContext)
	code := srv.DeleteByID(controlsellId)
	ctx.Response(code)
}

//InsertControlSell ...修改促销模板
func ModifyControlSell(ctx *web.Context){
	form := &models.ControlSells{}
	controlsellId,_ := strconv.Atoi(ctx.Param("id"))
	fmt.Println("接受到传参的id",controlsellId)
	if err := ctx.BindJSON(form) ; err != nil{
		ctx.Response(e.BAD_REQUEST)
		fmt.Println("错误修改")
		return
	}
	validate := validator.New()
	if err := validate.Struct(form); err != nil{
		ctx.Response(e.BAD_REQUEST)
		return
	}
	fmt.Println(form)
	srv := service.NewControlSellService(ctx.RequestContext)
	code := srv.UpdateControlSell(controlsellId,form)
	ctx.Response(code)
}

//InsertControlSell ...获取促销模板详情
func FindControlSellByID(ctx *web.Context){
	controlsellId,_ := strconv.Atoi(ctx.Param("id"))
	srv :=service.NewControlSellService(ctx.RequestContext)
	data := srv.ControlSellByID(controlsellId)
	ctx.ResponseData(e.SUCCESS,data)
}

//InsertControlSell ...获取促销模板列表
func FindControlSellList(ctx *web.Context){
	companyId,_ := strconv.Atoi(ctx.Param("id"))
	pager := &models.PagerArgs{}
	_ = pager.Bind(ctx)
	srv :=service.NewControlSellService(ctx.RequestContext)
	data := srv.ControlSellList(companyId,pager)
	ctx.ResponseData(e.SUCCESS,data)
}