package funString

import "unicode/utf8"

//返回指定的字符
func Ord(str string) int {

	r, _ := utf8.DecodeRune([]byte(str))

	return int(r)
}