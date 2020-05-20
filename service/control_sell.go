package service

import (
	"encoding/json"
	"fmt"
	"gomall-center/models"
	"gomall-center/pkg/e"
	"gomall-center/pkg/web"
	"strings"
)

type ControlSellService struct {
	Service
}

func NewControlSellService (content *web.RequestContext) *ControlSellService  {
	c := &ControlSellService{InitService(content)}
	return c
}

//InsertControlSell ...添加促销模板
func (c *ControlSellService) InsertControlSell (form *models.ControlSells) int{
	tx := c.Begin()
	SysConfigID := strings.Replace(strings.Trim(fmt.Sprint(form.SysConfigID), ""), " ", " ", -1)
	sysConfigId, err := json.Marshal(form.AreaOption)
	if err != nil {
		return e.ERROR
	}
	AreaOption := string(sysConfigId)
	OrdinaryScope := strings.Replace(strings.Trim(fmt.Sprint(form.OrdinaryScope), ""), " ", " ", -1)
	controlsell :=&models.ControlSell{
		CompanyID:form.CompanyID,
		CompanyName:form.CompanyName,
		ControlSellName:form.ControlSellName,
		ControlSellType:form.ControlSellType,
		AreaID:AreaOption,
		SysConfigID:SysConfigID,
		OrdinaryScope:OrdinaryScope,
	}
	if err := c.Create(&controlsell).Error ; err!= nil{
		tx.Rollback()
		return e.ERROR
	}
	tx.Commit()
	return e.SUCCESS
}

//DeleteByID ...根据促销模板ID删除促销模板
func (c *ControlSellService) DeleteByID (controlsellId int) int {
	tx := c.Begin()
	if err := c.Where("id = ?",controlsellId).Delete(&models.ControlSell{}).Error ;err != nil{
		tx.Rollback()
		return e.ERROR
	}
	tx.Commit()
	return e.SUCCESS
}

//UpdateControlSell ...根据促销模板ID修改促销模板内容
func (c *ControlSellService) UpdateControlSell (controlsellId int,form *models.ControlSells) int {
	tx := c.Begin()
	SysConfigID := strings.Replace(strings.Trim(fmt.Sprint(form.SysConfigID), ""), " ", " ", -1)
	sysConfigId, err := json.Marshal(form.AreaOption)
	if err != nil {
		return e.ERROR
	}
	AreaOption := string(sysConfigId)
	OrdinaryScope := strings.Replace(strings.Trim(fmt.Sprint(form.OrdinaryScope), ""), " ", " ", -1)
	result := &models.ControlSell{}
	c.First(result)
	result.CompanyID = form.CompanyID
	result.CompanyName = form.CompanyName
	result.ControlSellName = form.ControlSellName
	result.ControlSellType = form.ControlSellType
	result.AreaID = AreaOption
	result.SysConfigID = SysConfigID
	result. OrdinaryScope = OrdinaryScope
	if err := c.Where("id = ?" ,controlsellId).Save(&result).Error ; err!= nil{
		tx.Rollback()
		return e.ERROR
	}
	tx.Commit()
	return e.SUCCESS
}

//ControlSellByID ...根据促销模板ID获取促销模板详情
func (c *ControlSellService) ControlSellByID (controlsellId int) *models.ControlSell{
	results := &models.ControlSell{}
	if err := c.Where("id = ?",controlsellId).Find(&results).Error ;err !=nil{
		return nil
	}
	return results
}

//ControlSellList ...根据企业ID获取促销模板列表
func (c *ControlSellService) ControlSellList(companyId int,args *models.PagerArgs) *models.PagerData{
	var controlsell []*models.ControlSell
	var count int
	if args.KeyWord != "" {
		c.Where("title like ?", "'%"+args.KeyWord+"%'")
	}
	c.Where("company_id = ?",companyId).Offset((args.PageNum - 1) * args.PageSize).Limit(args.PageSize).Find(&controlsell).Count(&count)
	results := &models.PagerData{
		Count: count,
		Data:  controlsell,
	}
	return results
}