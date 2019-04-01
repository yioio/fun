package funFilesystem

import (
	"log"
	"os"
	"time"
)

//获取文件修改时间 返回unix时间戳
func Filemtime(path string) int64 {
	f, err := os.Open(path)
	if err != nil {
		log.Println("open file error")
		return time.Now().Unix()
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		log.Println("stat fileinfo error")
		return time.Now().Unix()
	}

	return fi.ModTime().Unix()
}
