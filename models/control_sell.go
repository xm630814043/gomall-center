package models

import "github.com/jinzhu/gorm"


//ControlSell struct 控销模板
type ControlSell struct {
	gorm.Model
	CompanyID 			int    `json:"company_id"`
	CompanyName 		string `json:"company_name"`
	ControlSellName 	string `json:"control_sell_name"`
	ControlSellType 	int    `json:"control_sell_type"`
	SysConfigID 		string `json:"sys_config_id"`
	AreaID				string `json:"area_id"`
	OrdinaryScope		string `json:"ordinary_scope"`
}

type AreaOption struct {
	ID       int   	`json:"id"`
	AreaName string `json:"area_name"`
}

//ControlSell struct 控销模板表单
type ControlSells struct {
	CompanyID 			int  			`json:"company_id" validate:"required"`
	CompanyName 		string 			`json:"company_name" validate:"required"`
	ControlSellName 	string 			`json:"control_sell_name" validate:"required"`
	ControlSellType 	int 			`json:"control_sell_type" validate:"required"`
	AreaOption  		[]AreaOption	`json:"area_option"`
	SysConfigID 		[]int 			`json:"sys_config_id"`
	OrdinaryScope		[]string 		`json:"ordinary_scope"`
}



