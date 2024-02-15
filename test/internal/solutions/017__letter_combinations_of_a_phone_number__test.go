package solutions_test

import (
	"github.com/Bilbottom/leetcode-go/internal/solutions"
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	var tests = []struct {
		digits string
		want   []string
	}{
		{digits: "23", want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{digits: "", want: []string{}},
		{digits: "2", want: []string{"a", "b", "c"}},
	}

	for _, test := range tests {
		got := solutions.LetterCombinations(test.digits)
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("LetterCombinations(%s) == %v, want %v", test.digits, got, test.want)
		}
	}
}
