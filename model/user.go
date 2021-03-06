package model

import (
	"database/sql"
	"fmt"
	"golangWebBook/db"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
	Email    string `json:"email"`
}

func FindUserByUsername(username string) (user User, err error) {
	query := "select id, pwd, username, email from users where username = ?"
	row := db.Db.QueryRow(query, username)
	user = User{
		Id:       0,
		Pwd:      "",
		Username: "",
		Email:    "",
	}
	err = row.Scan(&user.Id, &user.Pwd, &user.Username, &user.Email)
	if err != nil {
		return
	}
	return
}

func InsertUser(user *User) (int64, error) {
	query := "INSERT INTO users(username, pwd, email) values(?, ?, ?)"
	result, err := db.Db.Exec(query, user.Username, user.Pwd, user.Email)
	var lastInsertId int64
	if err != nil {
		return 0, err
	}
	if result != nil {
		lastInsertId, _ = result.LastInsertId()
	}
	return lastInsertId, nil
}

// 添加用户
func (user *User) AddUser() sql.Result {
	query := "INSERT INTO users(username, pwd, email) values(?, ?, ?)"
	result, err := db.Db.Exec(query, user.Username, user.Pwd, user.Email)
	if err != nil {
		panic(err)
	}
	return result
}

func (user *User) AddUserByPrepare() sql.Result {
	query := "INSERT INTO users(username, pwd, email) values(?, ?, ?)"
	stmt, err := db.Db.Prepare(query)
	if err != nil {
		return nil
	}
	result, err := stmt.Exec(user.Username, user.Pwd, user.Email)
	if err != nil {
		return nil
	}
	return result
}

// 根据id查询用户id
func (user *User) QueryUserById() User {
	query := "select id, username, pwd, email from users where id = ?"
	row := db.Db.QueryRow(query, user.Id)
	err := row.Scan(&user.Id, &user.Username, &user.Pwd, &user.Email)
	if err != nil {
		return User{}
	}
	return *user
}

func (user *User) QueryUserList() []User {
	var users []User
	query := "select id, username, pwd, email from users"
	rows, err := db.Db.Query(query)
	defer func() {
		err := rows.Close()
		if err != nil {
			fmt.Println("User QueryUserList Close Error ", err)
		}
	}()
	if err != nil {
		fmt.Println("User QueryUserList Error ", err)
		return []User{}
	}
	for rows.Next() {
		u := User{
			Id:       0,
			Username: "",
			Pwd:      "",
			Email:    "",
		}
		err = rows.Scan(&u.Id, &u.Username, &u.Pwd, &u.Email)
		if err != nil {
			panic(err)
		}
		users = append(users, u)
	}
	return users
}

func (user *User) UpdateUserNameById() sql.Result {
	query := "update users set username = ? where id = ?"
	result, err := db.Db.Exec(query, user.Username, user.Id)
	if err != nil {
		fmt.Println("UpdateUserNameById Error ", err)
		return nil
	}
	return result
}
