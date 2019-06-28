package funVar

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funVar"
	"log"
	"testing"
)

func Test_MapUnderlineStringToCamel(t *testing.T) {

	Convey("check： MapUnderlineStringToCamel", t, func() {
		underlineMapString := map[string]interface{}{
			"test_test":"test",
			"_test_ok":"_test_ok",
		}
		processRet2 := funVar.MapUnderlineStringToCamel(underlineMapString)
		log.Println("org data is ", underlineMapString)
		log.Println("MapUnderlineStringToCamel is  ", processRet2)
	})
}
