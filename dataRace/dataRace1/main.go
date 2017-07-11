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

func updateFooSlice(foolSlice FooSlice) {
	for {
		foo := &Foo{content: "new"}
		foolSlice[0] = foo
		time.Sleep(time.Second)
	}
}

func installHeepHandler(foolSlice FooSlice) {
	handler := func(w http.ResponseWriter, r *http.Request) {
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

//1、使用go run -race main.go运行
//2、在浏览器中访问http://localhost:8080,在Terminal中会输出 DATA RACE.
