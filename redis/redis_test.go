package redis

import (
	test "github.com/smartystreets/goconvey/convey"
	"testing"
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
