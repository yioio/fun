package funString

import "strings"


/**
 * 不建议外部使用 funString.Trim    strings.Trim 本身很简单
 */
func Trim(str string, characterMask ...string) string {
	mask := ""
	if len(characterMask) == 0 {

		mask = " \\t\\n\\r\\0\\x0B"

	} else {
		mask = characterMask[0]
	}
	return strings.Trim(str, mask)
}