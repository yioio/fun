package funString

import "unicode"

func Lcfirst(str string) string {

	for _, v := range str {
		u := string(unicode.ToLower(v))
		return u + str[len(u):]
	}
	return ""
}