package dao

import (
	"goblog/config"
	"goblog/models"
)

func CreateUser(user *models.User) error {
	result := config.MysqlDB.Create(user)
	return result.Error
}

func FindUserByNameAndPwd(username, password string) (*models.User, error) {
	var user models.User
	result := config.MysqlDB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, nil // 用户不存在
		}
		return nil, result.Error
	}
	if user.CheckPassword(password) == false {
		return nil, nil // 密码错误
	}

	return &user, nil
}
func FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := config.MysqlDB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, nil // 用户不存在
		}
		return nil, result.Error
	}

	return &user, nil
}
