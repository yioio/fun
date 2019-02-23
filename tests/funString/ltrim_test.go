package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Ltrim_test(t *testing.T) {

	Convey("删除字符串开头的空白字符（或其他字符）", t, func() {

		So(funString.Trim(" abc"), ShouldEqual, "abc")
		So(funString.Trim("&abc", "&"), ShouldEqual, "abc")
	})
}
