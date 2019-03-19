package redis

import (
	"fmt"
	test "github.com/smartystreets/goconvey/convey"
	"strconv"
	"testing"
	"time"
)

func Test_Connect(t *testing.T) {
	tests := []struct {
		name     string
		location string
		wantErr  bool
	}{
		{name: "localhost", location: localLocation, wantErr: false},
		{name: "localhostErrorPort", location: "127.0.0.1:6380", wantErr: true},
		{name: "emptyServer", location: "", wantErr: true},
	}

	test.Convey("connect the local redis", t, func() {
		for _, tt := range tests {
			test.Convey(tt.name, func() {
				conn, err := connect(tt.location)
				if tt.wantErr {
					test.So(err, test.ShouldNotBeNil)
					return
				}

				test.So(err, test.ShouldBeNil)
				defer func() {
					err := conn.Close()
					test.ShouldBeNil(err)
				}()
			})
		}
	})
}

func Test_GetRedisConnect(t *testing.T) {
	test.Convey("GetRedisConnect", t, func() {
		conn := GetRedisConnect()
		test.ShouldBeNil(conn.Err())
		defer conn.Close()
	})
}

func TestConn_SelectDb(t *testing.T) {
	test.Convey("TestConn_SelectDb", t, func() {
		conn := GetRedisConnect()
		defer conn.Close()

		conn.Select(1)
	})
}

func TestConn_Type(t *testing.T) {
	test.Convey("TestConn_Type", t, func() {
		conn := GetRedisConnect()
		defer conn.Close()

		randKey := strconv.FormatInt(time.Now().UnixNano(), 10)
		test.So(conn.Type(randKey), test.ShouldEqual, "none")

		conn.SET(randKey, time.Now().Unix())
		test.So(conn.Type(randKey), test.ShouldEqual, "string")

		conn.DEL(randKey)
	})
}

func TestConn_RENAME(t *testing.T) {
	test.Convey("TestConn_RENAME", t, func() {
		conn := GetRedisConnect()
		defer conn.Close()

		oldKey, newKey := strconv.FormatInt(time.Now().UnixNano(), 10), strconv.FormatInt(time.Now().UnixNano(), 10)

		test.Convey("RENAME dot not exists key", func() {
			renameResult := conn.RENAME(oldKey, newKey)
			test.So(renameResult, test.ShouldNotBeNil)
			fmt.Println(renameResult)
		})

		if !conn.EXISTS(oldKey) {
			conn.SET(oldKey, oldKey)
		}

		test.So(conn.EXISTS(newKey), test.ShouldBeFalse)
		test.So(conn.RENAME(oldKey, newKey), test.ShouldBeNil)

		test.So(conn.EXISTS(oldKey), test.ShouldBeFalse)
		test.So(conn.EXISTS(newKey), test.ShouldBeTrue)

		test.So(conn.GETString(newKey), test.ShouldEqual, oldKey)

		conn.DEL(newKey)
		test.So(conn.EXISTS(oldKey), test.ShouldBeFalse)
		test.So(conn.EXISTS(newKey), test.ShouldBeFalse)
	})
}

func TestConn_RENAMENX(t *testing.T) {
	test.Convey("TestConn_RENAME", t, func() {
		conn := GetRedisConnect()
		defer conn.Close()

		oldKey, newKey := strconv.FormatInt(time.Now().UnixNano(), 10), strconv.FormatInt(time.Now().UnixNano(), 10)

		//如果当前键不存在，失败
		test.Convey("oldKey dot not exists", func() {
			test.So(conn.RENAMENX(oldKey, newKey), test.ShouldNotBeNil)
		})

		if !conn.EXISTS(oldKey) {
			conn.SET(oldKey, oldKey)
		}
		test.So(conn.EXISTS(oldKey), test.ShouldBeTrue)

		//如果相同键替换，失败
		test.Convey("RENAMENX by same key", func() {
			test.So(conn.RENAMENX(oldKey, oldKey), test.ShouldBeFalse)
		})

		test.So(conn.EXISTS(newKey), test.ShouldBeFalse)

		//如果新键不存在，成功
		test.So(conn.RENAMENX(oldKey, newKey), test.ShouldBeTrue)

		test.So(conn.EXISTS(oldKey), test.ShouldBeFalse)
		test.So(conn.EXISTS(newKey), test.ShouldBeTrue)
		test.So(conn.GETString(newKey), test.ShouldEqual, oldKey)

		test.So(conn.DEL(newKey), test.ShouldBeTrue)

		test.So(conn.EXISTS(oldKey), test.ShouldBeFalse)
		test.So(conn.EXISTS(newKey), test.ShouldBeFalse)
	})
}

func TestConn_MOVE(t *testing.T) {
	test.Convey("TestConn_MOVE", t, func() {
		conn := GetRedisConnect()
		defer conn.Close()

		key := strconv.FormatInt(time.Now().UnixNano(), 10)

		conn.Select(0)
		test.So(conn.EXISTS(key), test.ShouldBeFalse)
		conn.Select(1)
		test.So(conn.EXISTS(key), test.ShouldBeFalse)

		//直接移动
		conn.Select(0)
		test.So(conn.MOVE(key, 1), test.ShouldBeFalse)
		test.So(conn.EXISTS(key), test.ShouldBeFalse)
		conn.Select(1)
		test.So(conn.EXISTS(key), test.ShouldBeFalse)

		//在0中添加key
		conn.Select(0)
		conn.SET(key, key)
		test.So(conn.EXISTS(key), test.ShouldBeTrue)
		conn.Select(1)
		test.So(conn.EXISTS(key), test.ShouldBeFalse)

		//从0移动到1
		conn.Select(0)
		test.So(conn.MOVE(key, 1), test.ShouldBeTrue)
		test.So(conn.EXISTS(key), test.ShouldBeFalse)

		conn.Select(1)
		test.So(conn.EXISTS(key), test.ShouldBeTrue)

		//再在0中添加相同key，然后移动
		conn.Select(0)
		conn.SET(key, key)
		test.So(conn.EXISTS(key), test.ShouldBeTrue)
		conn.Select(1)
		test.So(conn.EXISTS(key), test.ShouldBeTrue)

		//再次从0移动到1
		conn.Select(0)
		test.So(conn.MOVE(key, 1), test.ShouldBeFalse)
		test.So(conn.EXISTS(key), test.ShouldBeTrue)

		conn.Select(1)
		test.So(conn.EXISTS(key), test.ShouldBeTrue)

		//删除key
		conn.Select(0)
		test.So(conn.DEL(key), test.ShouldBeTrue)
		conn.Select(1)
		test.So(conn.DEL(key), test.ShouldBeTrue)
	})
}

func TestConn_RANDOMKEY(t *testing.T) {
	test.Convey("TestConn_RANDOMKEY", t, func() {
		conn := GetRedisConnect()
		defer conn.Close()

		conn.Select(15)
		test.So(conn.RANDOMKEY(), test.ShouldBeBlank)
	})
}

func TestConn_DBSIZE(t *testing.T) {
	test.Convey("TestConn_DBSIZE", t, func() {
		conn := GetRedisConnect()
		defer conn.Close()

		oldSize := conn.DBSIZE()
		key := strconv.FormatInt(time.Now().UnixNano(), 10)

		conn.SET(key, key)

		newSize := conn.DBSIZE()

		test.So(newSize, test.ShouldEqual, oldSize+1)

		test.So(conn.DEL(key), test.ShouldBeTrue)
		newSize = conn.DBSIZE()
		test.So(newSize, test.ShouldEqual, oldSize)
	})
}

func TestConn_Lock(t *testing.T) {
	test.Convey("", t, func() {
		conn := GetRedisConnect()
		defer conn.Close()

		key := strconv.FormatInt(time.Now().UnixNano(), 10)

		f, err := conn.Lock(key, key, 20)
		test.So(err, test.ShouldBeNil)
		test.So(f, test.ShouldBeTrue)

		f, err = conn.Lock(key, key, 20)
		test.So(err, test.ShouldBeNil)
		test.So(f, test.ShouldBeFalse)

		conn.UnLock(key)

		f, err = conn.Lock(key, key, 20)
		test.So(err, test.ShouldBeNil)
		test.So(f, test.ShouldBeTrue)

		conn.UnLock(key)
	})
}
