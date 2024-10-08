package main

import (
	"fmt"
	"tetris-optimizer/functions"
)

func main() {

	tetrominos := functions.ReadFile()
	fmt.Println(tetrominos)

}
