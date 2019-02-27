package funFilesystem

import (
	."github.com/smartystreets/goconvey/convey"
	"testing"
	"github.com/yioio/fun/funFilesystem"
	"runtime"
	"path"
)

func Test_Copy(t *testing.T)  {
	Convey("拷贝文件", t, func() {
		_, currentFile, _, _ := runtime.Caller(0)
		dir := path.Dir(currentFile)
		src := dir + "/testFiles/ok.txt"
		dst := dir + "/testFiles/ok2.txt"

		res := funFilesystem.Copy(src, dst)
		res2 := funFilesystem.File_exits(dst)
		res3 := funFilesystem.Delete(dst)

		So(res, ShouldEqual, true)
		So(res2, ShouldEqual, true)
		So(res3, ShouldEqual, true)
	})
}