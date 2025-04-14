package service

import (
	"errors"
	"goblog/dao"
	"goblog/models"
)

func RegisterUser(username, password, gender, school, email string, age uint) error {
	// 1. 检查用户名是否已经存在
	existingUser, _ := dao.FindUserByUsername(username)
	if existingUser != nil {
		return errors.New("用户名已存在")
	}

	// 2. 创建用户
	user := models.User{
		Username: username,
		Gender:   gender,
		Age:      age,
		School:   school,
		Email:    email,
	}

	// 3. 设置密码
	if err := user.SetPassword(password); err != nil {
		return err
	}

	// 4. 保存用户
	if err := dao.CreateUser(&user); err != nil {
		return err
	}

	return nil
}

func FindUserByNameAndPwd(username, password string) (*models.User, error) {
	user, err := dao.FindUserByNameAndPwd(username, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func FindUserByName(userName string) (*models.User, error) {
	user, err := dao.FindUserByUsername(userName)
	if err != nil {
		return nil, err
	}
	return user, nil
}
