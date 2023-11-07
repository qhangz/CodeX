package main

import (
	"errors"
	"fmt"
	// "github.com/gin-gonic/gin"
	// "goExample/middlewares"
	// "net/http"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// user类型绑定user表
type User struct {
	UserId                int    `gorm:"column:id_user"`
	UserName              string `gorm:"column:username"`
	Password              string `gorm:"column:password"`
	Email                 string `gorm:"column:email"`
	PersonalizedSignature string `gorm:"column:personal_signature"`
}

// 定义表名，返回当前struct绑定的mysql表名
func (u User) TableName() string {
	return "user"
}

func main() {
	// fmt.Println("Hello World!")

	// **********gin
	// r := gin.Default()
	// // r.Use(middlewares.Cors())

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello World!",
	// 	})
	// })

	// r.Run(":8080")

	// **********gorm
	//配置MySQL连接参数
	username := "root"       //账号
	password := "password"   //密码
	host := "127.0.0.1"      //数据库地址，可以是Ip或者域名
	port := 3306             //数据库端口
	DBname := "my_community" //数据库名

	//通过前面的数据库参数，拼接MYSQL DSN， 其实就是数据库连接串（数据源名称）
	//MYSQL dsn格式： {username}:{password}@tcp({host}:{port})/{Dbname}?charset=utf8&parseTime=True&loc=Local
	//类似{username}使用花括号包着的名字都是需要替换的参数
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, DBname)
	//连接MYSQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	} else {
		fmt.Println("连接数据库成功")
	}

	//定义一个用户，并初始化数据
	u := User{
		UserId:                0,
		UserName:              "name",
		Password:              "password",
		Email:                 "example@email.com",
		PersonalizedSignature: "signature",
	}

	//	插入一条用户数据
	//下面代码会自动生成SQL语句：INSERT INTO `user` (`id_user`,`username`,`password`,`email`,`personal_signature`) VALUES (2,'name','password','example@email.com','signature')
	// res := db.Create(&u)
	// if err := res.Error; err != nil {
	// 	fmt.Println("插入失败", err)
	// 	return
	// } else {
	// 	fmt.Println("插入成功")
	// }
	// fmt.Printf("影响行数res:%+v", res.RowsAffected) // 输出影响行数

	//	查询并返回第一条数据
	//定义需要保存数据的struct变量
	u = User{}
	// 自动生成sql： SELECT * FROM `user` WHERE username = 'hang'  ORDER BY `user`.`id` LIMIT 1
	result := db.Where("username = ? ", "hang").First(&u)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("找不到记录")
		return
	}
	//打印查询到的数据
	fmt.Println(u.UserName, u.Password)

	//	更新
	//自动生成Sql:  UPDATE `user` SET `username`="HANG"  WHERE id_user = '2'
	db.Model(&User{}).Where("id_user = ?", "2").Update("username", "HANG")

	//	删除
	//自动生成Sql： DELETE FROM `user` WHERE id_user = 3
	db.Where("id_user = ?", 3).Delete(&User{})

}
