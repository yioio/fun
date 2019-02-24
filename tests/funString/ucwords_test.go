package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Test_UcWords(t *testing.T) {

	Convey("将字符串中每个单词的首字母转换为大写", t, func() {

		So(funString.Ucfirst("hello world"), ShouldEqual, "Hello World")

	})
}
