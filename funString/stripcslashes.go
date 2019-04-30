package funString

import (
	"bytes"
)

//反引用一个使用 addcslashes() 转义的字符串
func Stripslashes(str string) string {

	var buf bytes.Buffer

	l, skip := len(str), false

	for i, char := range str {

		if skip {

			skip = false

			continue

		} else if char == '\\' {

			if i+1 < l && str[i+1] == '\\' {

				skip = true

			}

			continue
		}

		buf.WriteRune(char)
	}

	return buf.String()
}