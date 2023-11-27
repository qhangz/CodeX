package model

import (
	"time"
)

type User struct {
	ID           uint      `gorm:"primary_key" json:"id_user"`
	Username     string    `gorm:"unique_index" json:"username"`
	Password     string    `json:"password"`
	Email        string    `json:"email"`
	Age          int       `json:"age"`
	Summary      string    `json:"summary"`
	CreatedAt    time.Time `json:"created_at"`
	Avatar_imaeg string    `json:"avatar_image"`
}
