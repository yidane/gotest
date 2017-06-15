package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_searchRange(t *testing.T) {
	Convey("測試尋找相同值範圍", t, func() {
		So(func() bool {
			actual := searchRange([]int{1, 2, 3, 4, 4, 5, 6}, 4)
			return actual[0] == 3 && actual[1] == 4
		}(), ShouldEqual, true)
		So(func() bool {
			actual := searchRange([]int{}, 1)
			return actual[0] == -1 && actual[1] == -1
		}(), ShouldEqual, true)
		So(func() bool {
			actual := searchRange([]int{1, 2}, 3)
			return actual[0] == -1 && actual[1] == -1
		}(), ShouldEqual, true)
		So(func() bool {
			actual := searchRange([]int{1, 3}, 1)
			return actual[0] == 0 && actual[1] == 0
		}(), ShouldEqual, true)
		So(func() bool {
			actual := searchRange([]int{2, 2}, 2)
			return actual[0] == 0 && actual[1] == 1
		}(), ShouldEqual, true)
		So(func() bool {
			actual := searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
			return actual[0] == 3 && actual[1] == 4
		}(), ShouldEqual, true)
	})
}
