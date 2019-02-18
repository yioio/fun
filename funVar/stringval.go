package funVar

import(
	"fmt"
)

func Stringval(value interface{}) string {

	valueString := ""
	switch value.(type) {
		case []byte: {
			valueString := fmt.Sprintf("%s", (string(value.([]byte))))
			return valueString
		}
		case string:
			return valueString
		default:			
			valueString := fmt.Sprintf("%v", value)
			return valueString
	}

	return valueString
}
