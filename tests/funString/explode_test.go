package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Explode_test(t *testing.T) {

	Convey("使用一个字符串分割另一个字符串", t, func() {

		So(funString.Explode("|", "abc|bcd|aaa"), ShouldEqual, []string{"abc","bcd","aaa"})

	})
}