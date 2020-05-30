package service

import (
	"gomall-center/models"
	"gomall-center/pkg/e"
	"gomall-center/pkg/web"
	"strings"
)
//PromotionActivity ...
type PromotionPatternService struct {
	Service
}

//NewPromotionPattern ...
func NewPromotionPattern(content *web.RequestContext) *PromotionPatternService  {
	c := &PromotionActivity{InitService(content)}
	return (*PromotionPatternService)(c)
}

func (c *PromotionPatternService) InsertPromotionPattern (promotionActivityId int ,form *models.PromotionPatternFrom) int {
	promotionActivity := &models.ProductAttrVo{}
	promotionRelation := &models.ProductRelation{}
	var productA []int
	var productB []int
	c.Table("t_promotion_activity").Where("id = ?" ,promotionActivityId).Find(&promotionActivity)
	c.Table("t_product_relation").Where("promotion_activity_id = ?" ,promotionActivity).Find(&promotionRelation)
	for _,f := range form.ID{
		start :=strings.Index(promotionRelation.ProductId, string(f))
		if start != -1{
			productA =append(productA,f)
		}else{
			productB =append(productB,f)
		}
	}
	return e.SUCCESS
}
