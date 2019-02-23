package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Strpos_test(t *testing.T) {

	Convey("查找字符串首次出现的位置", t, func() {

		So(funString.Strpos("abc","b",0), ShouldEqual, 1)
		So(funString.Strpos("abc","d",0), ShouldEqual, -1)

	})
}