package funArray

import (
	"testing"
	"log"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funArray"
)

//
func Test_Array_intersect(t *testing.T) {

	Convey("检查： Array_intersect", t, func() {
		t1_arr1 :=[]interface{}{"test1", "test2"}
		t1_arr2 := []interface{}{"test2"}
		processRet1 := funArray.Array_intersect(t1_arr1, t1_arr2)
		log.Println("Array_intersect 1 return is ", processRet1)  // test1
		t2_arr1 :=[]interface{}{"test1", "test2"}
		t2_arr2 := []interface{}{"test1", "test2"}
		processRet2 := funArray.Array_intersect(t2_arr1, t2_arr2)   // []
		log.Println("Array_intersect 2 return is ", processRet2)
		//So(processRet1, ShouldEqual, processRet1)
		//So(processRet2, ShouldEqual, processRet2)
	})
}
