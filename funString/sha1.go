package funString

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha1(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}