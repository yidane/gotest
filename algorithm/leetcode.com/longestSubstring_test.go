package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	Convey("", t, func() {
		So(lengthOfLongestSubstring("abcabcabc"), ShouldEqual, 3)
		So(lengthOfLongestSubstring("bbbbb"), ShouldEqual, 1)
		So(lengthOfLongestSubstring("pwwkew"), ShouldEqual, 3)
		So(lengthOfLongestSubstring("abba"), ShouldEqual, 2)
		So(lengthOfLongestSubstring("uqinntq"), ShouldEqual, 4)
	})
}
