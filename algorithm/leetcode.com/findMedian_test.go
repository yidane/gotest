package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_findMedianSortedArrays(t *testing.T) {
	Convey("寻找中值", t, func() {
		So(findMedianSortedArrays([]int{1, 2}, []int{2}), ShouldEqual, 2)
	})
}
