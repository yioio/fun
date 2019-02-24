package funVar

import(
	"fmt"
	"strconv"
	"strings"
	// "math"
	// "reflect" 
)

func Intval(value interface{}) int {

	valueString := fmt.Sprintf("%v", value)
	valueString = strings.Split(valueString, ".")[0]

    i, err := strconv.Atoi(valueString)
    if err != nil {
		i, _ = strconv.Atoi(string("0"))
	}

	return i
}

func Int32val(value interface{}) int32 {
	int64Val := Int64val(value)

	return int32(int64Val)
}

func Int64val(value interface{}) int64 {

	valueString := fmt.Sprintf("%v", value)
	valueString = strings.Split(valueString, ".")[0]

	i, err := strconv.ParseInt(valueString, 10, 64)

    if err != nil {
		i, _ = strconv.ParseInt(string("0"), 10, 64)
	}

	return i
}
