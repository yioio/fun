package funVar

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funVar"
	"log"
	"testing"
)

func Test_MapCamelStringToUnderline(t *testing.T) {

	Convey("check： MapCamelStringToUnderline", t, func() {
		underlineMapString := map[string]interface{}{
			"testTest":"test",
			"Test_ok":"_test_ok",
		}
		processRet2 := funVar.MapCamelStringToUnderline(underlineMapString)
		log.Println("org data is  ", underlineMapString)
		log.Println("MapCamelStringToUnderline is ", processRet2)
	})
}
