package funArray

import(
	"github.com/yioio/fun/funVar"
	"math"
)
// ArraySlice type
type ArraySlice []interface{}

/**
 *  @param []interface{} inputArr  The array to work on
 *  @param int  size The size of each chunk
 */
func Array_chunk(inputArr ArraySlice, size int ) ArraySlice {
	length := len(inputArr)
	count := int(math.Ceil(funVar.Float64val(length) / funVar.Float64val(size)))
	chunckArr := make(ArraySlice, count)
	for i := 0; i < count-1; i++ {
		chunckArr[i] = inputArr[i*size : (i+1)*size]
	}
	chunckArr[count-1] = inputArr[size*(count-1):]

	return chunckArr
}

