package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_findDaffodilNumber(t *testing.T) {
	FindDaffodilNumber(1, 100000000)
}
