package funFilesystem

import (
	"os"
)

func Mkdir(path string, model os.FileMode, recursive bool) bool {
	var err error
	if recursive == true {
		err = os.MkdirAll(path, model)
	} else {
		err = os.Mkdir(path, model)
	}
	return err == nil
}