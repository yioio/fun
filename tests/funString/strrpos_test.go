package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Strrpos_test(t *testing.T) {

	Convey("计算指定字符串在目标字符串中最后一次出现的位置（不区分大小写）", t, func() {

		So(funString.Strrpos("ababcd","ab",0), ShouldEqual, 2)

	})
}