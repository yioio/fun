package temp

import (
	."github.com/smartystreets/goconvey/convey"
	"testing"
	"github.com/yioio/fun/funFilesystem"
)

func Test_Copy(t *testing.T)  {
	Convey("拷贝文件", t, func() {
		src := "C:\\Users\\chenhongsheng\\Desktop\\fun\\doc\\zh\\test.md"
		dst :=  "C:\\Users\\chenhongsheng\\Desktop\\fun\\doc\\zh\\test2.md"
		res := funFilesystem.Copy(src, dst)
		res2 := funFilesystem.File_exits(dst)
		res3 := funFilesystem.Delete(dst)

		So(res, ShouldEqual, true)
		So(res2, ShouldEqual, true)
		So(res3, ShouldEqual, true)
	})
}