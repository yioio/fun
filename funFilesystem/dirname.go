package funFilesystem

import (
	"path/filepath"
)

/**
  不建议外部使用 直接使用filepath.Dir即可
 */
func Dirname(path string) string {
	result := filepath.Dir(path)

	return result
}
