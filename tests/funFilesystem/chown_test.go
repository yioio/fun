package funFilesystem

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
	"github.com/yioio/fun/funFilesystem"
	"runtime"
	"path"
)

func Test_Chown(t *testing.T)  {
	Convey("修改文件所有者", t, func() {
		_, currentFile, _, _ := runtime.Caller(0)
		dir := path.Dir(currentFile)
		path := dir + "/testFiles/ok.txt"

		guid := os.Getuid()
		uid := os.Getuid()
		res := funFilesystem.Chown(path, guid, uid)

		So(res, ShouldEqual, true)
	})
}
