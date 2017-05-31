package testing

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

//http://studygolang.com/articles/1513
func Test_Goconvey(t *testing.T) {
	Convey("一加一", t, func() {
		num := 1 + 1
		So(num, ShouldEqual, 2)
	})
}
