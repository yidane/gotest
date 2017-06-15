package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_myAtoi(t *testing.T) {
	Convey("字符串转换为整形", t, func() {
		So(myAtoi(""), ShouldEqual, 0)
		So(myAtoi("1"), ShouldEqual, 1)
		So(myAtoi("+1"), ShouldEqual, 1)
		So(myAtoi("+-2"), ShouldEqual, 0)
		So(myAtoi("1 "), ShouldEqual, 1)
		So(myAtoi("1230"), ShouldEqual, 1230)
		So(myAtoi("-123"), ShouldEqual, -123)
		So(myAtoi("  -0012a42"), ShouldEqual, -12)
		So(myAtoi("2147483648"), ShouldEqual, 2147483647)
		So(myAtoi("-2147483648"), ShouldEqual, -2147483648)
	})
}
