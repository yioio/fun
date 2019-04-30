package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

//二进制安全比较字符串
func Test_Substr_compare(t *testing.T) {

	Convey("二进制安全比较字符串", t, func() {

		So(funString.Substr_compare("abcde", "bc", 1, 2), ShouldEqual, 0)
		So(funString.Substr_compare("abcde", "de", -2, 2), ShouldEqual, 0)
		So(funString.Substr_compare("abcde", "bcg", 1, 2), ShouldEqual, 0)
		So(funString.Substr_compare("abcde", "bc", 1, 3), ShouldEqual, 1)
		So(funString.Substr_compare("abcde", "cd", 1, 2), ShouldEqual, -1)
		So(funString.Substr_compare("abcde", "abc", 5, 1), ShouldEqual, -2)


	})

}
