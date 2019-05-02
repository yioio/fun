package funVar

import(
	"fmt"
)

func ArrayStringVal(origin []interface{}) []string {

	newRet := []string{}
	for _, value := range origin {
		newRet = append(newRet, fmt.Sprintf("%v", value))
	}

	return newRet
}