package funFilesystem

import (
	."github.com/smartystreets/goconvey/convey"
	"testing"
	"github.com/yioio/fun/funFilesystem"
	"runtime"
	"path"
)

func Test_chmod(t *testing.T)  {
	Convey("修改文件权限模式", t, func() {
		_, currentFile, _, _ := runtime.Caller(0)
		dir := path.Dir(currentFile)
		path := dir + "/testFiles/ok.txt"

		res := funFilesystem.Chmod(path, 0777)

		So(res, ShouldEqual, true)
	})
}
