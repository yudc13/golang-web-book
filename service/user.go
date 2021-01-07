package service

import (
	"errors"
	"golangWebBook/model"
)

func Login(username string, password string) (user model.User, err error) {
	u, err := model.FindUserByUsername(username)
	if err != nil {
		return model.User{}, errors.New("用户名或密码错误")
	}
	if u.Pwd != password {
		return model.User{}, errors.New("用户名或密码错误")
	}
	// 不返回密码字段
	user = model.User{
		Id:       u.Id,
		Username: u.Username,
		Email:    u.Email,
	}
	return user, nil
}

func Register(user *model.User) (userId int64, err error) {
	// 判断用户名是否已经注册过
	_, err = model.FindUserByUsername(user.Username)
	if err == nil {
		return 0, errors.New("用户名已经注册")
	}
	// 添加用户
	userId, err = model.InsertUser(user)
	return userId, err
}
