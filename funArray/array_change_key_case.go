package funArray

import(
	"strings"
	// "go.increase/fun/funString"
)

/**
 * 将数组中的所有键名修改为全大写或小写
 * 准确来讲，数组的key是数字，且下标从0开始，只有map才有改变key的需求
 * changeTo 可取值为 funArray.CASE_UPPER | funArray.CASE_LOWER
 */
func Array_change_key_case(arrayOrigin map[string]interface{}, changeTo int) map[string]interface{} {
	
	arrayReturn := make(map[string]interface{})
	for key, value := range arrayOrigin {

		keyNew := ""
		
		if (changeTo == CASE_UPPER) {
			keyNew = strings.ToUpper(key)
		} else if(changeTo == CASE_LOWER) {		
			keyNew = strings.ToLower(key)
		} else {
			keyNew = key
		}

		arrayReturn[keyNew] = value.(interface{})
	}

	return arrayReturn
}
