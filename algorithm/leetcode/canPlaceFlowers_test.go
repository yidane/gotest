package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_canPlaceFlowers(t *testing.T) {
	Convey("测试花坛放置", t, func() {
		So(canPlaceFlowers([]int{1}, 0), ShouldEqual, true)
		So(canPlaceFlowers([]int{0}, 1), ShouldEqual, true)
		So(canPlaceFlowers([]int{0, 0}, 1), ShouldEqual, true)
		So(canPlaceFlowers([]int{0, 0}, 2), ShouldNotEqual, true)
		So(canPlaceFlowers([]int{0, 0, 0}, 2), ShouldEqual, true)
		So(canPlaceFlowers([]int{0, 0, 1}, 1), ShouldEqual, true)
		So(canPlaceFlowers([]int{0, 1, 0}, 1), ShouldNotEqual, true)
		So(canPlaceFlowers([]int{1, 0, 0}, 1), ShouldEqual, true)
		So(canPlaceFlowers([]int{0, 1, 0, 1}, 1), ShouldEqual, false)
		So(canPlaceFlowers([]int{0, 0, 1, 0, 0}, 2), ShouldEqual, true)
	})
}
