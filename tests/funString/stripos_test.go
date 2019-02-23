package funString


import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Stripos_test(t *testing.T) {

	Convey("查找字符串首次出现的位置", t, func() {

		So(funString.Stripos("aBc", "b", 0), ShouldEqual, 1)

		So(funString.Stripos("aBc", "d", 0), ShouldEqual, -1)

		So(funString.Stripos("aBc", "b", 2), ShouldEqual, -1)


	})
}
