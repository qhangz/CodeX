package model

import (
	"gorm.io/gorm"
)

// commment model
type Comment struct {
	gorm.Model
	DiscussID   uint   // 外键，discussid
	Author      string `gorm:"type:varchar(20);not null" json:"author"`
	Content     string `gorm:"type:varchar(255);not null" json:"content"`
	Like_Number int64  `gorm:"type:int" json:"like_num"`
	View_Number int64  `gorm:"type:int" json:"view_num"`
}

