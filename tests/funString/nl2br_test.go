package funString


import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Test_Nl2br(t *testing.T) {

	Convey("函数把包含数据的二进制字符串转换为十六进制值", t, func() {

		So(funString.Nl2br("foo isn't\n bar", true), ShouldEqual, "foo isn't<br /> bar")

	})
}