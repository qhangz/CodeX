package service

import (
	"github.com/codex/dao"
	"github.com/codex/model"
)

func GetArticleHot() ([]model.Article, error) {
	return dao.GetArticleHot()
}