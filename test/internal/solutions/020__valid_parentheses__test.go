package solutions_test

import (
	"github.com/billwallis/leetcode-go/internal/solutions"
	"testing"
)

func TestIsValid(t *testing.T) {
	var tests = []struct {
		s    string
		want bool
	}{
		{s: "()", want: true},
		{s: "()[]{}", want: true},
		{s: "(]", want: false},
		{s: "([{([{}])}])", want: true},
		{s: "(", want: false},
		{s: ")", want: false},
	}

	for _, test := range tests {
		got := solutions.IsValid(test.s)
		if got != test.want {
			t.Fatalf("IsValid(%s) == %v, want %v", test.s, got, test.want)
		}
	}
}
