package funFilesystem

import (
	."github.com/smartystreets/goconvey/convey"
	"testing"
	"github.com/yioio/fun/funFilesystem"
)

func Test_Dirname(t *testing.T)  {
	Convey("获取路径的目录", t, func() {
		res := funFilesystem.Dirname("C:\\Users\\chenhongsheng\\Desktop\\fun\\doc\\zh\\test.md")

		So(res, ShouldEqual, "C:\\Users\\chenhongsheng\\Desktop\\fun\\doc\\zh")
	})
}

