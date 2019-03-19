package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	test "github.com/smartystreets/goconvey/convey"
)

var pool redis.Pool
var localLocation = "127.0.0.1:6379"

func init() {
	pool = redis.Pool{MaxActive: 2, MaxIdle: 5, Dial: func() (conn redis.Conn, e error) {
		return connect(localLocation)
	}}
}

func connect(location string) (redis.Conn, error) {
	return redis.Dial("tcp", location)
}

type Conn struct {
	redis.Conn
}

func GetRedisConnect() *Conn {
	con := pool.Get()

	return &Conn{con}
}

func (conn *Conn) Select(db int) {
	r, err := conn.Do("SELECT", db)
	test.So(err, test.ShouldBeNil)
	test.So(r, test.ShouldEqual, "OK")
}

func (conn *Conn) Close() {
	err := conn.Conn.Close()
	test.So(err, test.ShouldBeNil)
}

func (conn *Conn) EXISTS(key string) bool {
	r, err := redis.Bool(conn.Do("EXISTS", key))
	test.So(err, test.ShouldBeNil)

	return r
}

func (conn *Conn) Type(key string) string {
	r, err := redis.String(conn.Do("TYPE", key))
	test.So(err, test.ShouldBeNil)

	return r
}

func (conn *Conn) RENAME(key, newKey string) error {
	_, err := redis.String(conn.Do("RENAME", key, newKey))
	return err
}

func (conn *Conn) RENAMENX(key, newKey string) bool {
	r, err := redis.Bool(conn.Do("RENAMENX", key, newKey))
	if err != nil {
		fmt.Println(err)
	}
	return r
}

func (conn *Conn) MOVE(key string, db int) bool {
	r, err := redis.Bool(conn.Do("MOVE", key, db))
	test.So(err, test.ShouldBeNil)

	return r
}

func (conn *Conn) RANDOMKEY() string {
	if conn.DBSIZE() == 0 {
		return ""
	}

	r, err := redis.String(conn.Do("RANDOMKEY"))
	test.So(err, test.ShouldBeNil)
	return r
}

func (conn *Conn) DBSIZE() int64 {
	r, err := redis.Int64(conn.Do("DBSIZE"))
	test.So(err, test.ShouldBeNil)

	return r
}

func (conn *Conn) Lock(key, value string, ttl int) (bool, error) {
	r, err := redis.String(conn.Do("SET", key, value, "EX", ttl, "NX"))
	if err != nil {
		if err == redis.ErrNil {
			return false, nil
		}
		return false, err
	}

	return r == "OK", nil
}

func (conn *Conn) UnLock(key string) bool {
	return conn.DEL(key)
}
