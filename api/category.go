package api

import (
	"gomall-center/models"
	"gomall-center/pkg/e"
	"gomall-center/pkg/web"
	"gomall-center/service"
	"strconv"
)

//FindCategoryList ... 获取商品种类树
func FindCategoryList(ctx *web.Context) {
	srv := service.NewCategoryService(ctx.RequestContext)
	data := srv.CategoryList()
	ctx.ResponseData(e.SUCCESS, data)
}

//FindPropByCategoryID ...根据种类ID获取属性
func FindPropByCategoryID(ctx *web.Context) {
	id := ctx.Param("id")
	pager := &models.PagerArgs{}
	_ = pager.Bind(ctx)
	srv := service.NewCategoryService(ctx.RequestContext)
	categoryID, _ := strconv.Atoi(id)
	if categoryID == 0 {
		data := srv.PropList(pager)
		ctx.ResponseData(e.SUCCESS, data)
		return
	}
	data := srv.GetPropList(categoryID, pager)
	ctx.ResponseData(e.SUCCESS, data)
}
