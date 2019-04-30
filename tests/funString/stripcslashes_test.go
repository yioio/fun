package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Test_Stripcslashes(t *testing.T) {

	Convey("反引用一个使用 addcslashes() 转义的字符串", t, func() {

		So(funString.Stripslashes("Is your name O\\'reilly?"), ShouldEqual, "Is your name O'reilly?")

	})
}
