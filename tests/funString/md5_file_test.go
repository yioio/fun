package funString

import (
	"github.com/yioio/fun/funString"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Md5File(t *testing.T) {

	Convey("计算文件md5", t, func() {

		So(funString.Md5_file("./testFiles/pride-and-prejudice.txt"), ShouldEqual, "628a19b237928494a16e15f1bcd825cb")
	})
}
