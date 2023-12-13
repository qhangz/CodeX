package service

import (

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
