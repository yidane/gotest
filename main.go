package main

import (
	"net/http"

	"fmt"

	"github.com/yidane/gotest/genegrateData"
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

	fmt.Println(genegrateData.CreateNewStudent(1).ToString())
}
