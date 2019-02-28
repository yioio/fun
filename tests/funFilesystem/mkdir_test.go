
package funFilesystem

import (
	."github.com/smartystreets/goconvey/convey"
	"testing"
	"github.com/yioio/fun/funFilesystem"
	"runtime"
	"path"
)

func Test_mkdir(t *testing.T)  {
	Convey("创建目录", t, func() {
		_, currentFile, _, _ := runtime.Caller(0)
		dir := path.Dir(currentFile)
		newDir := dir + "/testFiles/test/test"

		exists := funFilesystem.File_exits(newDir)
		if exists == true {
			funFilesystem.Delete(dir)
		}

		res := funFilesystem.Mkdir(newDir, 755, true)
		res2 := funFilesystem.File_exits(newDir)

		So(res, ShouldEqual, true)
		So(res2, ShouldEqual, true)
	})
}
