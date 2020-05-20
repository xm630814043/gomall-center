package service

import (
	"gomall-center/models"
	"gomall-center/pkg/e"
	"gomall-center/pkg/web"
)

type ArticleService struct {
	Service
}

func NewArticleService(content *web.RequestContext) *ArticleService {
	c := &ArticleService{InitService(content)}
	return c
}

//InsertArticle ...添加文章
func (c *ArticleService) InsertArticle(form *models.ArticleContents) int {
	tx := c.Begin()
	var article = models.Article{Title: form.Title,Contents: form.Contents}
	if err := c.Create(&article).Error ; err !=nil{
		tx.Rollback()
		return e.ERROR
	}
	tx.Commit()
	return e.SUCCESS
}

//DeleteByID ...根据ID删除文章
func (c *ArticleService) DeleteByID (articleId int) int {
	tx := c.Begin()
	if err := c.Where("id = ?",articleId).Delete(&models.Article{}).Error ; err != nil{
		tx.Rollback()
		return e.ERROR
	}
	tx.Commit()
	return e.SUCCESS
}

//UpdateArticle ...根据ID修改文章内容
func (c *ArticleService) UpdateArticle (articleId int,form *models.ArticleContents) int {
	tx :=c.Begin()
	if err := c.Table("t_article").Where("id = ? ",articleId).Update(form).Error ; err != nil{
		tx.Rollback()
		return e.ERROR
	}
	tx.Commit()
	return e.SUCCESS
}

//ArticleByID ...根据ID获取文章详情
func (c *ArticleService) ArticleByID(articleId int) *models.Article {
	results := &models.Article{}
	if err := c.Where("id = ?", articleId).First(&results).Error ; err!= nil{
		return nil
	}
	return results
}

//ArticleList ...获取文章列表
func (c *ArticleService) ArticleList (args *models.PagerArgs) *models.PagerData  {
	var articles  []*models.Article
	var count int
	if args.KeyWord != "" {
		c.Where("title like ?", "'%"+args.KeyWord+"%'")
	}
	c.Offset((args.PageNum - 1) * args.PageSize).Limit(args.PageSize).Find(&articles).Count(&count)
	results := &models.PagerData{
		Count: count,
		Data:  articles,
	}
	return results
}

