package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func StrReplace_test(t *testing.T) {

	Convey("子字符串替换", t, func() {

		So(funString.StrReplace("ce", "12", "abcef", 1), ShouldEqual, "ab12f")

	})
}
