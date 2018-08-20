package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

//HOSTNAME 主机名
var HOSTNAME = ""

func init() {
	n, err := os.Hostname()
	if err == nil {
		HOSTNAME = n
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, fmt.Sprintf("Recived 'Hello world' from %v at %v", HOSTNAME, time.Now().Unix()))
	})

	log.Fatal(http.ListenAndServe(":8889", nil))
}
