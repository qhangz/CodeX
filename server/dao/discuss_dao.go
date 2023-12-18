package dao

import (
	"github.com/codex/db"
	"github.com/codex/model"
)

// publish discuss
func PublishDiscuss(newDiscuss model.Discuss) error {
	return db.DB.Create(&newDiscuss).Error
}

// get discuss info by discuss id from request
func GetDiscussInfo(discussID uint) (*model.Discuss, error) {
	discuss := model.Discuss{}
	db.DB.Where("id = ?", discussID).First(&discuss)

	db.DB.Model(&discuss).Association("Comment").Find(&discuss.Comment)

	return &discuss, nil
}

// get discuss list by category
// 返回id, title, summary, author, category, like_number, view_number, created_at，按照created_at降序排列
func GetDiscussList(category string) ([]model.DiscussList, error) {
	var discussList []model.DiscussList
	db.DB.Model(&model.Discuss{}).Select("id","title", "summary", "author", "like_number", "view_number", "created_at").Where("category = ?", category).Order("created_at desc").Scan(&discussList)
	return discussList, nil
}

