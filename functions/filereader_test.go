package functions

import (
	"os"
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	// Create a temporary test file
	tmpfile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	// Test data to simulate the content of the file
	testData := `##..
.#..
.#..
....

....
.##.
##..
....

....
#...
###.
....

....
##..
##..
....
`

	// Write the test data to the temporary file
	if _, err := tmpfile.Write([]byte(testData)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Override the os.Args to simulate passing the temp file as argument
	os.Args = []string{"cmd", tmpfile.Name()}

	// Call the function and store the result
	result := ReadFile()

	// Expected output
	expected := [][]string{
		{"AA..", ".A..", ".A..", "...."},
		{"....", ".BB.", "BB..", "...."},
		{"....", "C...", "CCC.", "...."},
		{"....", "DD..", "DD..", "...."},
	}

	// Check if the result matches the expected output
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
