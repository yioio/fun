package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Lcfirst_test(t *testing.T) {

	Convey("使一个字符串的第一个字符小写", t, func() {

		So(funString.Lcfirst("Abcd"), ShouldEqual, "abcd")

	})
}
