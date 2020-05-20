package service

import (
	"gomall-center/models"
	"gomall-center/pkg/e"
	"gomall-center/pkg/web"
)

//定义栏目主题
const (
	SubjectProduct = "product"
	SubjectShop    = "shop"
	SubjectArticle = "article"
)

//SubjectService ...
type SubjectService struct {
	Service
}

//NewSubjectService ...
func NewSubjectService(context *web.RequestContext) *SubjectService {
	c := &SubjectService{InitService(context)}
	return c
}

// ChannelList ...获取频道专题列表
func (c *SubjectService) ChannelList() []*models.Channel {
	var channels []*models.Channel
	var subjects []*models.Subject
	c.Find(&channels)
	c.Find(&subjects)
	for _, c := range channels {
		for _, s := range subjects {
			if int64(c.ID) == s.ChannelID {
				c.Subjects = append(c.Subjects, s)
			}
		}
	}
	return channels
}

//SubjectListByID ...根据频道id,专题id,获取频道专题的列表
func (c *SubjectService) SubjectListByID(subjectID int64, contentType string, limit int) interface{} {
	db := c.Where("id in (select object_id from t_subject_content_relation where subject_id = ? )", subjectID).Limit(limit)
	if contentType == SubjectProduct {
		var products []*models.Product
		db.Find(&products)
		return products
	} else if contentType == SubjectShop {
		var companyID []int64
		var shop []*models.Shop
		var shops []*models.ShopVo
		db.Find(&shop)
		for _, id := range shop {
			companyID = append(companyID, id.CompanyID)
		}
		c.Raw("SELECT s.id,s.company_id,c.logo,s.shop_name FROM t_shop as s,t_company as c WHERE s.company_id = c.id  and s.company_id in(?)", companyID).Scan(&shops)
		return shops
	} else if contentType == SubjectArticle {
		var articles []*models.Article
		db.Find(&articles)
		return articles
	}
	return nil
}

//DeleteByID ...	根据专题id,专题内容id,删除专题对应产品关系
func (c *SubjectService) DeleteByID(objectID int, subjects *models.SubjectContentRelation) int {
	tx := c.Begin()
	if err := c.Where("object_id = ? and subject_id = ?", objectID, subjects.SubjectID).Delete(&models.SubjectContentRelation{}).Error; err != nil{
		tx.Rollback()
		return e.ERROR
	}
	tx.Commit()
	return e.SUCCESS
}

//InsertSubject ...	根据专题id,专题内容id,添加专题对应产品关系
func (c *SubjectService) InsertSubject(subject *models.SubjectContentRelation) int {
	tx := c.Begin()
	var subjectContent = models.SubjectContentRelation{SubjectID: subject.SubjectID, ObjectID: subject.ObjectID}
	if err := c.Table("t_subject_content_relation").Create(&subjectContent).Error; err != nil{
		tx.Rollback()
		return e.ERROR
	}
	tx.Commit()
	return e.SUCCESS
}

//SubjectByID ...根据专题内容id,获取专题内容详情
func (c *SubjectService) SubjectByID(objectID int64, contentType string) interface{} {
	if contentType == SubjectProduct {
		products := &models.Product{}
		c.Where("id = ? ", objectID).Find(&products)
		return products
	} else if contentType == SubjectShop {
		shops := &models.Shop{}
		c.Where("id = ? ", objectID).Find(&shops)
		return shops
	} else if contentType == SubjectArticle {
		articles := &models.Article{}
		c.Where("id = ? ", objectID).Find(&articles)
		return articles
	}
	return nil
}
