package functions

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	board := [][]string{
		{".", ".", ".", "."},
		{".", ".", ".", "."},
		{".", ".", ".", "."},
		{".", ".", ".", "."},
	}

	tetrominos := [][]string{
		{"AA..", "AA..", "....", "...."},
		{".BB.", "BB..", "....", "...."},
	}

	expected := [][]string{
		{"A", "A", ".", "."},
		{"A", "A", "B", "B"},
		{".", "B", "B", "."},
		{".", ".", ".", "."},
	}

	result := Solve(board, tetrominos)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestSolveNoSolution(t *testing.T) {
	board := [][]string{
		{".", ".", "."},
		{".", ".", "."},
		{".", ".", "."},
	}

	tetrominos := [][]string{
		{"AA..", "AA..", "....", "...."},
		{"..BB", "BB..", "....", "...."},
	}

	expected := [][]string(nil)

	result := Solve(board, tetrominos)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected nil, but got %v", result)
	}
}
