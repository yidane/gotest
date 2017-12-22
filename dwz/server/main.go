package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/garyburd/redigo/redis"
)

var con redis.Conn

type RedisServer struct {
	Count int64
}

type RedisHashArg struct {
	Key, Value string
}

type RedisResult struct {
	Result string
}

//SaveToRedis save to redis
func (server *RedisServer) SaveToRedis(arg *[]RedisHashArg, r *RedisResult) error {

	for _, a := range *arg {
		if err := con.Send("HMSET", "dwz.cn", a.Key, a.Value); err != nil {
			r.Result = "faild"
			fmt.Println(err)
			return err
		}
		server.Count++
	}
	if err := con.Flush(); err != nil {
		r.Result = "faild"
		fmt.Println(err)
		return err
	}
	fmt.Println(server.Count)
	r.Result = "succeed"
	return nil
}

func main() {
	var err error
	con, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	con.Do("select", 1)

	defer con.Close()
	log.Println("connect redis success")

	rpc.Register(&RedisServer{})
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":8989")
	if err != nil {
		log.Println(err)
		return
	}

	http.Serve(l, nil)
}
