package funArray

//import (
//	"log"
//	)

func Array_fill(start_index int, num uint, value interface{}) map[interface{}]interface{} {
	fillArr := map[interface{}]interface{}{}
	var i uint
	for i = 1; i <= num; i++ {
		//log.Println(i)
		fillArr[start_index] = value
		start_index++
	}

	return fillArr
}
