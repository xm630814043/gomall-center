package api

import (
	"gomall-center/models"
	"gomall-center/pkg/e"
	"gomall-center/pkg/web"
	"gomall-center/service"
	"strconv"
)

//FindAccountList ...获取用户列表
func FindAccountList(ctx *web.Context) {
	pager := &models.PagerArgs{}
	_ = pager.Bind(ctx)
	Radio, _ := strconv.Atoi(ctx.Query("radio"))
	srv := service.NewAccountService(ctx.RequestContext)
	data := srv.AccountList(Radio,pager )
	ctx.ResponseData(e.SUCCESS, data)
}

//FindByAccount ...获取用户详情
func FindByAccount(ctx *web.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	srv := service.NewAccountService(ctx.RequestContext)
	data := srv.AccountByID(int(id))
	ctx.ResponseData(e.SUCCESS, data)
}
