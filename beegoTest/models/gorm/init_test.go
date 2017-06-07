package gorm

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Init(t *testing.T) {
	Convey("测试数据库自动迁移", t, func() {
		So(func() bool {
			return Init() == nil
		}(), ShouldEqual, true)
		So(HasUerTable(), ShouldEqual, true)
	})
}
