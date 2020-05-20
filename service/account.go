package service

import (
	"gomall-center/models"
	"gomall-center/pkg/web"
)

// AccountService ...
type AccountService struct {
	Service
}

//NewAccountService new AccountService
func NewAccountService(ctx *web.RequestContext) *AccountService {
	return &AccountService{
		Service: InitService(ctx),
	}
}

//AccountList ...获取用户列表
func (c *AccountService) AccountList(Radio int, args *models.PagerArgs) *models.PagerData {
	var accounts []*models.Account
	var count int
	db := c.Table("t_account").Where("company_id = ? ", Radio)
	if args.KeyWord != "" {
		db = db.Where("username like ?", "'%"+args.KeyWord+"%'")
	}
	db.Offset((args.PageNum - 1) * args.PageSize).Limit(args.PageSize).Find(&accounts).Count(&count)
	results := &models.PagerData{
		Count: count,
		Data:  accounts,
	}
	return results
}

//AccountByID ...根据用户ID获取用户详情
func (c *AccountService) AccountByID(id int) *models.Account {
	results := &models.Account{}
	c.Table("t_account").Where("id = ?", id).Find(&results)
	return results
}
