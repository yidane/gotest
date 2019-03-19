package redis

import (
	"fmt"
	test "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDel(t *testing.T) {
	test.Convey("TestDel", t, func() {
		conn := GetRedisConnect()
		defer conn.Close()

		const count = 22
		keys := make([]string, count)

		for i := 0; i < count; i++ {
			keys[i] = fmt.Sprintf("del:Key%v", i)
		}

		conn.Select(2)
		test.Convey("Set keys", func() {
			for i := 0; i < count; i++ {
				conn.SET(keys[i], keys[i])
			}
		})

		test.Convey("Exists keys before DELETE", func() {
			for i := 0; i < count; i++ {
				test.So(conn.EXISTS(keys[i]), test.ShouldBeTrue)
			}
		})

		test.Convey("DELETE keys", func() {
			test.So(conn.DEL(keys...), test.ShouldBeTrue)
		})

		test.Convey("Exists keys after DELETE", func() {
			for i := 0; i < count; i++ {
				test.So(conn.EXISTS(keys[i]), test.ShouldBeFalse)
			}
		})
	})
}

func TestString(t *testing.T) {
	test.Convey("TestString", t, func() {
		conn := GetRedisConnect()
		defer conn.Close()

		const testKey = "testStringKey"
		const oldValue = "00000"
		const newValue = "11111"

		test.Convey("Set OldValue", func() {
			conn.SET(testKey, oldValue)
			test.So(conn.EXISTS(testKey), test.ShouldBeTrue)
			test.So(conn.GETString(testKey), test.ShouldEqual, oldValue)
		})

		test.Convey("Set NewValue", func() {
			conn.SET(testKey, newValue)
			test.So(conn.EXISTS(testKey), test.ShouldBeTrue)
			test.So(conn.GETString(testKey), test.ShouldEqual, newValue)
		})

		test.Convey("SETNX exists key", func() {
			test.So(conn.SETNX(testKey, 123), test.ShouldBeFalse)
			test.So(conn.GETString(testKey), test.ShouldEqual, newValue)
		})

		test.Convey("Delete", func() {
			test.So(conn.DEL(testKey), test.ShouldBeTrue)
			test.So(conn.EXISTS(testKey), test.ShouldBeFalse)
		})
	})
}
