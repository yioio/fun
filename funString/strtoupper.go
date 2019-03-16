package funString

import(
	"strings"
)

/**
 * 不建议外部使用 funString.Strtoupper，因为 strings.ToUpper 本身很简单
 */
func Strtoupper(origin string) string {
	return strings.ToUpper(origin)
}
