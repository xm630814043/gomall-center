package models

import (
	"github.com/jinzhu/gorm"
	"gomall-center/pkg/enum"
)

// Account struct
type Account struct {
	gorm.Model
	Username    string        `json:"username"`
	Password    string        `json:"password"`
	Email       string        `json:"email"`
	Mobile      string        `json:"mobile"`
	UserType    enum.UserType `json:"user_type"`
	CompanyID   int64         `json:"company_id"`
	LastLoginIP string        `json:"last_login_ip"`
	Avatar      string        `json:"avatar"`
	PayPassword string        `json:"pay_password"`
	Sex         string        `json:"sex"`
}
