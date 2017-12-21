package redis

import "errors"
import "github.com/garyburd/redigo/redis"

//RedisHelper for operate redis
type RedisHelper struct {
	Server string
	con    *redis.Conn
}

func (r *RedisHelper) Dail() error {
	if len(r.Server) == 0 {
		return errors.New("Server can not be null")
	}

	c, err := redis.Dial("tcp", r.Server)
	if err != nil {
		return err
	}

	r.con = &c
	return nil
}
