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

		conn.SelectDb(1)
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
