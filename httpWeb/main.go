package main

import (
	"net/http"

	"fmt"

	"time"

	"log"

	"sync"

	"bytes"

	"runtime"

	"github.com/garyburd/redigo/redis"
	"github.com/julienschmidt/httprouter"
)

var Lock sync.Mutex
var Total int

//Index Say Welcome to client
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n", time.Now().String())
}

//Hello Say hello to client
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello %s!\n", ps.ByName("name"))
}

func (nf NotFoundInfo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "404!\n这是我的自定义404错误\n", time.Now().String())
}

//Tick 返回当前服务访问次数
func Tick(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	Lock.Lock()
	defer Lock.Unlock()
	Total++
	fmt.Fprint(w, Total)
}

var c redis.Conn

//Get 查询获取redis值
func Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// go func() {
	// 	Lock.Lock()
	// 	defer Lock.Unlock()
	// 	Total++
	// }()

	chn := make(chan string)
	var key = ps.ByName("id")
	//dialRedis()
	go func(key string) {
		d, err := c.Do("GET", key)
		if err != nil {
			chn <- err.Error()
		}
		if d == nil {
			chn <- "未找到相应值"
		} else {
			result := new(bytes.Buffer)
			result.Write(d.([]byte))
			chn <- result.String()
		}
	}(key)

	fmt.Fprint(w, <-chn)
}

//NotFoundInfo 自定义404
type NotFoundInfo struct {
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/Tick", Tick)
	router.GET("/GET/:id", Get)
	//router.NotFound = http.FileServer(http.Dir("public"))
	router.NotFound = new(NotFoundInfo)

	log.Fatal(http.ListenAndServe(":8888", router))
}

func init() {
	if !dialRedis() {
		fmt.Println("链接Redis服务失败")
	}
}

func dialRedis() bool {
	if c != nil {
		if err := c.Err(); err != nil {
			panic(err)
		}
		return true
	}

	var err error
	c, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}

	return true
}
