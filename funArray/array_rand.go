package funArray

import (
"log"
"math/rand"
"reflect"
"time"
)
/**
  *
  * array_rand — 从数组中随机取出一个或多个单元, 并返回数组的键
  * @params unit start_index  返回的数组的第一个索引值
  * @params int num 随机获取的数量。 必须大于或等于0
  *
  * 1.获取key 的个数
  * 2.随机获取key
  *
  * @return  map[int]interface{} 键的数据
  */
func ArrayRand(elements map[interface{}]interface{}, num int) map[interface{}]interface{} {
	returnKeys := make(map[interface{}]interface{})
	keys := Array_keys_interface(elements)
	keysLen := len(keys)
	log.Println("rand", keysLen, reflect.TypeOf(keysLen))
	// 随机获取
	var i  = 0
	for {
		if len(returnKeys) >= num {
			break
		}
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		number := r.Intn(keysLen)
		log.Println("randaaa", number)
		if Array_key_exists(number, returnKeys) == true {
			continue
		} else {
			returnKeys[i] = keys[number]
			i++
		}
	}
	return returnKeys
}

//func ArrayRand(elements []interface{}) []interface{} {
//	r := rand.New(rand.NewSource(time.Now().UnixNano()))
//	n := make([]interface{}, len(elements))
//	for i, v := range r.Perm(len(elements)) {
//		n[i] = elements[v]
//	}
//	return n
//}
