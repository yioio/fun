package funString

import "strconv"

func Hex2bin(str string) string {

	i, err := strconv.ParseInt(str, 16, 0)

	if err != nil {

		return ""
	}

	return strconv.FormatInt(i, 2)

}