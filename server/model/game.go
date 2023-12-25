package model

import (
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Route        string `gorm:"type:varchar(255);not null" json:"route"`
	Banner_Image string `gorm:"type:varchar(255);not null" json:"banner_image"`
	Game_Name    string `gorm:"type:varchar(255);not null" json:"game_name"`
}
