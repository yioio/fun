package funFilesystem

import (
	"os"
)

/**
  不建议外部使用 直接使用 os.Chmod 即可
 */
func Chmod(fileName string, model os.FileMode) bool {
	result := os.Chmod(fileName, model)

	return  result == nil
}
