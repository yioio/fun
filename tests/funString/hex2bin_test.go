package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Test_Hex2bin(t *testing.T) {

	Convey("函数把包含数据的二进制字符串转换为十六进制值", t, func() {

		So(funString.Hex2bin("3"), ShouldEqual, "11")

	})
}