package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Test_Addslashes(t *testing.T) {

	Convey("使用反斜线引用字符串", t, func() {

		So(funString.Addslashes("Is your name O'reilly?"), ShouldEqual, "Is your name O\\'reilly?")

	})
}