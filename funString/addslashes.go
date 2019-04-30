package funString

import "bytes"

//使用反斜线引用字符串
func Addslashes(str string) string {

	var buf bytes.Buffer

	for _, char := range str {

		switch char {

			case '\'', '"', '\\':

				buf.WriteRune('\\')
		}

		buf.WriteRune(char)
	}

	return buf.String()

}