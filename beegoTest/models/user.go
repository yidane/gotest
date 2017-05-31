package models

import "time"

//User 用户信息
type User struct {
	ID     int
	Name   string
	Gender bool
	Age    int
	Time   time.Time
}

//New 新建用户
func New() *User {
	return &User{
		ID:     1,
		Name:   "yidane",
		Gender: true,
		Age:    18,
		Time:   time.Now(),
	}
}
