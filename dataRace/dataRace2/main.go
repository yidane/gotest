package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Foo struct {
	content string
}

type FooSlice []*Foo

var mutex sync.Mutex

func updateFooSlice(foolSlice FooSlice) {
	for {
		mutex.Lock()
		foo := &Foo{content: "new"}
		foolSlice[0] = foo
		mutex.Unlock()
		time.Sleep(time.Second)
	}
}

func installHeepHandler(foolSlice FooSlice) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		defer mutex.Unlock()
		for _, foo := range foolSlice {
			if foo != nil {
				fmt.Fprintf(w, "foo: %v ", foo.content)
			}
		}
	}
	http.HandleFunc("/", handler)
}

func main() {
	foo1 := &Foo{content: "hey"}
	foo2 := &Foo{content: "yo"}
	fooSlice := FooSlice{foo1, foo2}

	installHeepHandler(fooSlice)
	go updateFooSlice(fooSlice)

	http.ListenAndServe(":8080", nil)
}
