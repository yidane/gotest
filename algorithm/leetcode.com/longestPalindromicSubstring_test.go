package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_longestPalindrome(t *testing.T) {
	Convey("寻找最大回环子字符串", t, func() {
		So(longestPalindrome("abab"), ShouldEqual, "aba")
		So(longestPalindrome("aba"), ShouldEqual, "aba")
		So(longestPalindrome("a"), ShouldEqual, "a")
		So(longestPalindrome("ccc"), ShouldEqual, "ccc")
	})
}
