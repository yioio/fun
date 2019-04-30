package funString

import "html"

// 将字符转换为 HTML 转义字符
func Htmlentities (str string) string {

	return html.EscapeString(str)
}