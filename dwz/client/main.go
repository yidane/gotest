package main

import (
	"log"
	"net/rpc"
)

type RedisHashArg struct {
	Key, Value string
}

type RedisResult struct {
	Result string
}

func main() {
	d, err := rpc.DialHTTP("tcp", "10.100.27.111:8989")
	if err != nil {
		log.Println(err)
		return
	}

	args := &RedisHashArg{Key: "yidane", Value: "yinsiwen"}
	result := RedisResult{}
	d.Call("RedisServer.SaveToRedis", args, &result)
}
