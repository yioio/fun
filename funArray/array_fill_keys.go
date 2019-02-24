package funArray

import(
	"errors"
)

func Array_combine(arrayKeys [] string, arrayValues []interface{}) (map[string]interface{}, error)  {
	combineArr := make(map[string]interface{})
	KeysLen := len(arrayKeys)
	valuesLen := len(arrayValues)
	if KeysLen != valuesLen {
		return combineArr, errors.New("键和值的数量不不一致")
	}
	for i := 0; i < KeysLen; i++ {
		combineArr[arrayKeys[i]] = arrayValues[i]
	}
	return combineArr, nil
}
