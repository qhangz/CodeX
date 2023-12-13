package service

import (

	"github.com/codex/dao"
	"github.com/codex/model"
)

// publish comment
func PublishComment(newComment model.Comment) error {
	return dao.PublishComment(newComment)
}
