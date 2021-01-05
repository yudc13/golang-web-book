package model

import (
	"testing"
)

func TestUser_AddUser(t *testing.T)  {
	user := &User{
		Username: "xiaohuang",
		Pwd:      "123456",
		Email:    "xiaohuang@163.com",
	}
	result := user.AddUser()
	rows, _ := result.RowsAffected()
	id, _ := result.LastInsertId()
	t.Logf("rows: %d, id: %v \n ", rows, id)
	if rows != 1 {
		t.Logf("addUser error")
	}
}

func TestUser_AddUserByPrepare(t *testing.T) {
	user := &User{
		Username: "huangbo",
		Pwd:      "123456",
		Email:    "huangbo@163.com",
	}
	result := user.AddUserByPrepare()
	rows, _ := result.RowsAffected()
	id, _ := result.LastInsertId()
	t.Logf("rows: %d, id: %v \n ", rows, id)
	if rows != 1 {
		t.Logf("addUser error")
	}
}

func TestUser_QueryUserById(t *testing.T) {
	user := &User{
		Id: 7,
		Username: "",
		Pwd: "",
		Email: "",
	}
	user.QueryUserById()
	t.Logf("user: %v \n", user)
}

func TestUser_QueryUserList(t *testing.T) {
	user := &User{}
	users := user.QueryUserList()
	t.Logf("users: %v ", users)
}

func TestUser_UpdateUserNameById(t *testing.T) {
	user := User{
		Id: 1,
		Username: "李佳佳",
	}
	result := user.UpdateUserNameById()
	rows, _ := result.RowsAffected()
	t.Logf("rows: %d \n ", rows)
	if rows != 1 {
		t.Logf("UpdateUserNameById error")
	}
}