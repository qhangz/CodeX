package service

import (
    "github.com/codex/dao"
    "github.com/codex/model"
)

func GetUserByUsername(username string) (*model.User, error) {
    return dao.GetUserByUsername(username)
}