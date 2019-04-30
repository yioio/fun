package funString

import (
	"bytes"
)

//在字符串所有新行之前插入 HTML 换行标记

func Nl2br(str string, isXhtml bool) string {

	r, n, runes := '\r', '\n', []rune(str)

	var br []byte

	if isXhtml {

		br = []byte("<br />")

	} else {

		br = []byte("<br>")
	}

	skip := false

	length := len(runes)

	var buf bytes.Buffer

	for i, v := range runes {


		if skip {
			skip = false
			continue
		}

		switch v {

			case n, r:

				if (i+1 < length) && (v == r && runes[i+1] == n) || (v == n && runes[i+1] == r) {
					buf.Write(br)
					skip = true
					continue
				}

				buf.Write(br)

			default:

				buf.WriteRune(v)
		}
	}

	return buf.String()

}