package functions

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile() [][]string {
	if len(os.Args) != 2 {
		fmt.Println("Error")
		os.Exit(0)
	}

	args := os.Args[1]

	filename, err := os.ReadFile(args)
	if err != nil {
		fmt.Println("Error")
		os.Exit(0)
	}
	tetrominos := [][]string{}
	eachTetro := []string{}
	var letter rune = 'A'

	file := strings.Split(string(filename), "\n")

	for _, str := range file {
		if str != "" {
			s := ""
			for _, c := range str {
				if c == '#' {
					s += string(letter)
				} else if c == '.' {
					s += "."
				} else {
					fmt.Println("Error")
					os.Exit(0)
				}
			}
			eachTetro = append(eachTetro, s)
		} else {
			tetrominos = append(tetrominos, eachTetro)
			eachTetro = []string{}
			letter++
		}
	}
	return tetrominos
}
