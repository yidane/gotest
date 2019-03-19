package redis

import (
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

func (conn *Conn) SelectDb(db int) {
	r, err := conn.Do("SELECT", db)
	test.So(err, test.ShouldBeNil)
	test.So(r, test.ShouldEqual, "OK")
}

func (conn *Conn) Close() {
	err := conn.Conn.Close()
	test.So(err, test.ShouldBeNil)
}
