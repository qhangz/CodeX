package service

import (
	// "time"

	"github.com/codex/dao"
	"github.com/codex/model"
)

// publish discuss
func PublishDiscuss(newDiscuss model.Discuss) error {
	return dao.PublishDiscuss(newDiscuss)
}

// get discuss info by discuss id from request
func GetDiscussInfo(discussID uint) (*model.Discuss, error) {
	return dao.GetDiscussInfo(discussID)
}

// get discuss list by category
func GetDiscussList(category string) ([]model.DiscussList, error) {
	var discussList []model.DiscussList
	discussList, err := dao.GetDiscussList(category)
	if err != nil {
		return nil, err
	}

	// 2023-12-17T16:08:43.42+08:00去掉后面的.42+08:00和T
	for i := range discussList {
		discussList[i].CreatedAt = discussList[i].CreatedAt[:10] + " " + discussList[i].CreatedAt[11:19]
	}

	return discussList, nil
}
