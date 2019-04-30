package funString

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
	"testing"
)

func Test_StrShuffle(t *testing.T) {

	Convey("使一个字符串的第一个字符小写", t, func() {

		So(funString.StrShuffle("abcdef"), ShouldNotEqual, "abcdef")

	})
}
