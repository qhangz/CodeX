package model

import (
	"gorm.io/gorm"
)

// article model
type Article struct {
	gorm.Model
	Article_ID     uint   `gorm:"primary_key" json:"article_id"`
	Author         string `json:"author"`
	Title          string `json:"title"`
	Summary        string `json:"summary"`
	View_Number    int64  `json:"view_number"`
	Like_Number    int64  `json:"like_number"`
	Comment_Number int64  `json:"comment_number"`
}
