package api

import (
	"gomall-center/models"
	"gomall-center/pkg/e"
	"gomall-center/pkg/web"
	"gomall-center/service"
	"gopkg.in/go-playground/validator.v9"
	"strconv"
)

//AddArticle ...添加文章
func AddArticle(ctx *web.Context) {
	form := &models.ArticleContents{}
	if err := ctx.BindJSON(form); err != nil {
		ctx.Response(e.BAD_REQUEST)
		return
	}
	validate := validator.New()
	if err := validate.Struct(form); err != nil {
		ctx.Response(e.BAD_REQUEST)
		return
	}
	srv := service.NewArticleService(ctx.RequestContext)
	code := srv.InsertArticle(form)
	ctx.Response(code)
}

//RemoveArticle ...删除文章
func RemoveArticle(ctx *web.Context)  {
	articleId,_ := strconv.Atoi(ctx.Param("id"))
	srv := service.NewArticleService(ctx.RequestContext)
	code := srv.DeleteByID(articleId)
	ctx.Response(code)
}

//ModifyArticle ...修改文章内容
func ModifyArticle(ctx *web.Context)  {
	form := &models.ArticleContents{}
	articleId,_ := strconv.Atoi(ctx.Param("id"))
	if err := ctx.BindJSON(form) ; err!= nil{
		ctx.Response(e.BAD_REQUEST)
		return
	}
	validate := validator.New()
	if err := validate.Struct(form) ; err != nil{
		ctx.Response(e.BAD_REQUEST)
		return
	}
	srv := service.NewArticleService(ctx.RequestContext)
	code := srv.UpdateArticle(articleId, form)
	ctx.Response(code)
}

//FindByArticle ...获取文章详情
func FindByArticle(ctx *web.Context) {
	articleId, _ := strconv.Atoi(ctx.Param("id"))
	srv := service.NewArticleService(ctx.RequestContext)
	data := srv.ArticleByID(articleId)
	ctx.ResponseData(e.SUCCESS, data)
}

//FindArticleList ...获取文章列表
func FindArticleList(ctx *web.Context){
	pager := &models.PagerArgs{}
	_ = pager.Bind(ctx)
	srv := service.NewArticleService(ctx.RequestContext)
	data := srv.ArticleList(pager)
	ctx.ResponseData(e.SUCCESS, data)
}


