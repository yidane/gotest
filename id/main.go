package main

import (
	"fmt"
	"github.com/naoina/denco"
	"github.com/yidane/toolkits/id"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"sync"
	"sync/atomic"
)

func main() {
	go func() {
		http.ListenAndServe(":6789", nil)
	}()

	mux := denco.NewMux()
	handler, err := mux.Build([]denco.Handler{
		mux.GET("/", index),
		mux.GET("/id/:stub", newID),
	})

	if err != nil {
		panic(err)
	}

	log.Fatal(http.ListenAndServe(":8088", handler))
}

func index(w http.ResponseWriter, r *http.Request, params denco.Params) {
	w.Write([]byte("hello world"))
}

var snowflakes = make(map[int]*id.Snowflake)
var rwLock = sync.RWMutex{}

var total int32 = 0

func newID(w http.ResponseWriter, r *http.Request, params denco.Params) {
	atomic.AddInt32(&total, 1)
	fmt.Println(total)
	stubStr := params.Get("stub")
	stub, err := strconv.Atoi(stubStr)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	snowflake, ok := snowflakes[stub]

	if !ok {
		rwLock.Lock()
		defer rwLock.Unlock()
		snowflake, ok = snowflakes[stub]
		if !ok {
			snowflake, err = id.NewSnowflake(0, uint(stub), 0)
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			snowflakes[stub] = snowflake
		}
	}

	snowflakeID, err := snowflake.NewID()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(strconv.FormatInt(snowflakeID, 10)))
}
