package funString

import(
	"strings"
)

/**
 * 不建议外部使用 funString.Strtolower strings.ToUpper 本身很简单
 */
func Strtolower(origin string) string {
	return strings.ToLower(origin)
}