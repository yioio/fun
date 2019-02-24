package funFilesystem

import (
	"path/filepath"
	"strings"
)

/**
 * basename ( string $path [, string $suffix ] ) : string
 */
func Basename(path, suffix string) string {
	return strings.TrimSuffix(filepath.Base(path), suffix)
}