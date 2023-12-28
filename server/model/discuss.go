package model

import (
	"gorm.io/gorm"
)

// discuss model
type Discuss struct {
	gorm.Model
	// UserID   uint   // 外键，用户id
	Author      string `gorm:"type:varchar(20);not null" json:"author"`
	Title       string `gorm:"type:varchar(255);not null" json:"title"`
	Summary     string `gorm:"type:varchar(255);not null" json:"summary"`
	Content     string `gorm:"type:longtext;not null" json:"content"`
	Category    string `gorm:"type:varchar(20);not null" json:"category"`
	Like_Number int64  `gorm:"type:int" json:"like_num"`
	View_Number int64  `gorm:"type:int" json:"view_num"`
	Comment     []Comment
}

// return discuss list by category
type DiscussList struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Author      string `gorm:"type:varchar(20);not null" json:"author"`
	Title       string `gorm:"type:varchar(20);not null" json:"title"`
	Summary     string `gorm:"type:varchar(20);not null" json:"summary"`
	Like_Number int64  `gorm:"type:int" json:"like_num"`
	View_Number int64  `gorm:"type:int" json:"view_num"`
	CreatedAt   string `gorm:"type:varchar(20);not null" json:"created_at"`
}

// return discuss list by username
type MineDiscuss struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Author      string `gorm:"type:varchar(20);not null" json:"author"`
	Title       string `gorm:"type:varchar(20);not null" json:"title"`
	Summary     string `gorm:"type:varchar(20);not null" json:"summary"`
	Category    string `gorm:"type:varchar(20);not null" json:"category"`
	Like_Number int64  `gorm:"type:int" json:"like_num"`
	View_Number int64  `gorm:"type:int" json:"view_num"`
	CreatedAt   string `gorm:"type:varchar(20);not null" json:"created_at"`
}

// return the top discuss list
type TopDiscuss struct {
	ID uint `gorm:"primary_key" json:"id"`
	// Rank  int    `gorm:"type:int" json:"rank"`
	Title string `gorm:"type:varchar(20);not null" json:"title"`
	// View_Number int64  `gorm:"type:int" json:"view_num"`
}
