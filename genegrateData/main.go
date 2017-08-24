package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

var studentData map[int]Student
var ch chan int
var total int

func init() {
	size := 100000
	studentList := GenegrateStudentList(size)
	studentData = make(map[int]Student, size)
	for _, s := range studentList {
		studentData[s.No] = *s
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()*5)

	ch = make(chan int)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		ch <- 1

		no := r.URL.Query()["id"]

		if len(no) == 0 {
			w.Write([]byte("不存在参数id"))
			return
		}
		id, err := strconv.Atoi(no[0])
		if err != nil {
			w.Write([]byte("参数不合法"))
			return
		}
		if s, f := studentData[id]; f {
			w.Write([]byte(s.ToString()))
		} else {
			w.Write([]byte("不存在相关信息"))
		}
	})

	go func() {
		for {
			select {
			case <-ch:
				total++
			default:
			}
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println(total)
		}
	}()

	log.Fatal(http.ListenAndServe(":8888", nil))
}
