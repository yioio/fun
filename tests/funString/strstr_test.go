package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Test_Strstr(t *testing.T) {

	Convey("查找字符串的首次出现", t, func() {

		So(funString.Strstr("abcd", "bc", false), ShouldEqual, "bcd")
		So(funString.Strstr("abcd", "bc", true), ShouldEqual, "a")

	})
}
