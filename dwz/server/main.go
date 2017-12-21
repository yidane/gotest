package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type RedisServer struct {
}

type RedisHashArg struct {
	Key, Value string
}

type RedisResult struct {
	Result string
}

//SaveToRedis save to redis
func (server *RedisServer) SaveToRedis(arg *RedisHashArg, r *RedisResult) error {
	log.Println(arg.Key, arg.Value)
	r.Result = "succeed"
	return nil
}

func main() {
	rpc.Register(&RedisServer{})
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Println(err)
		return
	}

	http.Serve(l, nil)
}
