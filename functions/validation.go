package functions

func Validate(tetrominos [][]string) bool {
	if len(tetrominos) > 26 {
		return false
	}
	for _, eachtetro := range tetrominos {
		if len(eachtetro) != 4 {
			return false
		}
		for _, str := range eachtetro {
			if len(str) != 4 {
				return false
			}
		}
		if !countConnection(eachtetro) {
			return false
		}

	}
	return true
}

func countConnection(eachTetro []string) bool {
	countC := 0
	countConnect := 0
	for i, str := range eachTetro {
		for j, c := range str {
			if c != '.' {
				countC++
				if i > 0 && eachTetro[i-1][j] == byte(c) {
					countConnect++
				}
				if i < len(eachTetro)-1 && eachTetro[i+1][j] == byte(c) {
					countConnect++
				}
				if j > 0 && eachTetro[i][j-1] == byte(c) {
					countConnect++
				}
				if j < len(str)-1 && eachTetro[i][j+1] == (byte(c)) {
					countConnect++
				}
			}
		}
	}
	// if countConnect < 6 || countC != 4 {
	// 	return false
	// } else {
	// 	return true
	// }
	return countC == 4 && (countConnect == 6 || countConnect == 8)
}
