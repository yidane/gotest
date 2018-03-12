package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"sync/atomic"
)

var currentUserCount int64 = 0

func addUser(w http.ResponseWriter, r *http.Request) {
	atomic.AddInt64(&currentUserCount, 1)
	w.Write([]byte("ok"))
}

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(currentUserCount)

	w.Write([]byte(strconv.FormatInt(currentUserCount, 10)))
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world.")
	})

	http.HandleFunc("/add", addUser)
	http.HandleFunc("/get", getUser)

	log.Fatal(http.ListenAndServe(":8888", nil))
}
