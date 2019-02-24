package funArray
//
//import(
//	"reflect"
//	"sort"
//)

import (
    "testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funArray"
    "log"
    "reflect"
)

func Test_Array_combine(t *testing.T) {
	
	Convey("检查： Array_combine", t, func() {
		// 1.数量不等
		t1_arrKeys := []string{"1", "2", "3","4"}
		T1_arrValues :=[]interface{}{"test", t1_arrKeys, []int{123} }
		processRet1 , isOk1 := funArray.Array_combine(t1_arrKeys, T1_arrValues)
		log.Println("Test_Array_combine  1 return ", isOk1, processRet1)
		// 2.数量不等
		t2_arrKeys := []string{"1", "2", "3","ok"}
		T2_arrValues :=[]interface{}{"test", t2_arrKeys, []int{123}, "ok" }
		processRet2 , isOk2 := funArray.Array_combine(t2_arrKeys, T2_arrValues)
		log.Println("Test_Array_combine  2 return ", processRet2 , isOk2)
		log.Println("Test_Array_combine   return eme", processRet2["ok"])
		log.Println("Test_Array_combine   return eme Type", reflect.TypeOf(processRet2["ok"]))
        //So(processRet[0], ShouldEqual, ret[0])
        //So(processRet[1], ShouldEqual, ret[1])
	})
}
