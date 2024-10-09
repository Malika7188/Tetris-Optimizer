package functions

import (
	"reflect"
	"testing"
)

func TestTrimTetris(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]string
		expected [][]string
	}{
		{
			name: "Multiple tetrominos with varying empty space",
			input: [][]string{
				{
					"B...",
					"BB..",
					"B...",
					"....",
				},
				{
					"..C.",
					".CC.",
					"..C.",
					"....",
				},
			},
			expected: [][]string{
				{
					"B",
					"BB",
					"B",
				},
				{
					"C",
					"CC",
					"C",
				},
			},
		},
		{
			name: "Tetromino with no empty space",
			input: [][]string{
				{
					"DDDD",
					"....",
					"....",
					"....",
				},
			},
			expected: [][]string{
				{
					"DDDD",
				},
			},
		},
		{
			name: "Empty input",
			input: [][]string{
				{
					"....",
					"....",
					"....",
					"....",
				},
			},
			expected: [][]string{
				{},
			},
		},
		{
			name: "Multiple tetrominos with different shapes",
			input: [][]string{
				{
					"E...",
					"EE..",
					".E..",
					"....",
				},
				{
					"..F.",
					"..F.",
					".FF.",
					"....",
				},
				{
					"..G.",
					"GGG.",
					"....",
					"....",
				},
			},
			expected: [][]string{
				{
					"E",
					"EE",
					".E",
				},
				{
					"F",
					"F",
					"FF",
				},
				{
					"G",
					"GGG",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TrimTetris(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("TrimTetris() = %v, want %v", result, tt.expected)
			}
		})
	}
}