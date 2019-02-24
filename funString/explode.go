package funString

import "strings"


/**
 * 不建议外部使用 funString.Explode    strings.Split 本身很简单
 */
func Explode(delimiter, str string) []string {

	return strings.Split(str, delimiter)
}
