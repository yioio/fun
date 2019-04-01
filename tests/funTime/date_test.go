package funTime

import (
    "testing" 
	. "github.com/smartystreets/goconvey/convey" 
	"github.com/yioio/fun/funTime"
	"github.com/yioio/fun/funVar"
	"strings"
	"os/exec"
	// "fmt"
	"time"
)

func Test_Date(t *testing.T) {

	Convey("Test Date", t, func() {

		cmd := exec.Command("date", "+%s")
		out, _ := cmd.Output()
		newStr:= strings.Trim(string(out), "\n")
		shelltime := funVar.Int64val(newStr)

		//时区正确下的比较
		timestamp := funTime.Strtotime_Date("2017-05-11")
		date := funTime.Date(timestamp)
		So(date, ShouldEqual, "2017-05-11")

		timestamp = time.Now().Unix()
		date = funTime.Date(timestamp)
		today := funTime.Date_today()
		tomorrow := funTime.Date_tomorrow()
		So(date, ShouldEqual, funTime.Date(shelltime))
		So(funTime.Date_now() , ShouldEqual, funTime.Date_today())
		
		So(timestamp - funTime.Strtotime_Date(today) < 86400, ShouldBeTrue)
		So(shelltime - timestamp < 86400, ShouldBeTrue)
		So(funTime.Strtotime_Date(tomorrow) - 86400 <= shelltime, ShouldBeTrue)
		So(tomorrow, ShouldEqual, funTime.Date_daysafter(1, timestamp))
		So(today, ShouldEqual, funTime.Date_daysafter(0, timestamp))
		So(today, ShouldEqual, funTime.Date_daysbefore(1, timestamp + 86400))
	})
}
