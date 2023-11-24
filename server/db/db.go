package db

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/codex/config"
)

var DB *gorm.DB

func InitDB() {
    var err error
    dbConfig := config.GetDatabaseConfig()
    DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
        dbConfig.Username,
        dbConfig.Password,
        dbConfig.Host,
        dbConfig.Port,
        dbConfig.DBName))
    if err != nil {
        panic(err)
    }
}