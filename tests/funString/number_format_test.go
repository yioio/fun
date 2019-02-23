package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func NumberFormat_test(t *testing.T) {

	Convey("使用一个字符串分割另一个字符串", t, func() {

		So(funString.NumberFormat(1234.5678,2,".",""), ShouldEqual, "1234.57")

	})
}
