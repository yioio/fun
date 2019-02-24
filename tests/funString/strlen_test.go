package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Test_Strlen(t *testing.T) {

	Convey("使用一个字符串分割另一个字符串", t, func() {

		So(funString.Strlen("abc"), ShouldEqual, 3)

	})
}
