package solutions_test

import (
	"github.com/Bilbottom/leetcode-go/internal/solutions"
	"testing"
)

func TestMinLengthOfStrings(t *testing.T) {
	var tests = []struct {
		strings []string
		want    int
	}{
		{strings: []string{}, want: 0},
		{strings: []string{"", ""}, want: 0},
		{strings: []string{"a", "b", "cd"}, want: 1},
		{strings: []string{"", "a", "bcd"}, want: 0},
	}

	for _, test := range tests {
		got := solutions.MinLengthOfStrings(test.strings)
		if got != test.want {
			t.Fatalf("MinLengthOfStrings(%v) == %d, want %d", test.strings, got, test.want)
		}
	}
}

func TestAreCharsAtPositionEqual(t *testing.T) {
	var tests = []struct {
		strings  []string
		position int
		want     bool
	}{
		{strings: []string{"a", "b"}, position: 0, want: false},
		{strings: []string{"ab", "ac", "ad"}, position: 0, want: true},
		{strings: []string{"ab", "ac", "ad"}, position: 1, want: false},
		{strings: []string{"a", "ab", "abc"}, position: 0, want: true},
	}

	for _, test := range tests {
		got := solutions.AreCharsAtPositionEqual(test.strings, test.position)
		if got != test.want {
			t.Fatalf("AreCharsAtPositionEqual(%v, %d) == %v, want %v", test.strings, test.position, got, test.want)
		}
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	var tests = []struct {
		strs []string
		want string
	}{
		{strs: []string{"flower", "flow", "flight"}, want: "fl"},
		{strs: []string{"dog", "racecar", "car"}, want: ""},
	}

	for _, test := range tests {
		got := solutions.LongestCommonPrefix(test.strs)
		if got != test.want {
			t.Fatalf("LongestCommonPrefix(%v) == %s, want %s", test.strs, got, test.want)
		}
	}
}
