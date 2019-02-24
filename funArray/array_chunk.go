package funArray

import(
	"strings"

)

type mapArray map[string]interface{}

const (
	CaseLOWER = "lower"
	CaseUPPER = "upper"
)
/**
 * mapArray keys case to lower or upper
 */
func Array_change_key_case(mapArray map[string]interface{}, caseString string) map[string]interface{} {
	var caseArr =  mapArray
	if caseString == CaseUPPER {
		caseArr = Array_change_key_case_upper(mapArray)
	} else if caseString == CaseLOWER {
		caseArr = Array_change_key_case_upper(mapArray)
	}
	return caseArr
}

/**
 * CASE_UPPER
 */
func Array_change_key_case_upper(arr mapArray) mapArray {
	var caseMapArray = mapArray{}
	for k, v := range arr {
		caseMapArray[strings.ToUpper(k)] = v
	}
	return  caseMapArray
}

/**
 * CASE_LOWER
 */
func Array_change_key_case_lower(arr mapArray) mapArray {
	var caseMapArray = mapArray{}
	for k, v := range arr {
		caseMapArray[strings.ToLower(k)] = v
	}
	return  caseMapArray
}

