package dao 
import (
	"github.com/codex/db"
	"github.com/codex/model"
)

func GetArticleHot() ([]model.Article, error) {
	var article []model.Article
	if err := db.DB.Where("is_hot = ?", true).Find(&article).Error; err != nil {
		return nil, err
	}
	return article, nil
}