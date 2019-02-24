package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Test_Stristr(t *testing.T) {

	Convey("查找字符串的首次出现", t, func() {

		So(funString.Strstr("aBcd", "bc", false), ShouldEqual, "bcd")
		So(funString.Strstr("abCd", "bc", true), ShouldEqual, "a")

	})
}