package leetcode

import (
	"testing"

	"math"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_power(t *testing.T) {
	Convey("Power", t, func() {
		So(myPow(0, 0), ShouldEqual, math.Pow(0, 0))
		So(myPow(3, -2), ShouldEqual, math.Pow(3, -2))
		So(myPow(10, 0), ShouldEqual, math.Pow(10, 0))
		So(myPow(1.3, -3), ShouldEqual, math.Pow(1.3, -3))
		So(myPow(0, -3), ShouldEqual, math.Pow(0, -3))
		So(myPow(6, 6), ShouldEqual, math.Pow(6, 6))
		So(myPow(math.MaxFloat64, 2), ShouldEqual, math.Pow(math.MaxFloat64, 2))
		So(myPow(0.00001, 2147483647), ShouldEqual, math.Pow(0.00001, 2147483647))
		So(myPow(0, 0), ShouldEqual, math.Pow(0, 0))
		So(myPow(0, 0), ShouldEqual, math.Pow(0, 0))
		So(myPow(0, 0), ShouldEqual, math.Pow(0, 0))
		So(myPow(0, 0), ShouldEqual, math.Pow(0, 0))
		So(myPow(0, 0), ShouldEqual, math.Pow(0, 0))
		So(myPow(0, 0), ShouldEqual, math.Pow(0, 0))
	})
}
