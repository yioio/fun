package funString

//将字符串分割成小块

func Chunk_split(body string, chunklen uint, end string) string{

	if end == "" {
		end = "\r\n"
	}
	runes, erunes := []rune(body), []rune(end)

	l := uint(len(runes))
	if l <= 1 || l < chunklen {
		return body + end
	}
	ns := make([]rune, 0, len(runes)+len(erunes))
	var i uint
	for i = 0; i < l; i += chunklen {
		if i+chunklen > l {
			ns = append(ns, runes[i:]...)
		} else {
			ns = append(ns, runes[i:i+chunklen]...)
		}
		ns = append(ns, erunes...)
	}


	return string(ns)

}
