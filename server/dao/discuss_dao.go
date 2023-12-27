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
	err := db.DB.Where("id = ?", discussID).First(&discuss)
	if err.Error != nil {
		return nil, err.Error
	}

	db.DB.Model(&discuss).Association("Comment").Find(&discuss.Comment)
	return &discuss, nil
}

// get discuss list by category
// 返回id, title, summary, author, category, like_number, view_number, created_at，按照created_at降序排列
func GetDiscussList(category string) ([]model.DiscussList, error) {
	var discussList []model.DiscussList
	err := db.DB.Model(&model.Discuss{}).
		Select("id", "title", "summary", "author", "like_number", "view_number", "created_at").
		Where("category = ?", category).Order("created_at desc").
		Scan(&discussList)
	if err.Error != nil {
		return nil, err.Error
	}
	return discussList, nil
}

// get the top title of discuss list from pre to end number
func GetDiscussTop(pre int, end int) ([]model.TopDiscuss, error) {
	var topDiscussList []model.TopDiscuss
	err := db.DB.Model(&model.Discuss{}).
		Select("id", "title").
		Order("view_number desc").Limit(end - pre + 1).Offset(pre - 1).
		Scan(&topDiscussList)
	if err.Error != nil {
		return nil, err.Error
	}
	return topDiscussList, nil
}

// add discuss view number
func AddDiscussView(discussID uint) error {
	return db.DB.Model(&model.Discuss{}).Where("id = ?", discussID).Update("view_number", db.DB.Raw("view_number + ?", 1)).Error
}

// add discuss like number
func AddDiscussLike(discussID uint) error {
	return db.DB.Model(&model.Discuss{}).Where("id = ?", discussID).Update("like_number", db.DB.Raw("like_number + ?", 1)).Error
}