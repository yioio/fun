package funFilesystem

import (
    "testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funFilesystem"
)

func Test_File_get_contents(t *testing.T) {
	
	Convey("检查获取 文件内容", t, func() {
		ret1 := funFilesystem.File_get_contents("/go/src/tuan_access/tools/initRouters.goxxxxxxxxx")
		ret2 := funFilesystem.File_get_contents("file_get_contents_test.go")
		
        So(ret1 == "", ShouldEqual, true)
        So(ret2[0:7], ShouldEqual, "package")
        So(ret2[0:21], ShouldEqual, "package funFilesystem")
	})
}
