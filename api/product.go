package api

import (
	"gomall-center/models"
	"gomall-center/pkg/e"
	"gomall-center/pkg/web"
	"gomall-center/service"
	"strconv"

	"gopkg.in/go-playground/validator.v9"
)

//FindProductBasicByStatus ...根据发布状态获取产品基本信息
func FindProductBasicByStatus(ctx *web.Context) {
	pager := &models.PagerArgs{}
	_ = pager.Bind(ctx)
	status, _ := strconv.Atoi(ctx.Query("radio"))
	srv := service.NewProductService(ctx.RequestContext)
	data := srv.ProductBasicByStatus(pager, status)
	ctx.ResponseData(e.SUCCESS, data)
}

//FindProductDetails ...查找产品详情
func FindProductDetails(ctx *web.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	srv := service.NewProductService(ctx.RequestContext)
	data := srv.ProductDetails(id)
	ctx.ResponseData(e.SUCCESS, data)
}

//ModifyProductStatus ...更改商品发布状态
func ModifyProductStatus(ctx *web.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	product := &models.ProductPutStatic{}
	if err := ctx.BindJSON(product); err != nil {
		ctx.Response(e.BAD_REQUEST)
		return
	}
	validate := validator.New()
	if err := validate.Struct(product); err != nil {
		ctx.Response(e.BAD_REQUEST)
		return
	}
	srv := service.NewProductService(ctx.RequestContext)
	code := srv.UpdateProductStatus(id, product)
	ctx.Response(code)
}

//FindPropValueByPropID ...根据属性ID获取属性值
func FindPropValueByPropID(ctx *web.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	pager := &models.PagerArgs{}
	_ = pager.Bind(ctx)
	srv := service.NewProductService(ctx.RequestContext)
	data := srv.PropValue(id, pager)
	ctx.ResponseData(e.SUCCESS, data)
}

//AddPropValByPropID ...根据属性ID增加属性值
func AddPropValByPropID(ctx *web.Context) {
	propVal := &models.ProductPropValue{}
	if err := ctx.BindJSON(propVal); err != nil {
		ctx.Response(e.BAD_REQUEST)
		return
	}
	validate := validator.New()
	if err := validate.Struct(propVal); err != nil {
		ctx.Response(e.BAD_REQUEST)
		return
	}
	srv := service.NewProductService(ctx.RequestContext)
	code := srv.InsertPropValue(propVal)
	ctx.Response(code)
}

//RemovePropValue ...删除属性的值根据ID
func RemovePropValue(ctx *web.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	srv := service.NewProductService(ctx.RequestContext)
	code := srv.DeletePropValue(id)
	ctx.Response(code)
}

//AddProp ...添加一个属性
func AddProp(ctx *web.Context) {
	propForm := &models.PropForm{}
	if err := ctx.BindJSON(propForm); err != nil {
		ctx.Response(e.BAD_REQUEST)
		return
	}
	validate := validator.New()
	if err := validate.Struct(propForm); err != nil {
		ctx.Response(e.BAD_REQUEST)
		return
	}
	proptype, _ := strconv.Atoi(propForm.PropType)
	inputType, _ := strconv.Atoi(propForm.InputType)
	prop := &models.ProductProp{
		ID:        0,
		PropName:  propForm.PropName,
		PropType:  proptype,
		ShopID:    1,
		InputType: inputType,
		IsDefault: 0,
	}
	categoryID, _ := strconv.Atoi(ctx.Query("categoryId"))
	parentID, _ := strconv.Atoi(ctx.Query("parentId"))
	srv := service.NewProductService(ctx.RequestContext)
	code := srv.InsertProp(categoryID, parentID, prop)
	ctx.Response(code)
}

//ModifyPropValue ...修改属性的值
func ModifyPropValue(ctx *web.Context) {
	propVal := &models.ProductPropValue{}
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := ctx.BindJSON(propVal); err != nil {
		ctx.Response(e.BAD_REQUEST)
		return
	}
	validate := validator.New()
	if err := validate.Struct(propVal); err != nil {
		ctx.Response(e.BAD_REQUEST)
		return
	}
	srv := service.NewProductService(ctx.RequestContext)
	code := srv.UpdatePropValue(propVal, id)
	ctx.Response(code)
}

//ModifyProp ...修改
func ModifyProp(ctx *web.Context) {
	propForm := &models.PropForm{}
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := ctx.BindJSON(propForm); err != nil {
		ctx.Response(e.BAD_REQUEST)
		return
	}
	validate := validator.New()
	if err := validate.Struct(propForm); err != nil {
		ctx.Response(e.BAD_REQUEST)
		return
	}
	inputtype, _ := strconv.Atoi(propForm.InputType)
	proptype, _ := strconv.Atoi(propForm.PropType)
	prop := &models.ProductProp{
		PropName:  propForm.PropName,
		PropType:  proptype,
		InputType: inputtype,
	}
	srv := service.NewProductService(ctx.RequestContext)
	code := srv.UpdateProp(prop, id)
	ctx.Response(code)
}

//RemoveProp ... 删除一个属性根据ID
func RemoveProp(ctx *web.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	categoryID, _ := strconv.Atoi(ctx.Query("categoryId"))
	parentID, _ := strconv.Atoi(ctx.Query("parentId"))
	srv := service.NewProductService(ctx.RequestContext)
	code := srv.DeleteProp(id, categoryID, parentID)
	ctx.Response(code)
}
