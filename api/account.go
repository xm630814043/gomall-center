package api

import (
	"gomall-center/models"
	"gomall-center/pkg/e"
	"gomall-center/pkg/web"
	"gomall-center/service"
)

// CompanyList 获取企业列表
func CompanyList(ctx *web.Context) {
	pager := &models.PagerArgs{}
	_ = pager.Bind(ctx)
	srv := service.NewCompanyService(ctx.RequestContext)
	data := srv.GetCompanyList(pager)
	ctx.ResponseData(e.SUCCESS, data)
}
