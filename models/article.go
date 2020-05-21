package models

import "github.com/jinzhu/gorm"

//Article ...轮播图
type Article struct {
	gorm.Model
	Title     string `json:"title"`
	ImageURL  string `json:"image_url"`
	Contents  string `json:"contents"`
	SubjectID int64  `json:"subject_id"`
	Tags 	  string `json:"tags"`
}
