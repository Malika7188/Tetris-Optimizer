package main

import (
	"fmt"
	"math"
	"tetris-optimizer/functions"
)

func main() {

	tetrominos := functions.ReadFile()
	if !functions.Validate(tetrominos) {
		fmt.Println("ERROR")
		return
	}
	tetrominos = functions.TrimTetris(tetrominos)
	// fmt.Println(tetrominos)

	size := int(math.Ceil(math.Sqrt(float64(len(tetrominos)) * 4)))
	var finalBoard [][]string
	for {
		board := Board(size)
		finalBoard = functions.Solve(board, tetrominos)
		if finalBoard != nil {
			break
		}
		size++
	}
	for _, boad := range finalBoard {
		for _, c := range boad {
			fmt.Print(c)
		}
		fmt.Println()
	}

}
func Board(size int) [][]string {
	boad := make([][]string, size)
	for i := range boad {
		boad[i] = make([]string, size)
		for j := range boad[i] {
			boad[i][j] = "."
		}
	}
	return boad
}
