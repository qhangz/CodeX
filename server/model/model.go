package model

import (
	"gorm.io/gorm"
)

// user model
type User struct {
	gorm.Model
	Username     string `gorm:"type:varchar(20);not null" json:"username"`
	Password     string `gorm:"size:255;not null" json:"password"`
	Email        string `gorm:"type:varchar(20);not null;unique" json:"email"`
	Age          int    `gorm:"type:int;not null" json:"age"`
	Summary      string `gorm:"type:varchar(255);not null" json:"summary"`
	Avatar_image string `gorm:"type:varchar(255);not null" json:"avatar_image"`
	// UserProfile   []UserProfile
}

// user profile model
type UserProfile struct {
	gorm.Model
	UserID uint
	User   User
	Tag    []Tag `gorm:"many2many:user_profile_tags;"` // 多对多关系
}

// user tag model
type Tag struct {
	gorm.Model
	UserProfileID uint
	Name          string `gorm:"type:varchar(20);not null" json:"name"`
}

// discuss model
type Discuss struct {
	gorm.Model
	// UserID   uint   // 外键，用户id
	Author      string `gorm:"type:varchar(20);not null" json:"author"`
	Title       string `gorm:"type:varchar(20);not null" json:"title"`
	Summary     string `gorm:"type:varchar(20);not null" json:"summaty"`
	Content     string `gorm:"type:varchar(255);not null" json:"content"`
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

// commment model
type Comment struct {
	gorm.Model
	DiscussID   uint   // 外键，discussid
	Author      string `gorm:"type:varchar(20);not null" json:"author"`
	Content     string `gorm:"type:varchar(255);not null" json:"content"`
	Like_Number int64  `gorm:"type:int" json:"like_num"`
	View_Number int64  `gorm:"type:int" json:"view_num"`
}

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
