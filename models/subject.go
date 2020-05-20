package models

import "github.com/jinzhu/gorm"

//Channel ...频道
type Channel struct {
	gorm.Model
	ChannelName string     `json:"channel_name"` //  频道名称
	Icon        string     `json:"icon"`         // 频道图标
	Sort        int        `json:"sort"`         //  排序
	ShowStatus  int8       `json:"show_status"`  // 显示状态
	Subjects    []*Subject `json:"subjects"`
}

//Subject ...专题
type Subject struct {
	gorm.Model
	Title       string `json:"title"`        // 栏目标题
	Pic         string `json:"pic"`          // 栏目图片
	Description string `json:"description"`  // 栏目说明
	ShowStatus  int8   `json:"show_status"`  //  栏目显示状态
	ChannelID   int64  `json:"channel_id"`   //  所属频道id
	SubjectType int8   `json:"subject_type"` //  栏目类型   //  1 店铺  2 产品  3 商城动态
}

//SubjectContentRelation ...专题对应产品关系表
type SubjectContentRelation struct {
	ID        int64  `json:"id"`
	SubjectID int64  `json:"subject_id" validate:"required"`
	ObjectID  int64  `json:"object_id" validate:"required"`
	TypeName  string `json:"type_name"`
}
