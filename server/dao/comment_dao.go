package dao

import (
	"github.com/codex/db"
	"github.com/codex/model"
)

// publish comment
func PublishComment(newComment model.Comment) error {
	return db.DB.Create(&newComment).Error
}

// add comment view number by comment id from request
func AddCommentView(commentID uint) error {
	return db.DB.Model(&model.Comment{}).Where("id = ?", commentID).Update("view_number", db.DB.Raw("view_number + ?", 1)).Error
}

// add comment like number by comment id from request
func AddCommentLike(commentID uint) error {
	return db.DB.Model(&model.Comment{}).Where("id = ?", commentID).Update("like_number", db.DB.Raw("like_number + ?", 1)).Error
}
