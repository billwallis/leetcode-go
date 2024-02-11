package solutions_test

import (
	"github.com/Bilbottom/leetcode-go/internal/solutions"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	var tests = []struct {
		num  int
		want string
	}{
		{num: 3, want: "III"},
		{num: 58, want: "LVIII"},
		{num: 1994, want: "MCMXCIV"},
	}

	for _, test := range tests {
		got := solutions.IntToRoman(test.num)
		if got != test.want {
			t.Fatalf("IntToRoman(%d) == %s, want %s", test.num, got, test.want)
		}
	}
}
