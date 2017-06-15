package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_reverse(t *testing.T) {
	Convey("整数转置", t, func() {
		So(reverse(1), ShouldEqual, 1)
		So(reverse(12), ShouldEqual, 21)
		So(reverse(123), ShouldEqual, 321)
		So(reverse(112345), ShouldEqual, 543211)
		So(reverse(10), ShouldEqual, 1)
		So(reverse(100), ShouldEqual, 1)
		So(reverse(-123), ShouldEqual, -321)
		So(reverse(90100), ShouldEqual, 109)
		So(reverse(-90000), ShouldEqual, -9)
	})
}
