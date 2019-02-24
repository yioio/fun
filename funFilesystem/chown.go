package funFilesystem

import (
	"os"
)

/**
  不建议外部使用  直接使用os.Chown 即可
   window 下不可用
   On Windows, it always returns the syscall.EWINDOWS error, wrapped in *PathError.
 */
func Chown(fileName string, uid int, gid int) bool {
	result := os.Chown(fileName, uid, gid)

	return result == nil
}
