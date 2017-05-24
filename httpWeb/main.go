package main

import (
	"net/http"

	"fmt"

	"time"

	"log"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n", time.Now().String())
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello %s!\n", ps.ByName("name"))
}

func (nf NotFoundInfo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "404!\n这是我的自定义404错误\n", time.Now().String())
}

type NotFoundInfo struct {
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	//router.NotFound = http.FileServer(http.Dir("public"))
	router.NotFound = new(NotFoundInfo)

	log.Fatal(http.ListenAndServe(":8888", router))
}
