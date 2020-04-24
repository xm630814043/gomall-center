package service

import (
	"gomall-center/models"
	"gomall-center/pkg/web"
)

// CompanyService 企业业务
type CompanyService struct {
	Service
}

const (
	// 列表查询返回字段
	selectListFields = "id, company_name, business_license, legal_person_name"
)

//NewCompanyService new companyservice
func NewCompanyService(ctx *web.RequestContext) *CompanyService {
	return &CompanyService{
		Service: InitService(ctx),
	}
}

// GetCompanyList 获取企业列表
func (c *CompanyService) GetCompanyList(args *models.PagerArgs) []*models.Company {
	var result []*models.Company

	db := c.DB.Select(selectListFields)
	if args.KeyWord != "" {
		db = db.Where("company_name like ?", "'%"+args.KeyWord+"%'")
	}
	db.Offset((args.PageNum - 1) * args.PageSize).Limit(args.PageSize).Find(result)
	return result
}
