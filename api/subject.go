package api

import (
	"gomall-center/models"
	"gomall-center/pkg/e"
	"gomall-center/pkg/web"
	"gomall-center/service"
	"gopkg.in/go-playground/validator.v9"
	"strconv"
)

// FindChannelList ...获取频道专题
func FindChannelList(ctx *web.Context) {
	srv := service.NewSubjectService(ctx.RequestContext)
	data := srv.ChannelList()
	ctx.ResponseData(e.SUCCESS, data)
}

// FindBySubjectList ...获取频道专题内容列表
func FindBySubjectList(ctx *web.Context) {
	subjectID, _ := strconv.ParseInt(ctx.Query("id"), 10, 64)
	contentType := ctx.Query("contentType")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	if limit == 0 {
		limit = 10
	}
	srv := service.NewSubjectService(ctx.RequestContext)
	data := srv.SubjectListByID(subjectID, contentType, limit)
	ctx.ResponseData(e.SUCCESS, data)
}

// RemoveSubject ...删除专题对应产品关系
func RemoveSubject(ctx *web.Context) {
	objectID, _ := strconv.Atoi(ctx.Param("id"))
	subjectContent := &models.SubjectContentRelation{}
	_ = ctx.ShouldBindQuery(subjectContent)
	srv := service.NewSubjectService(ctx.RequestContext)
	code := srv.DeleteByID(objectID, subjectContent)
	ctx.Response(code)
}

// AddSubject ...添加专题对应产品关系
func AddSubject(ctx *web.Context) {
	form := &models.SubjectContentRelation{}
	if err := ctx.BindJSON(form); err != nil {
		ctx.Response(e.BAD_REQUEST)
		return
	}
	validate := validator.New()
	if err := validate.Struct(form); err != nil {
		ctx.Response(e.BAD_REQUEST)
		return
	}
	srv := service.NewSubjectService(ctx.RequestContext)
	code := srv.InsertSubject(form)
	ctx.Response(code)
}

// FindBySubject  ...获取频道专题内容详情
func FindBySubject(ctx *web.Context) {
	objectID, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	contentType := ctx.Query("contentType")
	srv := service.NewSubjectService(ctx.RequestContext)
	data := srv.SubjectByID(objectID, contentType)
	ctx.ResponseData(e.SUCCESS, data)
}
