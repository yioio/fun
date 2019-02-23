package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Trim_test(t *testing.T) {

	Convey("去除字符串首尾处的空白字符（或者其他字符）", t, func() {

		So(funString.Trim("abc "), ShouldEqual, "abc")
		So(funString.Trim(" abc"), ShouldEqual, "abc")
		So(funString.Trim(" abc "), ShouldEqual, "abc")
		So(funString.Trim("abc&", "&"), ShouldEqual, "abc")
		So(funString.Trim("&abc", "&"), ShouldEqual, "abc")
		So(funString.Trim("&abc&", "&"), ShouldEqual, "abc")
	})
}
