package funTime

import (
    "testing"
	. "github.com/smartystreets/goconvey/convey"
	"go.increase/fun/funTime"
	"go.increase/fun/funVar"
	// "reflect"
	"os/exec"
	"strings"
)

func Test_Time(t *testing.T) {
	
	Convey("Test Time", t, func() {
		
		cmd := exec.Command("date", "+%s")
		out, _ := cmd.Output()
		newStr:= strings.Trim(string(out), "\n")
		shelltime := funVar.Int64val(newStr)
		
		timestamp := funTime.Time()

		So(shelltime > 0, ShouldBeTrue)
		So(timestamp > 0, ShouldBeTrue)
		So(timestamp - shelltime < 10, ShouldBeTrue)
		
        // So(fmt.Sprintf("%v", out), ShouldEqual, fmt.Sprintf("%v", timestamp), ShouldBeTrue)
	})
}
