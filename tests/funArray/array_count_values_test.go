package funArray

import (
	"testing"
	"log"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funArray"
)

//
func Test_Array_count_values(t *testing.T) {
	Convey("检查： Test_Array_count_values", t, func() {
		//t1_arrKeys := []string{"1", "2", "3","4"}
		t1_arr1 :=[]interface{}{"test1","test1",0,0,0,0,"test1", "test2", "test2"}
		t1_res := funArray.Array_count_values(t1_arr1);
		log.Println("Test_Array_count_values return", t1_res)
		So(t1_res, ShouldEqual, t1_res)
	})
}
