package funString

import "strings"

/**
 * 不建议外部使用 funString.StrReplace    strings.Replace 本身很简单
 */
func StrReplace(search, replace, subject string, count int) string {

	return strings.Replace(subject, search, replace, count)
}