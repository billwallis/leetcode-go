package solutions_test

import (
	"github.com/Bilbottom/leetcode-go/internal/solutions"
	"testing"
)

func TestIsMatch(t *testing.T) {
	t.Skip("IsMatch not implemented yet.")

	var tests = []struct {
		s    string
		p    string
		want bool
	}{
		{s: "aa", p: "a", want: false},
		{s: "aa", p: "a*", want: true},
		{s: "ab", p: ".*", want: true},
		{s: "match", p: "m*h", want: false},
		{s: "mmmmh", p: "m*h", want: true},
		{s: "match", p: "m.*h", want: true},
		{s: "match", p: "m.*m", want: false},
		{s: "aa", p: "b*aa", want: true},
		{s: "aaaaa", p: "a.*aa", want: true},
		{s: "aabaa", p: ".*aa.*b.*aa.*", want: true},
		{s: "aabaa", p: ".*a.*b.*a.*", want: true},
	}

	for _, test := range tests {
		got := solutions.IsMatch(test.s, test.p)
		if got != test.want {
			t.Fatalf("IsMatch(%s, %s) == %v, want %v", test.s, test.p, got, test.want)
		}
	}
}
