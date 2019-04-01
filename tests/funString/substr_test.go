package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Test_Substr(t *testing.T) {

	Convey("使用一个字符串分割另一个字符串", t, func() {

		So(funString.Substr("abcdef", 1, 2), ShouldEqual, "bc")

	})
}