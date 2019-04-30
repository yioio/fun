package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Test_Implode(t *testing.T) {

	Convey("将一个一维数组的值转化为字符串", t, func() {

		So(funString.Implode("|", []string{"abc","aaa"}), ShouldEqual, "abc|aaa")

	})
}