package funString

import (
	"bytes"
	"unicode"
)

//打断字符串为指定数量的字串
func Wordwrap(str string, width uint, br string) string {
	if br == "" {
		br = "\n"
	}
	init := make([]byte, 0, len(str))
	buf := bytes.NewBuffer(init)
	var current uint
	var wordbuf, spacebuf bytes.Buffer
	for _, char := range str {
		if char == '\n' {
			if wordbuf.Len() == 0 {
				if current+uint(spacebuf.Len()) > width {
					current = 0
				} else {
					current += uint(spacebuf.Len())
					spacebuf.WriteTo(buf)
				}
				spacebuf.Reset()
			} else {
				current += uint(spacebuf.Len() + wordbuf.Len())
				spacebuf.WriteTo(buf)
				spacebuf.Reset()
				wordbuf.WriteTo(buf)
				wordbuf.Reset()
			}
			buf.WriteRune(char)
			current = 0
		} else if unicode.IsSpace(char) {
			if spacebuf.Len() == 0 || wordbuf.Len() > 0 {
				current += uint(spacebuf.Len() + wordbuf.Len())
				spacebuf.WriteTo(buf)
				spacebuf.Reset()
				wordbuf.WriteTo(buf)
				wordbuf.Reset()
			}
			spacebuf.WriteRune(char)
		} else {
			wordbuf.WriteRune(char)
			if current+uint(spacebuf.Len()+wordbuf.Len()) > width && uint(wordbuf.Len()) < width {
				buf.WriteString(br)
				current = 0
				spacebuf.Reset()
			}
		}
	}

	if wordbuf.Len() == 0 {
		if current+uint(spacebuf.Len()) <= width {
			spacebuf.WriteTo(buf)
		}
	} else {
		spacebuf.WriteTo(buf)
		wordbuf.WriteTo(buf)
	}
	return buf.String()
}