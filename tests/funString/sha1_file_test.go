package funString


import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Test_Sha1_file(t *testing.T) {

	Convey("使用反斜线引用字符串", t, func() {

		So(funString.Sha1_file("md5_test.go"), ShouldEqual, "25933fc94d3ae37245bddfb3b7cb04350e840272")

	})
}