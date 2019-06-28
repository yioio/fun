package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Test_Wordwrap(t *testing.T) {

	Convey("打断字符串为指定数量的字串", t, func() {

		So(funString.Wordwrap("hello world", 6, "<br/>"), ShouldEqual, "hello<br/>world")

	})

}
