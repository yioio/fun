package funString

import "strconv"

func Bin2hex(str string) string {

	i, err := strconv.ParseInt(str, 2, 0)

	if err != nil {

		return ""
	}

	return strconv.FormatInt(i, 16)

}