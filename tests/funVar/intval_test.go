package funVar

import(
    "testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funVar"
	"reflect"
)

func Test_Intval(t *testing.T) {

	Convey("测试Intval传入值为空", t, func() {
        So(reflect.TypeOf(funVar.Intval(string("0"))), ShouldEqual, reflect.TypeOf(int(0)))
        So(reflect.TypeOf(funVar.Intval(string(""))), ShouldEqual, reflect.TypeOf(int(0)))
        So(funVar.Intval(string("")), ShouldEqual, int(0))
        So(funVar.Intval(string("1")), ShouldEqual, int(1))
        So(funVar.Intval(string("-1")), ShouldEqual, int(-1))
        So(funVar.Intval(string("-1.1")), ShouldEqual, int(-1))
        So(funVar.Intval(int(22)), ShouldEqual, int(22))
        So(funVar.Intval(int8(6)), ShouldEqual, int(6))
        So(funVar.Intval(int32(33)), ShouldEqual, int(33))
        So(funVar.Intval(int64(33)), ShouldEqual, int(33))
        So(funVar.Intval(float64(33.33)), ShouldEqual, int(33)) 
        So(funVar.Intval(float64(0.01)), ShouldEqual, int(0))
        So(funVar.Intval(float64(-0.01)), ShouldEqual, int(0))
	})
}

func Test_Int32val(t *testing.T) {

	Convey("测试Int32val", t, func() {
        So(reflect.TypeOf(funVar.Int32val(string("0"))), ShouldEqual, reflect.TypeOf(int32(0)))
        So(reflect.TypeOf(funVar.Int32val(string(""))), ShouldEqual, reflect.TypeOf(int32(0)))
        So(funVar.Int32val(string("")), ShouldEqual, int32(0))
        So(funVar.Int32val(string("1")), ShouldEqual, int32(1))
        So(funVar.Int32val(int(22)), ShouldEqual, int32(22))
        So(funVar.Int32val(int8(6)), ShouldEqual, int32(6))
        So(funVar.Int32val(int32(33)), ShouldEqual, int32(33))
        So(funVar.Int32val(int64(33)), ShouldEqual, int32(33))
        So(funVar.Int32val(float64(33.33)), ShouldEqual, int32(33))
	})
}

func Test_Int64val(t *testing.T) {

	Convey("测试Int64val", t, func() {
        So(reflect.TypeOf(funVar.Int64val(string("0"))), ShouldEqual, reflect.TypeOf(int64(0)))
        So(reflect.TypeOf(funVar.Int64val(string(""))), ShouldEqual, reflect.TypeOf(int64(0)))
        So(funVar.Int64val(string("")), ShouldEqual, int64(0))
        So(funVar.Int64val(string("1")), ShouldEqual, int64(1))
        So(funVar.Int64val(string("-1.1")), ShouldEqual, int64(-1))
        So(funVar.Int64val(int(22)), ShouldEqual, int64(22))
        So(funVar.Int64val(int8(6)), ShouldEqual, int64(6))
        So(funVar.Int64val(int32(33)), ShouldEqual, int64(33))
        So(funVar.Int64val(int64(33)), ShouldEqual, int64(33))
        So(funVar.Int64val(float64(33.33)), ShouldEqual, int64(33))
        So(funVar.Int64val("1548857206"), ShouldEqual, int64(1548857206))
        So(funVar.Int64val(1548857206), ShouldEqual, int64(1548857206))
	})
}
 