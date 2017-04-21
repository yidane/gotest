package genegrateData

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/garyburd/redigo/redis"
)

func Test_CreateNewStudent(t *testing.T) {
	student := CreateNewStudent(1)
	fmt.Println(student)
	fmt.Println(student.ToString())
	student1 := ReCreateNewStudent(student.ToString())
	fmt.Println(student1)
	// if student.No != student1.No {
	// 	t.Error("序列化反序列化出错")
	// }
}

func Benchmark_CreateStudent(b *testing.B) {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < b.N; i++ {
		var student = CreateNewStudent(i).ToString()
		_, err := c.Do("SET", strconv.Itoa(i), student)
		if err != nil {
			fmt.Println(err)
		}
	}
}
