package dao

import (
	"github.com/codex/db"
	"github.com/codex/model"
)

// publish comment
func PublishComment(newComment model.Comment) error {
	return db.DB.Create(&newComment).Error
}

