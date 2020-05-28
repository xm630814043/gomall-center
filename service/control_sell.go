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
//formControlSell  接受到的from表单数组转换字符串
func formControlSell (form *models.ControlSells) ([]string,int) {
	var forms []string
	SysConfigA := strings.Replace(strings.Trim(fmt.Sprint(form.SysConfigID), ""), " ", " ", -1)
	SysConfigB := strings.Replace(SysConfigA, "[", "", -1 )
	SysConfig := strings.Replace(SysConfigB, "]", "", -1 )
	forms = append(forms,SysConfig )
	AreaID, err := json.Marshal(form.AreaOption)
	if err != nil {
		return nil,e.ERROR
	}
	AreaOptionA := string(AreaID)
	AreaOptionB :=strings.Replace(AreaOptionA, "[", "{", -1 )
	AreaOption :=strings.Replace(AreaOptionB, "]", "}", -1 )
	forms = append(forms,AreaOption )
	OrdinaryScopeA := strings.Replace(strings.Trim(fmt.Sprint(form.OrdinaryScope), ""), " ", " ", -1)
	OrdinaryScopeB :=strings.Replace(OrdinaryScopeA, "[", "", -1 )
	OrdinaryScope :=strings.Replace(OrdinaryScopeB, "]", "", -1 )
	forms = append(forms,OrdinaryScope )
	return forms,e.SUCCESS
}
//InsertControlSell ...添加控销模板
func (c *ControlSellService) InsertControlSell (form *models.ControlSells) int{
	forms,_ := formControlSell(form)
	controlSell :=&models.ControlSell{
		CompanyID:form.CompanyID,
		CompanyName:form.CompanyName,
		ControlSellName:form.ControlSellName,
		ControlSellType:form.ControlSellType,
		SysConfigID:forms[0],
		AreaID:forms[1],
		OrdinaryScope:forms[2],
	}
	if err := c.Create(&controlSell).Error ; err!= nil{
		return e.ERROR
	}
	return e.SUCCESS
}

//DeleteByID ...根据控销模板ID删除控销模板
func (c *ControlSellService) DeleteByID (controlSellId int) int {
	if err := c.Where("id = ?",controlSellId).Delete(&models.ControlSell{}).Error ;err != nil{
		return e.ERROR
	}
	return e.SUCCESS
}

//UpdateControlSell ...根据控销模板ID修改控销模板内容
func (c *ControlSellService) UpdateControlSell (controlSellId int,form *models.ControlSells) int {
	forms,_ := formControlSell(form)

	result := &models.ControlSell{}
	c.First(result)
	result.ID = uint(controlSellId)
	result.CompanyID = form.CompanyID
	result.CompanyName = form.CompanyName
	result.ControlSellName = form.ControlSellName
	result.ControlSellType = form.ControlSellType
	result.SysConfigID = forms[0]
	result.AreaID = forms[1]
	result. OrdinaryScope = forms[2]
	if err := c.Table("t_control_sell").Save(&result).Error ; err!= nil{
		return e.ERROR
	}
	return e.SUCCESS
}

//ControlSellByID ...根据控销模板ID获取控销模板详情
func (c *ControlSellService) ControlSellByID (controlSellId int) *models.ControlSell{
	results := &models.ControlSell{}
	if err := c.Where("id = ?",controlSellId).First(&results).Error ;err !=nil{
		return nil
	}
	return results
}

//ControlSellList ...根据企业ID获取控销模板列表
func (c *ControlSellService) ControlSellList(companyId int,args *models.PagerArgs) *models.PagerData{
	var controlSell []*models.ControlSell
	var count int
	db := c.Table("t_control_sell")
	if args.KeyWord != "" {
		db = db.Where("control_sell_name like ?", "%"+args.KeyWord+"%")
	}
	db = db.Where("company_id = ? and deleted_at is null",companyId).Count(&count)
	if err := db.Offset((args.PageNum - 1) * args.PageSize).Limit(args.PageSize).Find(&controlSell).Error ;err !=nil{
		return nil
	}
	results := &models.PagerData{
		Count: count,
		Data:  controlSell,
	}
	return results
}