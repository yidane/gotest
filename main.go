package main

import (
	"log"
	"net/http"
	"runtime"

	"github.com/gorilla/mux"
)

func handle(hw http.ResponseWriter, request *http.Request) {

}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	r := mux.NewRouter()
	r.HandleFunc("/student", handle).Methods("GET")
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal(err)
	}
}
