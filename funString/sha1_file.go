package funString

import (
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
)

//计算文件的 sha1 散列值
func Sha1_file (filename string) string{

	data, err := ioutil.ReadFile(filename)

	if err != nil {

		return ""
	}
	hash := sha1.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))

}