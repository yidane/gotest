package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"
)

func handle(hw http.ResponseWriter, request *http.Request) {
	hw.Write([]byte("123"))
}

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	// r := mux.NewRouter()
	// r.HandleFunc("/student", handle).Methods("GET")
	// err := http.ListenAndServe(":8082", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//fmt.Println(genegrateData.CreateNewStudent(1).ToString())

	var n = 1200

	var max = 0
	var t = 0
	for m := 1; m < 60; m++ {
		i := n / m
		r := n - i*m
		if r > m {
			continue
		}

		fmt.Printf("%v %s %v = %v	%v	\n", n, "%", m, i, r)

		if r > max {
			max = r
			t = m
		}
	}

	go func() {
		fmt.Println("runtime.Goexit()")
		runtime.Goexit()
		fmt.Println("runtime.Goexit() end")
	}()

	fmt.Println(max)
	fmt.Println(t)

	time.Sleep(time.Millisecond)
	os.Exit(0)
}
