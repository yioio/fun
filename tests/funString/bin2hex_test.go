package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Test_Bin2hex(t *testing.T) {

	Convey("函数把包含数据的二进制字符串转换为十六进制值", t, func() {

		So(funString.Bin2hex("11"), ShouldEqual, "3")

	})
}