package service

import (
	// "time"

	"github.com/codex/dao"
	"github.com/codex/model"
	// "golang.org/x/text/number"
)

// Convert the time format of CreatedAt to 2021-01-01 00:00:00
func ConvertCreatedAt(discussList []model.DiscussList) []model.DiscussList {
	for i := range discussList {
		discussList[i].CreatedAt = discussList[i].CreatedAt[:10] + " " + discussList[i].CreatedAt[11:19]
	}
	return discussList
}

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
	// var discussList []model.DiscussList
	discussList, err := dao.GetDiscussList(category)
	if err != nil {
		return nil, err
	}

	// convert time format
	discussList = ConvertCreatedAt(discussList)

	return discussList, nil
}

// get the top title of discuss list
func GetDiscussTop(pre int, end int) ([]model.TopDiscuss, error) {
	// var topDiscussList []model.TopDiscuss
	topDiscussList, err := dao.GetDiscussTop(pre, end)
	if err != nil {
		return nil, err
	}
	return topDiscussList, nil
}
