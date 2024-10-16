package functions

import (
	"strings"
)

func TrimTetris(tetrominos [][]string) [][]string {
	leters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	trimedTetris := [][]string{}
	containLetter := []string{}

	for _, eachTetro := range tetrominos {

		for _, str := range eachTetro {
			for _, l := range leters {
				if strings.Contains(str, string(l)) {
					containLetter = append(containLetter, str)
					break
				}

			}
		}
		trimedTetris = append(trimedTetris, containLetter)
		containLetter = []string{}
	}

	newTrimedtetro := [][]string{}

	for _, eachTetro := range trimedTetris {
		hasLetters := make([]bool, 4)
		for j := 0; j < 4; j++ {
			for i := 0; i < len(eachTetro); i++ {
				for _, c := range leters {
					if eachTetro[i][j] == byte(c) {
						hasLetters[j] = true
					}
				}
			}
		}
		newSlice := []string{}
		for _, str := range eachTetro {
			newStr := ""
			for i, c := range str {
				if hasLetters[i] {
					newStr += string(c)
				}
			}
			newSlice = append(newSlice, newStr)
		}
		newTrimedtetro = append(newTrimedtetro, newSlice)
	}
	return newTrimedtetro
}
