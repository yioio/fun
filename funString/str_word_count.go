package funString

import "strings"

//返回字符串中单词的使用情况
func Str_word_count(str string) []string {

	return strings.Fields(str)
}
