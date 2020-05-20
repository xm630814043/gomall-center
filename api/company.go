package api

import (
	"gomall-center/models"
	"gomall-center/pkg/e"
	"gomall-center/pkg/web"
	"gomall-center/service"
	"strconv"

	"gopkg.in/go-playground/validator.v9"
)

//FindCompanyList ...获取企业列表
func FindCompanyList(ctx *web.Context) {
	pager := &models.PagerArgs{}
	_ = pager.Bind(ctx)
	Radio, _ := strconv.Atoi(ctx.Query("radio"))
	pager.Filter = make(map[string]interface{})
	pager.Filter["company_status"] = Radio
	srv := service.NewCompanyService(ctx.RequestContext)
	data := srv.CompanyList(pager)
	ctx.ResponseData(e.SUCCESS, data)
}

//FindByCompany ...获取企业详情
func FindByCompany(ctx *web.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	srv := service.NewCompanyService(ctx.RequestContext)
	data := srv.CompanyByID(id)
	ctx.ResponseData(e.SUCCESS, data)
}

//ModifyCompany ...修改企业状态
func ModifyCompany(ctx *web.Context) {
	form := &models.CompanyPut{}
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err := ctx.BindJSON(form); err != nil {
		ctx.Response(e.BAD_REQUEST)
		return
	}
	validate := validator.New()
	if err := validate.Struct(form); err != nil {
		ctx.Response(e.BAD_REQUEST)
		return
	}
	srv := service.NewCompanyService(ctx.RequestContext)
	data := srv.UpdateCompany(id, form)
	ctx.ResponseData(e.SUCCESS, data)
}
