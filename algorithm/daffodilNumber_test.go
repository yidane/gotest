package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_findDaffodilNumber(t *testing.T) {
	Convey("寻找水仙数", t, func() {
		FindDaffodilNumber(1, 100000000)
	})
}
