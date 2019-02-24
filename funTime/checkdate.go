package funTime

func Checkdata(m int, d int, y int) bool{
	if m <= 0 || d <= 0 || y <= 0 {
		return false
	}
	if m>12 {
		return false
	}
	var isLeapYear bool = false
	if (y%4 == 0 && y%100 != 0) || y%400 == 0{
		isLeapYear = true
	}

	if m == 2 {
		if isLeapYear {
			if d > 29 {
				return false
			}
		}else{
			if d > 28 {
				return false
			}
		}
	}else{
		if m == 1 || m == 3 || m == 5 || m == 7 || m == 8 || m == 10 || m == 12{
			if d > 31 {
				return false
			}
		}else{
			if d > 30 {
				return false
			}
		}
	}
	return true
}