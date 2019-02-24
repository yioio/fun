package funTime

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funTime"
	"testing"
)

func Test_Checkdate(t *testing.T) {

	Convey("Test Checkdata", t, func() {
		So(funTime.Checkdata(12, 31, 2000), ShouldEqual, true)
		So(funTime.Checkdata(2, 29, 2001), ShouldEqual, false)
	})
}


