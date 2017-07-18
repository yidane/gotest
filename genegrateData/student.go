package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Student struct {
	No     int
	Name   string
	Sex    string
	YuWen  int
	ShuXue int
	YingYu int
	WuLi   int
	HuaXue int
}

func CreateNewStudent(no int) *Student {
	student := &Student{}
	student.No = no
	if no%2 == 0 {
		student.Sex = "男"
	} else {
		student.Sex = "女"
	}
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	student.YuWen = creatScore(no, rand)
	student.ShuXue = creatScore(no, rand)
	student.WuLi = creatScore(no, rand)
	student.YingYu = creatScore(no, rand)
	student.HuaXue = creatScore(no, rand)

	return student
}

func ReCreateNewStudent(jsonData string) *Student {
	if jsonData == "" {
		fmt.Println("empty json data")
	}
	var student Student
	err := json.Unmarshal([]byte(jsonData), student)
	if err == nil {
		return nil
	}

	return &student
}

func (student Student) ToString() string {
	b, err := json.Marshal(student)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}

func creatScore(no int, rand *rand.Rand) int {
	return 44 + 56*rand.Intn(100)/100
}
