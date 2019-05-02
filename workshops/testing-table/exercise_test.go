package exercise_test

import (
	"strings"
	"testing"
)

// Write a table drive test for strings.Index
// https://golang.org/pkg/strings/#Index
// Use the following test conditions
// |------------------------------------------------|
// | Value                     | Substring | Answer |
// |===========================|===========|========|
// | "Gophers are amazing!"    | "are"     | 8      |
// | "Testing in Go is fun."   | "fun"     | 17     |
// | "The answer is 42."       | "is"      | 11     |
// |------------------------------------------------|

// Rewrite the above test for strings.Index using subtests
// Enable parallel testing in the subtests
func TestStringsIndex(t *testing.T) {
	tcs := []struct {
		word   string
		substr string
		exp    int
	}{
		{word: "Gophers are amazing!", substr: "are", exp: 8},
		{word: "Testing in Go is fun.", substr: "fun", exp: 17},
		{word: "The answer is 42.", substr: "is", exp: 11},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(tc.word, func(t *testing.T) {
			t.Parallel()
			got := strings.Index(tc.word, tc.substr)
			t.Logf("testing %q", tc.word)
			if got != tc.exp {
				t.Errorf("unexpected value '%s' index of '%s' got: %d, exp: %d", tc.substr, tc.word, got, tc.exp)
			}
		})
	}
}
