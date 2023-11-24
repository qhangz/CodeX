package dao

import (
    "github.com/codex/db"
    "github.com/codex/model"
)

func GetUserByUsername(username string) (*model.User, error) {
    var user model.User
    if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}