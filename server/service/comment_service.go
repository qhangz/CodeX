package service

import (
	"github.com/codex/dao"
	"github.com/codex/model"
)

// publish comment
func PublishComment(newComment model.Comment) error {
	comment := newComment
	// get the avatar according to the author name
	avatar, err := dao.GetUserAvatar(comment.Author)
	if err != nil {
		return err
	}
	comment.Avatar = avatar
	return dao.PublishComment(comment)
}
