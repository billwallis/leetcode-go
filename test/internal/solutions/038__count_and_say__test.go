package solutions_test

import (
	"github.com/Bilbottom/leetcode-go/internal/solutions"
	"testing"
)

func TestSay(t *testing.T) {
	var tests = []struct {
		word string
		want string
	}{
		{word: "1", want: "11"},
		{word: "11", want: "21"},
		{word: "21", want: "1211"},
		{word: "1211", want: "111221"},
	}

	for _, test := range tests {
		got := solutions.Say(test.word)
		if got != test.want {
			t.Fatalf("Say(%s) == %s, want %s", test.word, got, test.want)
		}
	}
}

func TestCountAndSay(t *testing.T) {
	var tests = []struct {
		n    int
		want string
	}{
		{n: 1, want: "1"},
		{n: 4, want: "1211"},
	}

	for _, test := range tests {
		got := solutions.CountAndSay(test.n)
		if got != test.want {
			t.Fatalf("CountAndSay(%d) == %s, want %s", test.n, got, test.want)
		}
	}
}
