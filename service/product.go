package service

import (
	"gomall-center/models"
	"gomall-center/pkg/e"
	"gomall-center/pkg/web"
)

const (
	// 列表查询返回字段---商品
	selectGoodsListFields = "id,  product_name,	pic"
)
const (
	// 列表查询返回字段---商品基本信息
	selectProductsBasicInfoList = "id,  product_name,	producer"
)

// ProductService 企业业务
type ProductService struct {
	Service
}

//NewProductService ...
func NewProductService(ctx *web.RequestContext) *ProductService {
	return &ProductService{
		Service: InitService(ctx),
	}
}

//ProductBasicByStatus ...商品基本信息
func (c *ProductService) ProductBasicByStatus(args *models.PagerArgs, status int) *models.PagerData {
	var result []*models.ProductBasicInfo
	var count int
	db := c.DB.Select(selectProductsBasicInfoList)
	if args.KeyWord != "" {
		db = db.Where("product_name like ? AND publish_status=? AND deleted_at is null", "%"+args.KeyWord+"%", status)
	} else {
		db = db.Where("publish_status=? AND deleted_at is null", status)
	}
	db.Table("t_product").Count(&count).Offset((args.PageNum - 1) * args.PageSize).Limit(args.PageSize).Find(&result)
	pagerdata := &models.PagerData{
		Count: count,
		Data:  result,
	}
	return pagerdata
}

//ProductDetails ...商品要审核的信息
func (c *ProductService) ProductDetails(id int) *models.ProductAudit {
	result := new(models.ProductAudit)
	db := c.DB.Select("publish_status,description,pic,id")
	db.Table("t_product").Where("id=?", id).First(&result)
	return result
}

//UpdateProductStatus ...修改商品发布状态
func (c *ProductService) UpdateProductStatus(id int, product *models.ProductPutStatic) int {
	transaction := c.DB.Begin()
	err := c.DB.Model(&models.Product{}).Where("id=?", id).Update("publish_status", product.PublishStatus).Error
	if err != nil {
		transaction.Rollback()
		return e.ERROR
	}
	transaction.Commit()
	return e.SUCCESS
}

//PropValue ...根据属性ID获取属性值
func (c *ProductService) PropValue(id int, args *models.PagerArgs) *models.PagerData {
	var result []*models.ProductPropValue
	var count int
	db := c.DB
	db.Table("t_product_prop_value").Where("prop_id=?", id).Count(&count).Offset((args.PageNum - 1) * args.PageSize).Limit(args.PageSize).Find(&result)
	pagerdate := &models.PagerData{
		Count: count,
		Data:  result,
	}
	return pagerdate
}

//InsertPropValue ...增加属性值根据属性ID
func (c *ProductService) InsertPropValue(value *models.ProductPropValue) int {
	transaction := c.DB.Begin()
	err := c.DB.Create(value).Error
	if err != nil {
		transaction.Rollback()
		return e.ERROR
	}
	transaction.Commit()
	return e.SUCCESS
}

//DeletePropValue ...删除属性的值
func (c *ProductService) DeletePropValue(id int) int {
	transaction := c.DB.Begin()
	if err := c.DB.Where("id = ?", id).Delete(&models.ProductPropValue{}).Error; err != nil {
		transaction.Rollback()
		return e.ERROR
	}
	transaction.Commit()
	return e.SUCCESS
}

//InsertProp ...添加属性
func (c *ProductService) InsertProp(categoryID int, parentID int, prop *models.ProductProp) int {
	transaction := c.DB.Begin()
	lastProp := &models.ProductProp{}
	db := c.DB
	if err := db.Create(&prop).Error; err != nil {
		transaction.Rollback()
		return e.ERROR
	}
	db.Last(lastProp)

	addCategoryProp := &models.CategoryProp{
		ID:         0,
		CategoryID: categoryID,
		PropID:     lastProp.ID,
	}
	if err := db.Table("t_category_prop").Create(addCategoryProp).Error; err != nil {
		transaction.Rollback()
		return e.ERROR
	}
	transaction.Commit()
	return e.SUCCESS
}

//UpdatePropValue ...修改商品属性的值
func (c *ProductService) UpdatePropValue(prop *models.ProductPropValue, id int) int {
	transaction := c.DB.Begin()
	err := c.DB.Model(&models.ProductPropValue{}).Where("id=?", id).Update("prop_value", prop.PropValue).Error
	if err != nil {
		transaction.Rollback()
		return e.ERROR
	}
	transaction.Commit()
	return e.SUCCESS
}

//UpdateProp ...更改prop
func (c *ProductService) UpdateProp(prop *models.ProductProp, id int) int {
	transaction := c.DB.Begin()
	prop.ID = id
	err := c.DB.Model(&prop).Where("id=?", id).Update(map[string]interface{}{"prop_name": prop.PropName, "prop_type": prop.PropType, "input_type": prop.InputType}).Error
	if err != nil {
		transaction.Rollback()
		return e.ERROR
	}
	transaction.Commit()
	return e.SUCCESS
}

//DeleteProp ...删除一个属性根据属性ID
func (c *ProductService) DeleteProp(id int, categoryID int, parentID int) int {
	transaction := c.DB.Begin()
	db := c.DB
	var number int
	db.Table("t_category_prop").Where("prop_id=?", id).Count(&number)
	if number == 1 {
		if err := db.Where("prop_id=?", id).Delete(&models.ProductPropValue{}).Error; err != nil {
			transaction.Rollback()
			return e.ERROR
		}
		if err := db.Where("id=?", id).Delete(&models.ProductProp{}).Error; err != nil {
			transaction.Rollback()
			return e.ERROR
		}
	}
	if err := db.Where("prop_id=? AND category_id=?", id, categoryID).Delete(&models.CategoryProp{}).Error; err != nil {
		transaction.Rollback()
		return e.ERROR
	}
	transaction.Commit()
	return e.SUCCESS
}
