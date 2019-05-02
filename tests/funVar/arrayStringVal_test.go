package funVar

import(
    "testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funVar"
	"reflect"
)

func Test_ArrayStringVal(t *testing.T) {

	Convey("测试ArrayStringVal", t, func() {
        So(reflect.TypeOf(funVar.ArrayStringVal([]interface{}{"1", "abc"})), ShouldEqual, reflect.TypeOf([]string{""}))
        So(funVar.ArrayStringVal([]interface{}{"1", "abc"})[0], ShouldEqual, []string{"1", "abc"}[0])
        So(funVar.ArrayStringVal([]interface{}{"1", "abc"})[1], ShouldEqual, []string{"1", "abc"}[1])
	})
}
