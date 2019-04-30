package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Test_Chr(t *testing.T) {

	Convey("返回指定的字符", t, func() {

		So(funString.Chr(65), ShouldEqual, "A")

	})
}