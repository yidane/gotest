package redis

import (
	"github.com/garyburd/redigo/redis"
	test "github.com/smartystreets/goconvey/convey"
)

func (conn *Conn) GET(key string) interface{} {
	r, err := conn.Do("GET", key)
	test.So(err, test.ShouldBeNil)

	return r
}

func (conn *Conn) GETString(key string) string {
	r, err := redis.String(conn.Do("GET", key))
	test.So(err, test.ShouldBeNil)

	return r
}

func (conn *Conn) GETInt(key string) int {
	r, err := redis.Int(conn.Do("GET", key))
	test.So(err, test.ShouldBeNil)

	return r
}

func (conn *Conn) GETIntSlice(key string) []int {
	r, err := redis.Ints(conn.Do("GET", key))
	test.So(err, test.ShouldBeNil)

	return r
}

func (conn *Conn) GETInt64(key string) int64 {
	r, err := redis.Int64(conn.Do("GET", key))
	test.So(err, test.ShouldBeNil)

	return r
}

func (conn *Conn) GETInt64Map(key string) map[string]int64 {
	r, err := redis.Int64Map(conn.Do("GET", key))
	test.So(err, test.ShouldBeNil)

	return r
}

func (conn *Conn) GETInt64Slice(key string) []int64 {
	r, err := redis.Int64s(conn.Do("GET", key))
	test.So(err, test.ShouldBeNil)

	return r
}

func (conn *Conn) GETFloat64(key string) float64 {
	r, err := redis.Float64(conn.Do("GET", key))
	test.So(err, test.ShouldBeNil)

	return r
}

func (conn *Conn) GETIFloat64Slice(key string) []float64 {
	r, err := redis.Float64s(conn.Do("GET", key))
	test.So(err, test.ShouldBeNil)

	return r
}

//若不存在，直接写入值；若存在，则覆盖旧值
func (conn *Conn) SET(key string, values interface{}) {
	r, err := redis.String(conn.Do("SET", key, values))
	test.So(err, test.ShouldBeNil)
	test.So(r, test.ShouldEqual, "OK")
}

func (conn *Conn) MSET(key string, values map[string]interface{}) {
	if len(values) == 0 {
		return
	}

	args := make([]interface{}, len(values)*2)
	i := 0
	for k, v := range values {
		args[i] = k
		args[i+1] = v
		i += 2
	}

	r, err := redis.String(conn.Do("MSET", args))
	test.So(err, test.ShouldBeNil)
	test.So(r, test.ShouldEqual, "OK")
}

func (conn *Conn) MSETNX(key string, values map[string]interface{}) {
	if len(values) == 0 {
		return
	}

	args := make([]interface{}, len(values)*2)
	i := 0
	for k, v := range values {
		args[i] = k
		args[i+1] = v
		i += 2
	}

	r, err := redis.String(conn.Do("MSETNX", args))
	test.So(err, test.ShouldBeNil)
	test.So(r, test.ShouldEqual, "OK")
}

//SET if Not eXists
func (conn *Conn) SETNX(key string, values interface{}) bool {
	r, err := redis.Int(conn.Do("SETNX", key, values))
	test.So(err, test.ShouldBeNil)

	return r == 1
}

func (conn *Conn) DEL(keys ...string) bool {
	if len(keys) == 0 {
		return true
	}

	if len(keys) == 1 {
		c, err := redis.Int(conn.Do("DEL", keys[0]))
		test.So(err, test.ShouldBeNil)

		if c == 1 {

		}

		return true
	}

	return conn.dels(keys...)
}

func (conn *Conn) dels(keys ...string) bool {
	if len(keys) == 0 {
		return true
	}

	err := conn.Send("MULTI")
	test.So(err, test.ShouldBeNil)

	for _, key := range keys {
		if len(key) == 0 {
			continue
		}

		err = conn.Send("DEL", key)
		test.So(err, test.ShouldBeNil)

	}

	err = conn.Send("EXEC")
	test.So(err, test.ShouldBeNil)

	return true
}
