package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Test_Ord(t *testing.T) {

	Convey("转换字符串第一个字节为 0-255 之间的值", t, func() {

		So(funString.Ord("A"), ShouldEqual, 65)

	})
}
