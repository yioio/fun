package funFilesystem

import (
	"io"
	"os"
)

func Copy(source string, dest string) bool  {
	src, err := os.Open(source)
	if err != nil {
		return false
	}
	// 使用defer不管后面的代码流程如何影响，这个文件能够被自动关闭
	defer src.Close()
	dst, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return false
	}
	// 使用defer不管后面的代码流程如何影响，这个文件能够被自动关闭
	defer dst.Close()

	_, error := io.Copy(dst, src)

	return  error == nil
}
