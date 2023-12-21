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
	UserProfile  []UserProfile
}

// user profile model
type UserProfile struct {
	gorm.Model
	UserID uint
	// User   User
	// Name          string `gorm:"type:varchar(20);not null" json:"name"`
	// Age           int    `gorm:"type:int;not null" json:"age"`
	// Summary       string `gorm:"type:varchar(255);not null" json:"summary"`
	DiscussNumber int64 `gorm:"type:int" json:"discuss_number"`
	FansNumber    int64 `gorm:"type:int" json:"fans_number"`
	FollowNumber  int64 `gorm:"type:int" json:"follow_number"`
	Tag           []Tag `gorm:"many2many:user_profile_tags;"` // 多对多关系
}

// user tag model
type Tag struct {
	gorm.Model
	UserProfileID uint
	Name          string `gorm:"type:varchar(20);not null" json:"name"`
}

// return user info
type UserInfo struct {
	gorm.Model
	Username     string `gorm:"type:varchar(20);not null" json:"username"`
	Email        string `gorm:"type:varchar(20);not null;unique" json:"email"`
	Age          int    `gorm:"type:int;not null" json:"age"`
	Summary      string `gorm:"type:varchar(255);not null" json:"summary"`
	Avatar_image string `gorm:"type:varchar(255);not null" json:"avatar_image"`
	UserProfile  []UserProfile
}
