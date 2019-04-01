package funVar

import(
	"fmt"
	"strconv"
)

func Float64val(value interface{}) float64 {

	valueString := fmt.Sprintf("%v", value)
	valueFloat,err := strconv.ParseFloat(valueString, 64)

    if err != nil {
		valueFloat = 0.0
	}
	
	return valueFloat
}

// func Float32val(value interface{}) float32 {
// 	valueString := fmt.Sprintf("%v", value)
// 	valueFloat,err := strconv.ParseFloat(valueString, 32)
//     if err != nil {
// 		valueFloat,_ = strconv.ParseFloat("0", 32)
// 	}
// 	return valueFloat
// }
//