package functions

import (
	"strings"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]string
		expected bool
	}{
		{
			name: "Valid tetrominos - 3 connections",
			input: [][]string{
				{"....", "..##", "..##", "...."},
				{"....", ".###", "...#", "...."},
			},
			expected: true,
		},
		{
			name: "Valid tetromino - 4 connections",
			input: [][]string{
				{"....", ".##.", ".##.", "...."},
			},
			expected: true,
		},
		{
			name: "Invalid tetromino - wrong size",
			input: [][]string{
				{"...", "..##", "..##", "...."},
			},
			expected: false,
		},
		{
			name: "Invalid tetromino - more than 4 blocks",
			input: [][]string{
				{"....", "..##", "..##", "...#"},
			},
			expected: false,
		},
		{
			name: "Valid tetromino - I shape (3 connections)",
			input: [][]string{
				{"..#.", "..#.", "..#.", "..#."},
			},
			expected: true,
		},
		{
			name: "Valid tetromino - L shape (3 connections)",
			input: [][]string{
				{"....", "...#", ".###", "...."},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Validate(tt.input)
			if got != tt.expected {
				t.Errorf("Validate() = %v, want %v", got, tt.expected)
				t.Errorf("Input:\n%s", formatTetromino(tt.input[0]))
			}
		})
	}
}

// Helper function to format a tetromino for error output
func formatTetromino(tetromino []string) string {
	return strings.Join(tetromino, "\n")
}
