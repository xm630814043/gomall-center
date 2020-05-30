package api

import (
	"github.com/go-playground/validator/v10"
	"gomall-center/models"
	"gomall-center/pkg/e"
	"gomall-center/pkg/web"
	"gomall-center/service"
	"strconv"
)

//Discounts ...折扣
func FIndDiscounts  (ctx *web.Context){
	promotionActivityId,_:= strconv.Atoi(ctx.Query("id"))
	form := &models.PromotionPatternFrom{}
	if err := ctx.BindJSON(form);err != nil{
		ctx.Response(e.BAD_REQUEST)
		return
	}
	validate := validator.New()
	if err :=validate.Struct(form);err != nil{
		ctx.Response(e.BAD_REQUEST)
		return
	}
	srv := service.NewPromotionPattern(ctx.RequestContext)
	code := srv.InsertPromotionPattern(promotionActivityId,form)
	ctx.Response(code)
}

//StandBy ...立减
func StandBy  (ctx web.Context){

}

//FullReduction ...满额减
func QuotaReduction  (ctx web.Context){

}

//FullReduction ...满折减
func FullReduction  (ctx web.Context){

}

//FullVolume ...满折减
func FullVolume  (ctx web.Context){

}

//FullVolume ...满额赠
func GiveFull  (ctx web.Context){

}

//FullVolume ...满折赠
func GiveFullFold  (ctx web.Context){

}