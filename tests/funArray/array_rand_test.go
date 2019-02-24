package funArray

import (
	"testing"
	"log"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funArray"
)

//
func Test_Array_rand(t *testing.T) {
	Convey("检查： Array_rand", t, func() {
		orgArr := map[interface{}]interface{}{"test":"test",0:"0", "22":0, "00":"89", "220":"123", "te":"888"}
		randElement := funArray.ArrayRand(orgArr, 5)   // []
		log.Println("Test_Array_rand1 return is ", randElement)
		// 获取到的第一个随机元素
		Frist := orgArr[randElement[0]]
		log.Println("Test_Array_rand1 first is ", Frist)
	})
}
