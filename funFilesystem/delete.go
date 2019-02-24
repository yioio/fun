package funFilesystem

import "os"

/**
 不建议外部使用 直接使用os.Remove即可
 */
func Delete(path string) bool {
	result := os.Remove(path)

	return result == nil
}
