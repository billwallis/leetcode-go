package solutions_test

import (
	"github.com/billwallis/leetcode-go/internal/solutions"
	"testing"
)

func TestLeftTrim(t *testing.T) {
	var tests = []struct {
		s    string
		c    string
		want string
	}{
		{s: "000000", c: "0", want: ""},
		{s: "     1", c: " ", want: "1"},
		{s: "  1   ", c: " ", want: "1   "},
		{s: "012345", c: "0", want: "12345"},
	}

	for _, test := range tests {
		got := solutions.LeftTrim(test.s, test.c)
		if got != test.want {
			t.Fatalf("LeftTrim(%s, %s) == %s, want %s", test.s, test.c, got, test.want)
		}
	}
}

func TestWithin(t *testing.T) {
	var tests = []struct {
		number int
		min    int
		max    int
		want   int
	}{
		{number: 0, min: 0, max: 0, want: 0},
		{number: 5, min: 0, max: 10, want: 5},
		{number: -1, min: 0, max: 10, want: 0},
		{number: 11, min: 0, max: 10, want: 10},
	}

	for _, test := range tests {
		got := solutions.Within(test.number, test.min, test.max)
		if got != test.want {
			t.Fatalf("Within(%d, %d, %d) == %d, want %d", test.number, test.min, test.max, got, test.want)
		}
	}
}

func TestMyAtoi(t *testing.T) {
	var tests = []struct {
		s    string
		want int
	}{
		{s: "42", want: 42},
		{s: "   -42", want: -42},
		{s: "4193 with words", want: 4193},
		{s: "-91283472332", want: -2147483648},
		{s: "words and 987", want: 0},
		{s: "", want: 0},
		{s: "+", want: 0},
		{s: "    0000000000000   ", want: 0},
		{s: "   +0 123", want: 0},
		{s: "20000000000000000000", want: 2147483647},
		{s: "    +00001 123 and this is a whole bunch of other text", want: 1},
	}

	for _, test := range tests {
		got := solutions.MyAtoi(test.s)
		if got != test.want {
			t.Fatalf("MyAtoi(%s) == %d, want %d", test.s, got, test.want)
		}
	}
}
