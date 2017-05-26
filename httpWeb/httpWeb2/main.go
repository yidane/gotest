package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func main() {
	// 限制为1个CPU
	runtime.GOMAXPROCS(runtime.NumCPU())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world.")
	})

	log.Fatal(http.ListenAndServe(":8888", nil))
}
