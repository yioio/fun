package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Ucfirst_test(t *testing.T) {

	Convey("将字符串的首字母转换为大写", t, func() {

		So(funString.Ucfirst("hello"), ShouldEqual, "Hello")

	})
}