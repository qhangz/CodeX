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

// add comment view number by comment id from request
func AddCommentView(commentID uint) error {
	return dao.AddCommentView(commentID)
}

// add comment like number by comment id from request
func AddCommentLike(commentID uint) error {
	return dao.AddCommentLike(commentID)
}