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