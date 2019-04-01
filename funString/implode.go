package funString

import "strings"

/**
 * 不建议外部使用 funString.Implode    strings.Join 本身很简单
 */
func Implode(glue string, pieces []string) string {

	return strings.Join(pieces, glue)
}
