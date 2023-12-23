package dao

import (
	"github.com/codex/db"
	"github.com/codex/model"
)

// get all user infomation by username
func GetUserByUsername(username string) (*model.User, error) {
	// var user model.User
	user := model.User{}
	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	db.DB.Model(&user).Association("UserProfile").Find(&user.UserProfile)
	return &user, nil
}

// get userinfo by username
// func GetUserInfoByUsername(username string) (*model.UserInfo, error) {
// 	var userInfo model.UserInfo
// 	if err := db.DB.Model(&model.User{}).
// 		Select("id", "created_at", "username", "email", "age", "summary", "avatar_image").
// 		Where("username = ?", username).
// 		Scan(&userInfo).Error; err != nil {
// 		return nil, err
// 	}
// 	// db.DB.Model(&userInfo).Association("UserProfile").Find(&userInfo.UserProfile)
// 	return &userInfo, nil
// }

// get user by email
func GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// get user by id
func GetUserById(id int) (*model.User, error) {
	var user model.User
	if err := db.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// is email exist
func IsEmailExist(email string) bool {
	var user model.User
	db.DB.Where("email = ?", email).First(&user)
	return user.ID != 0
}

// is username exist
func IsUsernameExist(username string) bool {
	var user model.User
	db.DB.Where("username = ?", username).First(&user)
	return user.ID != 0
}

// user register(create new user)
func Register(user model.User) error {
	return db.DB.Create(&user).Error
}

// user list
func GetUserList() ([]model.User, error) {
	var users []model.User
	// 返回所有用户的信息
	if err := db.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// update username
func UpdateUsername(user *model.User, newUsername string) error {
	return db.DB.Model(&user).Update("username", newUsername).Error
}

// func UpdateUsername(user model.User, newUsername string) error {
// 	// 将user中username字段更新为newUsername
// 	return db.DB.Model(&user).Where("username = ?", user.Username).Update("username", newUsername).Error
// }

// update password
func UpdatePassword(user *model.User, newPassword string) error {
	return db.DB.Model(&user).Update("password", newPassword).Error
}

// update email
func UpdateEmail(user *model.User, newEmail string) error {
	return db.DB.Model(&user).Update("email", newEmail).Error
}

// update age
func UpdateAge(user *model.User, newAge int) error {
	return db.DB.Model(&user).Update("age", newAge).Error
}

// update summary
func UpdateSummary(user *model.User, newSummary string) error {
	return db.DB.Model(&user).Update("summary", newSummary).Error
}

// delete user
func DeleteUser(user *model.User) error {
	return db.DB.Delete(&user).Error
}

// get user's avatar image accoring to username
func GetUserAvatar(username string) (string, error) {
	var avatar string
	err := db.DB.Model(&model.User{}).Select("avatar_image").Where("username=?", username).Scan(&avatar).Error
	if err != nil {
		return "", err
	}
	return avatar, nil
}
