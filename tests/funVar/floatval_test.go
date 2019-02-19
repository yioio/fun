package funVar

import(
    "testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funVar"
	"reflect"
)

func Test_Float64val(t *testing.T) {

	Convey("测试Float64val传入值为空", t, func() {
        So(reflect.TypeOf(funVar.Float64val(string("0"))), ShouldEqual, reflect.TypeOf(float64(0)))
		So(reflect.TypeOf(funVar.Float64val(string(""))), ShouldEqual, reflect.TypeOf(float64(0)))
		
        So(funVar.Float64val(string("")), ShouldEqual, float64(0))
        So(funVar.Float64val(string("1")), ShouldEqual, float64(1))
        So(funVar.Float64val(string("1.1")), ShouldEqual, float64(1.1))
        So(funVar.Float64val(string("-1")), ShouldEqual, float64(-1))
		So(funVar.Float64val(string("-1.1")), ShouldEqual, float64(-1.1))

        So(funVar.Float64val(int(0)), ShouldEqual, float64(0))
        So(funVar.Float64val(int(1)), ShouldEqual, float64(1))
		So(funVar.Float64val(int(-1)), ShouldEqual, float64(-1))
		
        So(funVar.Float64val(float32(0.1)), ShouldEqual, float64(0.1))
        So(funVar.Float64val(float32(1.0)), ShouldEqual, float64(1.0))
        So(funVar.Float64val(float32(-1)), ShouldEqual, float64(-1))
        So(funVar.Float64val(float32(-1.1)), ShouldEqual, float64(-1.1))
        So(funVar.Float64val(float32(-1.101)), ShouldEqual, float64(-1.101))
		So(funVar.Float64val(float32(-1.1)), ShouldNotEqual, float64(-1.12))
		
        So(funVar.Float64val(float64(0.1)), ShouldEqual, float64(0.1))
        So(funVar.Float64val(float64(1.0)), ShouldEqual, float64(1.0))
        So(funVar.Float64val(float64(-1)), ShouldEqual, float64(-1))
        So(funVar.Float64val(float64(-1.1)), ShouldEqual, float64(-1.1))
        So(funVar.Float64val(float64(-1.101)), ShouldEqual, float64(-1.101))
        So(funVar.Float64val(float64(-1.1)), ShouldNotEqual, float64(-1.12))
		
	})

	// Convey("测试Float32val传入值为空", t, func() {
    //     So(reflect.TypeOf(funVar.Float32val(string("0"))), ShouldEqual, reflect.TypeOf(float32(0)))
	// 	So(reflect.TypeOf(funVar.Float32val(string(""))), ShouldEqual, reflect.TypeOf(float32(0)))
		
    //     So(funVar.Float32val(string("")), ShouldEqual, float32(0))
    //     So(funVar.Float32val(string("1")), ShouldEqual, float32(1))
    //     So(funVar.Float32val(string("1.1")), ShouldEqual, float32(1.1))
    //     So(funVar.Float32val(string("-1")), ShouldEqual, float32(-1))
	// 	So(funVar.Float32val(string("-1.1")), ShouldEqual, float32(-1.1))

    //     So(funVar.Float32val(int(0)), ShouldEqual, float32(0))
    //     So(funVar.Float32val(int(1)), ShouldEqual, float32(1))
	// 	So(funVar.Float32val(int(-1)), ShouldEqual, float32(-1))
		
    //     So(funVar.Float32val(float32(0.1)), ShouldEqual, float32(0.1))
    //     So(funVar.Float32val(float32(1.0)), ShouldEqual, float32(1.0))
    //     So(funVar.Float32val(float32(-1)), ShouldEqual, float32(-1))
    //     So(funVar.Float32val(float32(-1.1)), ShouldEqual, float32(-1.1))
    //     So(funVar.Float32val(float32(-1.101)), ShouldEqual, float32(-1.101))
	// 	So(funVar.Float32val(float32(-1.1)), ShouldNotEqual, float32(-1.12))
		
    //     So(funVar.Float32val(float64(0.1)), ShouldEqual, float32(0.1))
    //     So(funVar.Float32val(float64(1.0)), ShouldEqual, float32(1.0))
    //     So(funVar.Float32val(float64(-1)), ShouldEqual, float32(-1))
    //     So(funVar.Float32val(float64(-1.1)), ShouldEqual, float32(-1.1))
    //     So(funVar.Float32val(float64(-1.101)), ShouldEqual, float32(-1.101))
    //     So(funVar.Float32val(float64(-1.1)), ShouldNotEqual, float32(-1.12))
		
	// })
}