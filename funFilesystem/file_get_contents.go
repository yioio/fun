package funFilesystem

import(
	"io/ioutil"
)

func File_get_contents(filename string) string {

	data, _ := File_get_contents_detail(filename)

	return string(data)
}

func File_get_contents_detail(filename string) (string, error) {

	data, err := ioutil.ReadFile(filename)

	return string(data), err
}