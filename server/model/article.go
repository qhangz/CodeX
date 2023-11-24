package model

import (
	"time"
)

type Article struct {
	Article_ID     uint      `gorm:"primary_key" json:"article_id"`
	Author         string    `json:"author"`
	Title          string    `json:"title"`
	Summary        string    `json:"summary"`
	Created_At     time.Time `json:"created_at"`
	View_Number    int64     `json:"view_number"`
	Like_Number    int64     `json:"like_number"`
	Comment_Number int64     `json:"comment_number"`
}

// type Article struct {
// 	Article_ID     primitive.ObjectID `bson:"_id"`
// 	Cover_URL      string             `json:"cover_url"`
// 	Title          string             `json:"title"`
// 	Author         string             `json:"author"`
// 	Summary        string             `json:"summary"`
// 	Content        string             `json:"content"`
// 	Likes          int64              `json:"likes"`
// 	Comment_Number int64              `json:"comment_number"`
// 	Comments       []Comment          `json:"comments" bson:"comments"`
// 	Created_At     time.Time          `json:"created_at"`
// 	Updated_At     time.Time          `json:"updated_at"`
// }
