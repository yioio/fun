package funFilesystem

import(
	"os"
)

func File_exits(path string) bool {

	ret, _ := File_exits_detail(path)

	return ret
}

func File_exits_detail(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
