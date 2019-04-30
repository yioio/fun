package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Test_Rtrim(t *testing.T) {

	Convey("删除字符串末端的空白字符（或者其他字符）", t, func() {

		So(funString.Rtrim("abc "), ShouldEqual, "abc")
		So(funString.Rtrim("abc&", "&"), ShouldEqual, "abc")
	})
}
