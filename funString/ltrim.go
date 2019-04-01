package funString

import "strings"

/**
 * 不建议外部使用 funString.Ltrim    strings.TrimLeft 本身很简单
 */
func Ltrim(str string, characterMask ...string) string {

	mask := ""
	if len(characterMask) == 0 {

		mask = " \\t\\n\\r\\0\\x0B"

	} else {
		mask = characterMask[0]
	}
	return strings.TrimLeft(str, mask)
}
