package funString

import (
	"strings"
)

//二进制安全比较字符串
func Substr_compare(main_str , str string, offset, length int) int{


	mainLength := len(main_str)
	strLength := len(str)

	if length<0 || offset>= mainLength {

		return -2
	}


	if offset <0 {


		start1 := offset+mainLength
		end1 := offset+mainLength+length

		start2 := offset+strLength
		end2 := offset+strLength+length

		if end1 >= mainLength {

			main_str = main_str[start1:mainLength-1]

		} else {

			main_str = main_str[start1:end1]
		}

		if end2 >= strLength {

			str = main_str[start2:strLength-1]

		} else {

			str = main_str[start2:end2]
		}


	} else {

		if offset+length >= mainLength {

			main_str = main_str[offset:]


		} else {

			main_str = main_str[offset:offset+length]

		}

		if offset+length > strLength {

			if offset == 0 {

				str = str[offset:]

			} else {

				str = str[offset-1:]
			}


		} else {

			if offset == 0 {

				str = str[offset:offset+length]

			} else {

				str = str[offset-1:offset+length-1]
			}

		}

	}

	return strings.Compare(main_str, str)
}
