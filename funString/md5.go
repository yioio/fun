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

	// w := md5.New()
	// io.WriteString(w, str)//将str写入到w中
	// md5str2 := fmt.Sprintf("%x", w.Sum(nil))  //w.Sum(nil)将w的hash转成[]byte格式

	// return mdtstr2
}

// func Md5_16(str string, raw_output bool) string {
// 	data := []byte(str)
// 	has := md5.Sum(data)
// 	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制

// 	fmt.Println(md5str1)

// 	// hasher := md5.New()
// 	// hasher.Write([]byte(text))
// 	// return hex.EncodeToString(hasher.Sum(nil))
// }