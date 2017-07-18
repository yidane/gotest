package main

import "github.com/garyburd/redigo/redis"
import "fmt"
import "strconv"

func Genegrate(total int) {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	if total < 10 {
		total = 10
	}
	for i := 0; i < total; i++ {
		_, err := c.Do("SET", strconv.Itoa(i), "hello")
		if err != nil {
			fmt.Println(err)
		}
	}
}

//GenegrateStudentList 生产数据
func GenegrateStudentList(total int) []*Student {
	students := make([]*Student, total)
	for i := 0; i < total; i++ {
		students[i] = CreateNewStudent(i)
	}

	return students
}
