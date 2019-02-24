package funString

import(
	"crypto/md5"
	// "fmt"
	// "io"
	"encoding/hex"
)

func Md5(text string) string {

	hasher := md5.New()
	hasher.Write([]byte(text))
	
	return hex.EncodeToString(hasher.Sum(nil))
}