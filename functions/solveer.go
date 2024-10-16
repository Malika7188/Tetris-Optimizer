package functions

func Solve(board [][]string, tetrominos [][]string) [][]string {
	if SolveRecursive(board, tetrominos, 0) {
		return board
	}
	return nil
}
func SolveRecursive(board, tetrominos [][]string, index int) bool {
	if index == len(tetrominos) {
		return true
	}
	eachTetro := tetrominos[index]
	for i := range board {
		for j := range board[i] {
			if canPlace(board, eachTetro, i, j) {
				Place(board, eachTetro, i, j)
				if SolveRecursive(board, tetrominos, index+1) {
					return true
				}
				remove(board, eachTetro, i, j)
			}
		}
	}
	return false
}

func canPlace(board [][]string, eachTetro []string, i, j int) bool {
	for distI := range eachTetro {
		for distJ, c := range eachTetro[distI] {
			if c != '.' {
				if i+distI > len(board)-1 || j+distJ > len(board[i])-1 || board[i+distI][j+distJ] != "." {
					return false
				}
			}
		}
	}
	return true
}
func Place(board [][]string, eachTetro []string, i, j int) {
	for distI := range eachTetro {
		for distJ, c := range eachTetro[distI] {
			if c != '.' {
				board[i+distI][j+distJ] = string(c)
			}
		}
	}
}
func remove(board [][]string, eachTetro []string, i, j int) {
	for distI := range eachTetro {
		for distJ, c := range eachTetro[distI] {
			if c != '.' {
				board[i+distI][j+distJ] = "."
			}
		}
	}
}
