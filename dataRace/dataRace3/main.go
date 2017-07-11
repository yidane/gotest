package main

import (
	"fmt"
	"net/http"
	"time"
)

type Foo struct {
	content string
}

type FooSlice []*Foo

var requestChan chan chan FooSlice

func installHeepHandler(fs FooSlice) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("begin handle")
		request := make(chan FooSlice)
		requestChan <- request
		foolSlice := <-request
		fmt.Println("read data from channel")
		for _, foo := range foolSlice {
			if foo != nil {
				fmt.Fprintf(w, "foo: %v ", foo.content)
			}
		}
	}
	http.HandleFunc("/", handler)
}

func updateFooSlice(fs FooSlice) {
	for {
		select {
		case ch := <-requestChan:
			{
				fmt.Println("ch")
				fooSliceCopy := make(FooSlice, len(fs))
				copy(fooSliceCopy, fs)
				ch <- fooSliceCopy
			}
		default:
			{
				fs = make(FooSlice, 2)
				fs[0] = &Foo{content: "new"}
				fs[1] = &Foo{content: time.Now().String()}
			}
		}
	}
}

func main() {
	requestChan = make(chan chan FooSlice)
	fs := FooSlice{
		&Foo{content: "hello"},
		&Foo{content: time.Now().String()},
	}
	installHeepHandler(fs)
	go updateFooSlice(fs)

	http.ListenAndServe(":8080", nil)
}
