package funString

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func Md5_file(filePath string) string {
	file, _ := os.Open(filePath)
	defer file.Close()
	hash := md5.New()
	io.Copy(hash, file);
	return hex.EncodeToString(hash.Sum(nil))
}
