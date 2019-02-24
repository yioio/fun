package funString

import "strings"


/**
 * 不建议外部使用 funString.Ucwords    strings.Title 本身很简单
 */
func Ucwords(str string) string {

	return strings.Title(str)
}
