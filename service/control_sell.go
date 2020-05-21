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

//InsertControlSell ...添加控销模板
func (c *ControlSellService) InsertControlSell (form *models.ControlSells) int{
	SysConfiga := strings.Replace(strings.Trim(fmt.Sprint(form.SysConfigID), ""), " ", " ", -1)
	SysConfigb := strings.Replace(SysConfiga, "[", "", -1 )
	SysConfig := strings.Replace(SysConfigb, "]", "", -1 )
	AreaID, err := json.Marshal(form.AreaOption)
	if err != nil {
		return e.ERROR
	}
	AreaOptiona := string(AreaID)
	AreaOptionb :=strings.Replace(AreaOptiona, "[", "{", -1 )
	AreaOption :=strings.Replace(AreaOptionb, "]", "}", -1 )
	OrdinaryScopea := strings.Replace(strings.Trim(fmt.Sprint(form.OrdinaryScope), ""), " ", " ", -1)
	OrdinaryScopeab :=strings.Replace(OrdinaryScopea, "[", "", -1 )
	OrdinaryScope :=strings.Replace(OrdinaryScopeab, "]", "", -1 )
	controlsell :=&models.ControlSell{
		CompanyID:form.CompanyID,
		CompanyName:form.CompanyName,
		ControlSellName:form.ControlSellName,
		ControlSellType:form.ControlSellType,
		AreaID:AreaOption,
		SysConfigID:SysConfig,
		OrdinaryScope:OrdinaryScope,
	}
	if err := c.Create(&controlsell).Error ; err!= nil{
		return e.ERROR
	}
	return e.SUCCESS
}

//DeleteByID ...根据控销模板ID删除控销模板
func (c *ControlSellService) DeleteByID (controlsellId int) int {
	if err := c.Where("id = ?",controlsellId).Delete(&models.ControlSell{}).Error ;err != nil{
		return e.ERROR
	}
	return e.SUCCESS
}

//UpdateControlSell ...根据控销模板ID修改控销模板内容
func (c *ControlSellService) UpdateControlSell (controlsellId int,form *models.ControlSells) int {
	SysConfiga := strings.Replace(strings.Trim(fmt.Sprint(form.SysConfigID), ""), " ", " ", -1)
	SysConfigb := strings.Replace(SysConfiga, "[", "", -1 )
	SysConfig := strings.Replace(SysConfigb, "]", "", -1 )
	AreaID, err := json.Marshal(form.AreaOption)
	if err != nil {
		return e.ERROR
	}
	AreaOptiona := string(AreaID)
	AreaOptionb :=strings.Replace(AreaOptiona, "[", "{", -1 )
	AreaOption :=strings.Replace(AreaOptionb, "]", "}", -1 )
	OrdinaryScopea := strings.Replace(strings.Trim(fmt.Sprint(form.OrdinaryScope), ""), " ", " ", -1)
	OrdinaryScopeab :=strings.Replace(OrdinaryScopea, "[", "", -1 )
	OrdinaryScope :=strings.Replace(OrdinaryScopeab, "]", "", -1 )
	result := &models.ControlSell{}
	c.First(result)
	result.ID = uint(controlsellId)
	result.CompanyID = form.CompanyID
	result.CompanyName = form.CompanyName
	result.ControlSellName = form.ControlSellName
	result.ControlSellType = form.ControlSellType
	result.AreaID = AreaOption
	result.SysConfigID = SysConfig
	result. OrdinaryScope = OrdinaryScope
	if err := c.Table("t_control_sell").Save(&result).Error ; err!= nil{
		return e.ERROR
	}
	return e.SUCCESS
}

//ControlSellByID ...根据控销模板ID获取控销模板详情
func (c *ControlSellService) ControlSellByID (controlsellId int) *models.ControlSell{
	results := &models.ControlSell{}
	if err := c.Where("id = ?",controlsellId).First(&results).Error ;err !=nil{
		return nil
	}
	return results
}

//ControlSellList ...根据企业ID获取控销模板列表
func (c *ControlSellService) ControlSellList(companyId int,args *models.PagerArgs) *models.PagerData{
	var controlsell []*models.ControlSell
	var count int
	if args.KeyWord != "" {
		c.Where("title like ?", "'%"+args.KeyWord+"%'")
	}
	if err := c.Where("company_id = ?",companyId).Offset((args.PageNum - 1) * args.PageSize).Limit(args.PageSize).Find(&controlsell).Count(&count).Error ;err !=nil{
		return nil
	}
	results := &models.PagerData{
		Count: count,
		Data:  controlsell,
	}
	return results
}