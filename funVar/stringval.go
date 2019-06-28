package funVar

import(
	"fmt"
	"reflect"
)

func Stringval(value interface{}) string {

	valueString := ""
	switch value.(type) {
		case []byte: {
			valueString = fmt.Sprintf("%s", (string(value.([]byte))))
		}
		default:
			valueString := fmt.Sprintf("%v", value)
			return valueString
	}
	return valueString
}

/**
	下划线字符串转成驼峰字符串
	如 abcd_a ==> abcdA
 */
func UnderlineStringToCamel(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			if i > 0 {
				d = d - 32
			}
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

/**
	驼峰字符串转成下划线+小些字母
	例如： 	abcdAccc ==> abcd_accc
			AbcdA    ==> abcd_a
	如果遇到有大写字母的地方，则将在前一位将加上下划线 "_", 将大写字母转成小字母（如果是首字母，则只是转成小谢即可）
 */
func CamelStringToUnderline(s string) string {
	returnString := ""
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if d >= 'A' && d <= 'Z' {
			// 如果遇到有大写字母的地方，则将在前一位将加上下划线 "_", 将大写字母转成小字母（如果是首字母，则只是转成小谢即可）
			ordNum := d + 32
			if i == 0 {
				returnString = returnString + "" + Stringval(string(ordNum))
			} else {
				returnString = returnString + "_" + Stringval(string(ordNum))
			}
		} else {
			returnString = returnString + Stringval(string(d))
		}
	}
	return returnString
}

/**
	将map下的所有键  由下划线转成驼峰
 */
func MapUnderlineStringToCamel(body map[string]interface{}) map[string]interface{} {
	keys := reflect.ValueOf(body).MapKeys()
	newMap := make(map[string]interface{})
	for _, v := range keys {
		strKey := v.Interface().(string)
		newMap[UnderlineStringToCamel(strKey)] = body[strKey]
	}
	return newMap
}

/**
	将map下的所有键  由驼峰转成下划线
 */
func MapCamelStringToUnderline(body map[string]interface{}) map[string]interface{} {
	keys := reflect.ValueOf(body).MapKeys()
	newMap := make(map[string]interface{})
	for _, v := range keys {
		strKey := v.Interface().(string)
		newMap[CamelStringToUnderline(strKey)] = body[strKey]
	}
	return newMap
}
